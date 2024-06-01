package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
)

type FindProfessional struct {
	professionalRepository repository.ProfessionalRepository
}

func NewFindProfessional(professionalRepository repository.ProfessionalRepository) *FindProfessional {
	return &FindProfessional{professionalRepository: professionalRepository}
}

func (f *FindProfessional) Execute(ctx context.Context, page, size int64) ([]*dto.ProfessionalOutput, error) {
	establishmentData, err := getEstablishmentData(ctx)
	if err != nil {
		return nil, err
	}
	professionals, err := f.professionalRepository.Find(ctx, establishmentData.ID(), page, size)
	if err != nil {
		return nil, err
	}
	output := make([]*dto.ProfessionalOutput, 0, len(professionals))
	for _, professional := range professionals {
		output = append(output, &dto.ProfessionalOutput{
			ID:   professional.ID(),
			Name: professional.Name(),
		})
	}
	return output, nil
}
