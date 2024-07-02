package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/application/repository"
)

type GetProfessionalOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetProfessionalUseCase struct {
	professionalRepository repository.ProfessionalRepository
}

func NewGetProfessionalUseCase(professionalRepository repository.ProfessionalRepository) *GetProfessionalUseCase {
	return &GetProfessionalUseCase{
		professionalRepository: professionalRepository,
	}
}

func (g *GetProfessionalUseCase) Execute(ctx context.Context, id string) (*GetProfessionalOutputDTO, error) {
	professional, err := g.professionalRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	output := &GetProfessionalOutputDTO{
		ID:   professional.ID(),
		Name: professional.Name(),
	}

	return output, nil
}
