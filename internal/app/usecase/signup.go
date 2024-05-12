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
	workPlanRepository      repository.WorkPlanRepository
}

func NewSignup(
	accountRepository repository.AccountRepository,
	professionalRepository repository.ProfessionalRepository,
	establishmentRepository repository.EstablishmentRepository,
	workPlanRepository repository.WorkPlanRepository,
) *Signup {
	return &Signup{
		accountRepository:       accountRepository,
		professionalRepository:  professionalRepository,
		establishmentRepository: establishmentRepository,
		workPlanRepository:      workPlanRepository,
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
	account, err = a.accountRepository.Save(ctx, account)
	if err != nil {
		return err
	}
	establishment, err := entity.NewEstablishment(account.ID, input.EstablishmentName, "slug")
	if err != nil {
		return err
	}
	establishment, err = a.establishmentRepository.Save(ctx, establishment)
	if err != nil {
		return err
	}
	professional, err := entity.NewProfessional(account.ID, establishment.ID, input.Name)
	if err != nil {
		return err
	}
	professional, err = a.professionalRepository.Save(ctx, professional)
	if err != nil {
		return err
	}
	workPlan, err := entity.DefaultWorkPlan()
	if err != nil {
		return err
	}
	workPlan.ProfessionalID = professional.ID
	_, err = a.workPlanRepository.Save(ctx, workPlan)
	if err != nil {
		return err
	}
	return nil
}
