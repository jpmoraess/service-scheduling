package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type CreateService struct {
	serviceRepository repository.ServiceRepository
}

func NewCreateService(serviceRepository repository.ServiceRepository) *CreateService {
	return &CreateService{serviceRepository: serviceRepository}
}

func (c *CreateService) Execute(ctx context.Context, input dto.CreateServiceInput) error {
	establishmentData, err := getEstablishmentData(ctx)
	if err != nil {
		return err
	}

	service, err := entity.NewService(establishmentData.ID(), input.Name, input.Description, vo.NewMoney(input.Price), time.Duration(input.DurationInMinutes), true)
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
