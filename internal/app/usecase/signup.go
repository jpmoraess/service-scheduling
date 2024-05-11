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
	professionalRepository  repository.ProfessionalRepository
	establishmentRepository repository.EstablishmentRepository
}

func NewSignup(
	accountRepository repository.AccountRepository,
	professionalRepository repository.ProfessionalRepository,
	establishmentRepository repository.EstablishmentRepository,
) *Signup {
	return &Signup{
		accountRepository:       accountRepository,
		professionalRepository:  professionalRepository,
		establishmentRepository: establishmentRepository,
	}
}

func (a *Signup) Execute(ctx context.Context, input dto.SignupInput) error {
	// TODO: validate duplicated e-mail, phone and slug (2PC)
	encpw, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return err
	}
	account, err := entity.NewAccount(entity.OwnerType, input.Name, input.Email, input.PhoneNumber, string(encpw))
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
	savedEstablishment, err := a.establishmentRepository.Save(ctx, establishment)
	if err != nil {
		return err
	}
	professional, err := entity.NewProfessional(savedAccount.ID, savedEstablishment.ID, input.Name)
	if err != nil {
		return err
	}
	_, err = a.professionalRepository.Save(ctx, professional)
	if err != nil {
		return err
	}
	return nil
}
