package repository

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type PasswordResetRepository interface {
	Save(context.Context, *entity.PasswordReset) error
	FindByToken(context.Context, string) (*entity.PasswordReset, error)
}
