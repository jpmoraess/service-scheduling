package usecase

import (
	"context"
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
)

type ResetPassword struct {
	accountRepository       repository.AccountRepository
	passwordResetRepository repository.PasswordResetRepository
}

func NewResetPassword(
	accountRepository repository.AccountRepository,
	passwordResetRepository repository.PasswordResetRepository,
) *ResetPassword {
	return &ResetPassword{
		accountRepository:       accountRepository,
		passwordResetRepository: passwordResetRepository,
	}
}

func (r *ResetPassword) Execute(ctx context.Context, token string, input dto.ResetPasswordInput) error {
	passwordReset, err := r.passwordResetRepository.FindByToken(ctx, token)
	if err != nil {
		return err
	}

	if !passwordReset.IsExpiryTimeValid() {
		return fmt.Errorf("invalid token")
	}

	account, err := r.accountRepository.Get(ctx, passwordReset.AccountID())
	if err != nil {
		return err
	}

	if err := account.ResetPassword(input.NewPassword); err != nil {
		return err
	}

	account, err = r.accountRepository.Save(ctx, account)
	if err != nil {
		return err
	}
	fmt.Println("password reset successfully: ", account.ID())

	return nil
}
