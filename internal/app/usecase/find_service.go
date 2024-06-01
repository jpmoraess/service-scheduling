package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
)

type FindService struct {
	serviceRepository repository.ServiceRepository
}

func NewFindService(serviceRepository repository.ServiceRepository) *FindService {
	return &FindService{
		serviceRepository: serviceRepository,
	}
}

func (f *FindService) Execute(ctx context.Context, establishmentID string) ([]*dto.ServiceOutput, error) {
	services, err := f.serviceRepository.FindByEstablishmentID(ctx, establishmentID)
	if err != nil {
		return nil, err
	}
	var output []*dto.ServiceOutput
	for _, service := range services {
		output = append(output, &dto.ServiceOutput{
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
