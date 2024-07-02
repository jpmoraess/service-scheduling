package usecase

import (
	"context"
	"errors"
	"github.com/jpmoraess/service-scheduling/internal/application/repository"
)

type ResetPasswordInputDTO struct {
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}

type ResetPasswordUseCase struct {
	accountRepository       repository.AccountRepository
	passwordResetRepository repository.PasswordResetRepository
}

func NewResetPasswordUseCase(
	accountRepository repository.AccountRepository,
	passwordResetRepository repository.PasswordResetRepository,
) *ResetPasswordUseCase {
	return &ResetPasswordUseCase{
		accountRepository:       accountRepository,
		passwordResetRepository: passwordResetRepository,
	}
}

func (r *ResetPasswordUseCase) Execute(ctx context.Context, token string, input ResetPasswordInputDTO) error {
	passwordReset, err := r.passwordResetRepository.FindByToken(ctx, token)
	if err != nil {
		return err
	}

	if !passwordReset.IsExpiryTimeValid() {
		return errors.New("expired token")
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

	return nil
}
