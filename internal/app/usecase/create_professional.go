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
	workPlanRepository      repository.WorkPlanRepository
}

func NewCreateProfessional(
	accountRepository repository.AccountRepository,
	professionalRepository repository.ProfessionalRepository,
	establishmentRepository repository.EstablishmentRepository,
	workPlanRepository repository.WorkPlanRepository,
) *CreateProfessional {
	return &CreateProfessional{
		accountRepository:       accountRepository,
		professionalRepository:  professionalRepository,
		establishmentRepository: establishmentRepository,
		workPlanRepository:      workPlanRepository,
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
	account, err = c.accountRepository.Save(ctx, account)
	if err != nil {
		return err
	}
	professional, err := entity.NewProfessional(account.ID, input.EstablishmentID, input.Name)
	if err != nil {
		return err
	}
	professional, err = c.professionalRepository.Save(ctx, professional)
	if err != nil {
		return err
	}
	workPlan, err := entity.DefaultWorkPlan()
	if err != nil {
		return err
	}
	workPlan.ProfessionalID = professional.ID
	_, err = c.workPlanRepository.Save(ctx, workPlan)
	if err != nil {
		return err
	}
	return nil
}
