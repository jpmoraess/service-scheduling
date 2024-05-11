package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

type Signup struct {
	accountRepository       repository.AccountRepository
	establishmentRepository repository.EstablishmentRepository
}

func NewSignup(accountRepository repository.AccountRepository, establishmentRepository repository.EstablishmentRepository) *Signup {
	return &Signup{
		accountRepository:       accountRepository,
		establishmentRepository: establishmentRepository,
	}
}

// TODO: validate duplicated e-mail, phone and slug (2PC)
func (a *Signup) Execute(ctx context.Context, input dto.SignupInput) error {
	encpw, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return err
	}
	account, err := entity.NewAccount(input.Name, input.Email, input.PhoneNumber, string(encpw))
	if err != nil {
		return err
	}
	savedAccount, err := a.accountRepository.Save(ctx, account)
	if err != nil {
		return err
	}
	establishment, err := entity.NewEstablishment(savedAccount.ID, input.EstablishmentName, "slug")
	if err != nil {
		return err
	}
	_, err = a.establishmentRepository.Save(ctx, establishment)
	if err != nil {
		return err
	}
	// TODO: criar profissional
	return nil
}
