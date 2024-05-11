package repository

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type ServiceRepository interface {
	Save(context.Context, *entity.Service) (*entity.Service, error)
	FindByEstablishmentID(context.Context, string) ([]*entity.Service, error)
}
