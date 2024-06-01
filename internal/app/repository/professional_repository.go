package repository

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type ProfessionalRepository interface {
	Save(context.Context, *entity.Professional) (*entity.Professional, error)
	Get(context.Context, string) (*entity.Professional, error)
	Find(ctx context.Context, establishmentID string, page int64, size int64) ([]*entity.Professional, error)
}
