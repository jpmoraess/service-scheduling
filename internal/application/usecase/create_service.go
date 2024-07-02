package usecase

import (
	"context"
	"time"

	"github.com/jpmoraess/service-scheduling/internal/application/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type CreateServiceInputDTO struct {
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	DurationInMinutes int64   `json:"durationInMinutes"`
}

type CreateServiceOutputDTO struct {
	ID string `json:"id"`
}

type CreateServiceUseCase struct {
	serviceRepository repository.ServiceRepository
}

func NewCreateServiceUseCase(serviceRepository repository.ServiceRepository) *CreateServiceUseCase {
	return &CreateServiceUseCase{serviceRepository: serviceRepository}
}

func (c *CreateServiceUseCase) Execute(ctx context.Context, input CreateServiceInputDTO) (*CreateServiceOutputDTO, error) {
	establishmentData, err := getEstablishmentData(ctx)
	if err != nil {
		return nil, err
	}

	service, err := entity.NewService(establishmentData.ID(), input.Name, input.Description, vo.NewMoney(input.Price), time.Duration(input.DurationInMinutes), true)
	if err != nil {
		return nil, err
	}

	service, err = c.serviceRepository.Save(ctx, service)
	if err != nil {
		return nil, err
	}

	return &CreateServiceOutputDTO{
		ID: service.ID(),
	}, nil
}
