package usecase

import (
	"context"
	"github.com/jpmoraess/service-scheduling/internal/application/repository"
)

type FindProfessionalOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FindProfessionalUseCase struct {
	professionalRepository repository.ProfessionalRepository
}

func NewFindProfessionalUseCase(professionalRepository repository.ProfessionalRepository) *FindProfessionalUseCase {
	return &FindProfessionalUseCase{professionalRepository: professionalRepository}
}

func (f *FindProfessionalUseCase) Execute(ctx context.Context, page, size int64) ([]*FindProfessionalOutputDTO, error) {
	establishmentData, err := getEstablishmentData(ctx)
	if err != nil {
		return nil, err
	}

	professionals, err := f.professionalRepository.Find(ctx, establishmentData.ID(), page, size)
	if err != nil {
		return nil, err
	}

	output := make([]*FindProfessionalOutputDTO, 0, len(professionals))
	for _, professional := range professionals {
		output = append(output, &FindProfessionalOutputDTO{
			ID:   professional.ID(),
			Name: professional.Name(),
		})
	}

	return output, nil
}
