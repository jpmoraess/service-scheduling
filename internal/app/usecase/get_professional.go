package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type GetProfessional struct {
	professionalRepository repository.ProfessionalRepository
}

func NewGetProfessional(professionalRepository repository.ProfessionalRepository) *GetProfessional {
	return &GetProfessional{
		professionalRepository: professionalRepository,
	}
}

func (g *GetProfessional) Execute(ctx context.Context, id string) (*entity.Professional, error) {
	professional, err := g.professionalRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return professional, nil
}
