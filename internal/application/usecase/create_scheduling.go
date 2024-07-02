package usecase

import (
	"context"
	"sync"

	"github.com/jpmoraess/service-scheduling/internal/application/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type CreateSchedulingInputDTO struct {
	ServiceID       string `json:"serviceID"`
	CustomerID      string `json:"customerID"`
	ProfessionalID  string `json:"professionalID"`
	EstablishmentID string `json:"establishmentID"`
	Date            string `json:"date"`
	Time            string `json:"time"`
}

type CreateSchedulingOutputDTO struct {
	ID string `json:"id"`
}

type CreateSchedulingUseCase struct {
	serviceRepository       repository.ServiceRepository
	customerRepository      repository.CustomerRepository
	professionalRepository  repository.ProfessionalRepository
	establishmentRepository repository.EstablishmentRepository
	schedulingRepository    repository.SchedulingRepository
}

func NewCreateSchedulingUseCase(
	serviceRepository repository.ServiceRepository,
	customerRepository repository.CustomerRepository,
	professionalRepository repository.ProfessionalRepository,
	establishmentRepository repository.EstablishmentRepository,
	schedulingRepository repository.SchedulingRepository,
) *CreateSchedulingUseCase {
	return &CreateSchedulingUseCase{
		serviceRepository:       serviceRepository,
		customerRepository:      customerRepository,
		professionalRepository:  professionalRepository,
		establishmentRepository: establishmentRepository,
		schedulingRepository:    schedulingRepository,
	}
}

func (c *CreateSchedulingUseCase) Execute(ctx context.Context, input CreateSchedulingInputDTO) (*CreateSchedulingOutputDTO, error) {
	resp, err := c.getData(ctx, input)
	if err != nil {
		return nil, err
	}

	scheduling, err := entity.NewScheduling(resp.service.ID(), resp.customer.ID(), resp.professional.ID(), resp.establishment.ID(), input.Date, input.Time)
	if err != nil {
		return nil, err
	}

	scheduling, err = c.schedulingRepository.Save(ctx, scheduling)
	if err != nil {
		return nil, err
	}

	return &CreateSchedulingOutputDTO{
		ID: scheduling.ID(),
	}, nil
}

type ResponseData struct {
	service       *entity.Service
	customer      *entity.Customer
	professional  *entity.Professional
	establishment *entity.Establishment
}

func (c *CreateSchedulingUseCase) getData(ctx context.Context, input CreateSchedulingInputDTO) (*ResponseData, error) {
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
