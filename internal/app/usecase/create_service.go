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
	serviceRepository repository.ServiceRepository
}

func NewCreateService(serviceRepository repository.ServiceRepository) *CreateService {
	return &CreateService{
		serviceRepository: serviceRepository,
	}
}

func (c *CreateService) Execute(ctx context.Context, input dto.CreateServiceInput) error {
	// TODO: validate if establishment exists
	service, err := entity.NewService(input.EstablishmentID, input.Name, input.Description, input.Price, time.Duration(input.DurationInMinutes))
	if err != nil {
		return err
	}
	savedService, err := c.serviceRepository.Save(ctx, service)
	if err != nil {
		return err
	}
	fmt.Printf("service created successfully: %+v\n", savedService)
	return nil
}
