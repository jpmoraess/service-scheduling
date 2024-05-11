package repository

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type ProfessionalRepository interface {
	Save(context.Context, *entity.Professional) (*entity.Professional, error)
}
