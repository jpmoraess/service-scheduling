package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
	"golang.org/x/crypto/bcrypt"
)

type CreateProfessional struct {
	accountRepository      repository.AccountRepository
	professionalRepository repository.ProfessionalRepository
}

func NewCreateProfessional(accountRepository repository.AccountRepository, professionalRepository repository.ProfessionalRepository) *CreateProfessional {
	return &CreateProfessional{accountRepository: accountRepository, professionalRepository: professionalRepository}
}

/**
	1. Validar e-mail e telefone do profissional cadastrado
	2. Enviar e-mail para profissional cadastrar sua senha e baixar o app
**/

func (c *CreateProfessional) Execute(ctx context.Context, input dto.CreateProfessionalInput) error {
	establishmentData, err := getEstablishmentData(ctx)
	if err != nil {
		return err
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

	professional, err := entity.NewProfessional(account.ID(), establishmentData.ID(), input.Name)
	if err != nil {
		return err
	}

	_, err = c.professionalRepository.Save(ctx, professional)
	if err != nil {
		return err
	}

	return nil
}
