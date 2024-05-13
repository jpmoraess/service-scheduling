package usecase

import (
	"context"
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type CreateScheduling struct {
	schedulingRepository repository.SchedulingRepository
}

func NewCreateScheduling(schedulingRepository repository.SchedulingRepository) *CreateScheduling {
	return &CreateScheduling{
		schedulingRepository: schedulingRepository,
	}
}

func (c *CreateScheduling) Execute(ctx context.Context, input dto.CreateSchedulingInput) error {
	// TODO: remover os parses de data daqui, utilizar value object??
	date, err := parseDateTime(input.Date, "2006-01-02")
	if err != nil {
		return err
	}

	time, err := parseDateTime(input.Time, "15:04")
	if err != nil {
		return err
	}

	// TODO: aplicar algumas validacoes

	scheduling, err := entity.NewScheduling(input.ServiceID, input.CustomerID, input.ProfessionalID, input.EstablishmentID, date, time)
	if err != nil {
		return err
	}

	scheduling, err = c.schedulingRepository.Save(ctx, scheduling)
	if err != nil {
		return err
	}
	fmt.Println("success scheduling: ", scheduling)
	return nil
}
