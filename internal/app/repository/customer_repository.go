package repository

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type CustomerRepository interface {
	Save(context.Context, *entity.Customer) (*entity.Customer, error)
	Get(context.Context, string) (*entity.Customer, error)
	GetByEstablishmentIDAndPhoneNumber(context.Context, string, string) (*entity.Customer, error)
}
