package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
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
