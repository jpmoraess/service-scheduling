package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type CreateService struct {
	serviceRepository       repository.ServiceRepository
	establishmentRepository repository.EstablishmentRepository
}

func NewCreateService(serviceRepository repository.ServiceRepository, establishmentRepository repository.EstablishmentRepository) *CreateService {
	return &CreateService{
		serviceRepository:       serviceRepository,
		establishmentRepository: establishmentRepository,
	}
}

func (c *CreateService) Execute(ctx context.Context, input dto.CreateServiceInput) error {
	tokenData, ok := ctx.Value("account").(*entity.Account)
	if !ok {
		return fmt.Errorf("error") // TODO: treat error better
	}

	establishment, err := c.establishmentRepository.GetByAccountID(ctx, tokenData.ID)
	if err != nil {
		return fmt.Errorf("establishment not found") // TODO: treat error better
	}

	service, err := entity.NewService(establishment.ID, input.Name, input.Description, input.Price, time.Duration(input.DurationInMinutes))
	if err != nil {
		return err
	}
	service, err = c.serviceRepository.Save(ctx, service)
	if err != nil {
		return err
	}
	fmt.Printf("service created successfully: %+v\n", service)
	return nil
}
