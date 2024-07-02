package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/application/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
	"golang.org/x/crypto/bcrypt"
)

type CreateProfessionalInputDTO struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type CreateProfessionalOutputDTO struct {
	ID string `json:"id"`
}

type CreateProfessionalUseCase struct {
	accountRepository      repository.AccountRepository
	professionalRepository repository.ProfessionalRepository
}

func NewCreateProfessionalUseCase(
	accountRepository repository.AccountRepository,
	professionalRepository repository.ProfessionalRepository,
) *CreateProfessionalUseCase {
	return &CreateProfessionalUseCase{
		accountRepository:      accountRepository,
		professionalRepository: professionalRepository,
	}
}

func (c *CreateProfessionalUseCase) Execute(ctx context.Context, input CreateProfessionalInputDTO) (*CreateProfessionalOutputDTO, error) {
	establishmentData, err := getEstablishmentData(ctx)
	if err != nil {
		return nil, err
	}

	encryptedPwd, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return nil, err
	}

	account, err := entity.NewAccount(vo.ProfessionalType, input.Name, input.Email, input.PhoneNumber, string(encryptedPwd))
	if err != nil {
		return nil, err
	}

	account, err = c.accountRepository.Save(ctx, account)
	if err != nil {
		return nil, err
	}

	professional, err := entity.NewProfessional(account.ID(), establishmentData.ID(), input.Name)
	if err != nil {
		return nil, err
	}

	professional, err = c.professionalRepository.Save(ctx, professional)
	if err != nil {
		return nil, err
	}

	return &CreateProfessionalOutputDTO{
		ID: professional.ID(),
	}, nil
}
