package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
)

type GetProfessional struct {
	professionalRepository repository.ProfessionalRepository
}

func NewGetProfessional(professionalRepository repository.ProfessionalRepository) *GetProfessional {
	return &GetProfessional{
		professionalRepository: professionalRepository,
	}
}

func (g *GetProfessional) Execute(ctx context.Context, id string) (*dto.ProfessionalOutput, error) {
	professional, err := g.professionalRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	professionalOutput := &dto.ProfessionalOutput{
		ID:   professional.ID(),
		Name: professional.Name(),
	}
	return professionalOutput, nil
}
