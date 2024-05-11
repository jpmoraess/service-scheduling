package repository

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type EstablishmentRepository interface {
	Save(context.Context, *entity.Establishment) (*entity.Establishment, error)
}
