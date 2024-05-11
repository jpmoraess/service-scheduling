package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

type CreateProfessional struct {
	accountRepository       repository.AccountRepository
	professionalRepository  repository.ProfessionalRepository
	establishmentRepository repository.EstablishmentRepository
}

func NewCreateProfessional(
	accountRepository repository.AccountRepository,
	professionalRepository repository.ProfessionalRepository,
	establishmentRepository repository.EstablishmentRepository,
) *CreateProfessional {
	return &CreateProfessional{
		accountRepository:       accountRepository,
		professionalRepository:  professionalRepository,
		establishmentRepository: establishmentRepository,
	}
}

func (c *CreateProfessional) Execute(ctx context.Context, input dto.CreateProfessionalInput) error {
	// TODO: validate existence of establishment
	encpw, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return err
	}
	account, err := entity.NewAccount(entity.ProfessionalType, input.Name, input.Email, input.PhoneNumber, string(encpw))
	if err != nil {
		return err
	}
	savedAccount, err := c.accountRepository.Save(ctx, account)
	if err != nil {
		return err
	}
	professional, err := entity.NewProfessional(savedAccount.ID, input.EstablishmentID, input.Name)
	if err != nil {
		return err
	}
	_, err = c.professionalRepository.Save(ctx, professional)
	if err != nil {
		return err
	}
	return nil
}
