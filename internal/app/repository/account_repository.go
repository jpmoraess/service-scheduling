package repository

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type AccountRepository interface {
	Save(context.Context, *entity.Account) (*entity.Account, error)
	GetAccountByEmail(context.Context, string) (*entity.Account, error)
}
