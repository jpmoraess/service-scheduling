package repository

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type SchedulingRepository interface {
	Save(context.Context, *entity.Scheduling) (*entity.Scheduling, error)
}
