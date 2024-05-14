package usecase

import (
	"context"
	"sync"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type CreateScheduling struct {
	serviceRepository       repository.ServiceRepository
	customerRepository      repository.CustomerRepository
	professionalRepository  repository.ProfessionalRepository
	establishmentRepository repository.EstablishmentRepository
	schedulingRepository    repository.SchedulingRepository
}

func NewCreateScheduling(
	serviceRepository repository.ServiceRepository,
	customerRepository repository.CustomerRepository,
	professionalRepository repository.ProfessionalRepository,
	establishmentRepository repository.EstablishmentRepository,
	schedulingRepository repository.SchedulingRepository,
) *CreateScheduling {
	return &CreateScheduling{
		serviceRepository:       serviceRepository,
		customerRepository:      customerRepository,
		professionalRepository:  professionalRepository,
		establishmentRepository: establishmentRepository,
		schedulingRepository:    schedulingRepository,
	}
}

func (c *CreateScheduling) Execute(ctx context.Context, input dto.CreateSchedulingInput) error {
	// TODO: remover os parses de data daqui, utilizar value object??
	date, err := parseDateTime(input.Date, "2006-01-02")
	if err != nil {
		return err
	}

	time, err := parseDateTime(input.Time, "15:04")
	if err != nil {
		return err
	}

	resp, err := c.getData(ctx, input)
	if err != nil {
		return err
	}

	scheduling, err := entity.NewScheduling(resp.service, resp.customer, resp.professional, resp.establishment, date, time)
	if err != nil {
		return err
	}

	_, err = c.schedulingRepository.Save(ctx, scheduling)
	if err != nil {
		return err
	}

	return nil
}

type ResponseData struct {
	service       *entity.Service
	customer      *entity.Customer
	professional  *entity.Professional
	establishment *entity.Establishment
}

func (c *CreateScheduling) getData(ctx context.Context, input dto.CreateSchedulingInput) (*ResponseData, error) {
	var (
		resp  = &ResponseData{}
		wg    = sync.WaitGroup{}
		errCh = make(chan error, 4)
	)

	wg.Add(4)
	go func() {
		defer wg.Done()
		service, err := c.serviceRepository.Get(ctx, input.ServiceID)
		if err != nil {
			errCh <- err
			return
		}
		resp.service = service
	}()
	go func() {
		defer wg.Done()
		customer, err := c.customerRepository.Get(ctx, input.CustomerID)
		if err != nil {
			errCh <- err
			return
		}
		resp.customer = customer
	}()
	go func() {
		defer wg.Done()
		professional, err := c.professionalRepository.Get(ctx, input.ProfessionalID)
		if err != nil {
			errCh <- err
			return
		}
		resp.professional = professional
	}()
	go func() {
		defer wg.Done()
		establishment, err := c.establishmentRepository.Get(ctx, input.EstablishmentID)
		if err != nil {
			errCh <- err
			return
		}
		resp.establishment = establishment
	}()
	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}
