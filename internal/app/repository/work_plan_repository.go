package repository

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type WorkPlanRepository interface {
	Save(context.Context, *entity.WorkPlan) (*entity.WorkPlan, error)
}
