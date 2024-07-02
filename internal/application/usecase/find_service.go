package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/application/repository"
)

type FindServiceOutputDTO struct {
	ID                string  `json:"ID"`
	EstablishmentID   string  `json:"establishmentID"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	DurationInMinutes int64   `json:"durationInMinutes"`
	Available         bool    `json:"available"`
}

type FindServiceUseCase struct {
	serviceRepository repository.ServiceRepository
}

func NewFindServiceUseCase(serviceRepository repository.ServiceRepository) *FindServiceUseCase {
	return &FindServiceUseCase{
		serviceRepository: serviceRepository,
	}
}

func (f *FindServiceUseCase) Execute(ctx context.Context, establishmentID string) ([]*FindServiceOutputDTO, error) {
	services, err := f.serviceRepository.FindByEstablishmentID(ctx, establishmentID)
	if err != nil {
		return nil, err
	}

	var output []*FindServiceOutputDTO
	for _, service := range services {
		output = append(output, &FindServiceOutputDTO{
			ID:                service.ID(),
			EstablishmentID:   service.EstablishmentID(),
			Name:              service.Name(),
			Description:       service.Description(),
			Price:             service.Price().AmountFloat64(),
			DurationInMinutes: int64(service.Duration()),
			Available:         service.Available(),
		})
	}

	return output, nil
}
