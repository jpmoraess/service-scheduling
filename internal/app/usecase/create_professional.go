package usecase

import (
	"context"
	"fmt"

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
	tokenData, ok := ctx.Value("account").(*entity.Account)
	if !ok {
		return fmt.Errorf("error") // TODO: treat error better
	}

	establishment, err := c.establishmentRepository.GetByAccountID(ctx, tokenData.ID)
	if err != nil {
		return fmt.Errorf("establishment not found") // TODO: treat error better
	}

	encpw, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return err
	}

	account, err := entity.NewAccount(entity.ProfessionalType, input.Name, input.Email, input.PhoneNumber, string(encpw))
	if err != nil {
		return err
	}

	account, err = c.accountRepository.Save(ctx, account)
	if err != nil {
		return err
	}

	professional, err := entity.NewProfessional(account.ID, establishment.ID, input.Name)
	if err != nil {
		return err
	}

	_, err = c.professionalRepository.Save(ctx, professional)
	if err != nil {
		return err
	}

	return nil
}
