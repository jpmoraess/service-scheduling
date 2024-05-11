package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
)

type FindServices struct {
	serviceRepository repository.ServiceRepository
}

func NewFindServices(serviceRepository repository.ServiceRepository) *FindServices {
	return &FindServices{
		serviceRepository: serviceRepository,
	}
}

func (f *FindServices) Execute(ctx context.Context, establishmentID string) ([]*dto.ServiceOutput, error) {
	services, err := f.serviceRepository.FindByEstablishmentID(ctx, establishmentID)
	if err != nil {
		return nil, err
	}
	var output []*dto.ServiceOutput
	for _, service := range services {
		output = append(output, &dto.ServiceOutput{
			ID:                service.ID,
			EstablishmentID:   service.EstablishmentID,
			Name:              service.Name,
			Description:       service.Description,
			Price:             service.Price,
			DurationInMinutes: int64(service.Duration),
			Available:         service.Available,
		})
	}
	return output, nil
}
