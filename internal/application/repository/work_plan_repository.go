package repository

import (
	"context"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type WorkPlanRepository interface {
	Save(ctx context.Context, plan *entity.WorkPlan) (*entity.WorkPlan, error)
}
