package mapper

import (
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
)

func ToSchedulingData(scheduling *entity.Scheduling) (*data.SchedulingData, error) {
	return &data.SchedulingData{
		ServiceID:       scheduling.ServiceID(),
		CustomerID:      scheduling.CustomerID(),
		ProfessionalID:  scheduling.ProfessionalID(),
		EstablishmentID: scheduling.EstablishmentID(),
	}, nil
}

func FromSchedulingData(data *data.SchedulingData) (*entity.Scheduling, error) {
	scheduling, err := entity.RestoreScheduling(data.ID.Hex(), data.ServiceID, data.CustomerID, data.ProfessionalID, data.EstablishmentID, data.Date, data.Time)
	if err != nil {
		fmt.Println("error to restore scheduling from database", err)
		return nil, err
	}
	return scheduling, nil
}
