package repository

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type EstablishmentRepository interface {
	Save(context.Context, *entity.Establishment) (*entity.Establishment, error)
	GetByAccountID(context.Context, string) (*entity.Establishment, error)
}
