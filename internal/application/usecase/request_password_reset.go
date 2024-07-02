package usecase

import (
	"context"
	"github.com/jpmoraess/service-scheduling/internal/application/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type RequestPasswordResetInputDTO struct {
	Email string `json:"email"`
}

type RequestPasswordResetUseCase struct {
	accountRepository       repository.AccountRepository
	passwordResetRepository repository.PasswordResetRepository
}

func NewRequestPasswordResetUseCase(
	accountRepository repository.AccountRepository,
	passwordResetRepository repository.PasswordResetRepository,
) *RequestPasswordResetUseCase {
	return &RequestPasswordResetUseCase{
		accountRepository:       accountRepository,
		passwordResetRepository: passwordResetRepository,
	}
}

func (r *RequestPasswordResetUseCase) Execute(ctx context.Context, input RequestPasswordResetInputDTO) error {
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

	return nil
}
