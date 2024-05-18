package usecase

import (
	"context"
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type RequestPasswordReset struct {
	accountRepository       repository.AccountRepository
	passwordResetRepository repository.PasswordResetRepository
}

func NewRequestPasswordReset(
	accountRepository repository.AccountRepository,
	passwordResetRepository repository.PasswordResetRepository,
) *RequestPasswordReset {
	return &RequestPasswordReset{
		accountRepository:       accountRepository,
		passwordResetRepository: passwordResetRepository,
	}
}

func (r *RequestPasswordReset) Execute(ctx context.Context, input dto.RequestPasswordResetInput) error {
	account, err := r.accountRepository.GetAccountByEmail(ctx, input.Email)
	if err != nil {
		return err
	}
	passwordReset, err := entity.NewPasswordReset(account.ID())
	if err != nil {
		return err
	}
	if err := r.passwordResetRepository.Save(ctx, passwordReset); err != nil {
		return err
	}
	// TODO: send e-mail
	fmt.Printf("sending password reset request: %+v", passwordReset)
	return nil
}
