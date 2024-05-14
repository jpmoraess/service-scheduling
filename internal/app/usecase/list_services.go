package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
)

type ListServices struct {
	serviceRepository repository.ServiceRepository
}

func NewListServices(serviceRepository repository.ServiceRepository) *ListServices {
	return &ListServices{
		serviceRepository: serviceRepository,
	}
}

func (f *ListServices) Execute(ctx context.Context, establishmentID string) ([]*dto.ServiceOutput, error) {
	services, err := f.serviceRepository.FindByEstablishmentID(ctx, establishmentID)
	if err != nil {
		return nil, err
	}
	var output []*dto.ServiceOutput
	for _, service := range services {
		output = append(output, &dto.ServiceOutput{
			ID:                service.GetID(),
			EstablishmentID:   service.GetEstablishmentID(),
			Name:              service.GetName(),
			Description:       service.GetDescription(),
			Price:             service.GetPrice().GetAmountFloat64(),
			DurationInMinutes: int64(service.GetDuration()),
			Available:         service.GetAvailable(),
		})
	}
	return output, nil
}
