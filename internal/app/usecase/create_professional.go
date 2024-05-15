package usecase

import (
	"context"
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
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
	authData, err := getAuthData(ctx)
	if err != nil {
		return err
	}

	establishment, err := c.establishmentRepository.GetByAccountID(ctx, authData.ID())
	if err != nil {
		return fmt.Errorf("establishment not found") // TODO: treat error better
	}

	encpw, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return err
	}

	account, err := entity.NewAccount(vo.ProfessionalType, input.Name, input.Email, input.PhoneNumber, string(encpw))
	if err != nil {
		return err
	}

	account, err = c.accountRepository.Save(ctx, account)
	if err != nil {
		return err
	}

	professional, err := entity.NewProfessional(account.ID(), establishment.ID(), input.Name)
	if err != nil {
		return err
	}

	_, err = c.professionalRepository.Save(ctx, professional)
	if err != nil {
		return err
	}

	return nil
}
