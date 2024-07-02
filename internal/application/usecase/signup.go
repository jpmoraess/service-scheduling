package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/application/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type SignupInputDTO struct {
	Name              string `json:"name"`
	EstablishmentName string `json:"establishmentName"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	Password          string `json:"password"`
}

type SignupUseCase struct {
	accountRepository       repository.AccountRepository
	professionalRepository  repository.ProfessionalRepository
	establishmentRepository repository.EstablishmentRepository
}

func NewSignupUseCase(
	accountRepository repository.AccountRepository,
	professionalRepository repository.ProfessionalRepository,
	establishmentRepository repository.EstablishmentRepository,
) *SignupUseCase {
	return &SignupUseCase{
		accountRepository:       accountRepository,
		professionalRepository:  professionalRepository,
		establishmentRepository: establishmentRepository,
	}
}

func (a *SignupUseCase) Execute(ctx context.Context, input SignupInputDTO) error {
	account, err := entity.NewAccount(vo.OwnerType, input.Name, input.Email, input.PhoneNumber, input.Password)
	if err != nil {
		return err
	}

	account, err = a.accountRepository.Save(ctx, account)
	if err != nil {
		return err
	}

	establishment, err := entity.NewEstablishment(account.ID(), input.EstablishmentName, "slug")
	if err != nil {
		return err
	}

	establishment, err = a.establishmentRepository.Save(ctx, establishment)
	if err != nil {
		return err
	}

	professional, err := entity.NewProfessional(account.ID(), establishment.ID(), input.Name)
	if err != nil {
		return err
	}

	_, err = a.professionalRepository.Save(ctx, professional)
	if err != nil {
		return err
	}
	return nil
}
