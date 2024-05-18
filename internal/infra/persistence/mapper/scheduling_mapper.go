package mapper

import (
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/util"
)

func ToSchedulingData(scheduling *entity.Scheduling) (*data.SchedulingData, error) {
	serviceID, err := util.GetObjectID(scheduling.ServiceID())
	if err != nil {
		return nil, err
	}
	customerID, err := util.GetObjectID(scheduling.CustomerID())
	if err != nil {
		return nil, err
	}
	professionalID, err := util.GetObjectID(scheduling.ProfessionalID())
	if err != nil {
		return nil, err
	}
	establishmentID, err := util.GetObjectID(scheduling.EstablishmentID())
	if err != nil {
		return nil, err
	}
	return &data.SchedulingData{
		ServiceID:       serviceID,
		CustomerID:      customerID,
		ProfessionalID:  professionalID,
		EstablishmentID: establishmentID,
		Date:            scheduling.Date().String(),
		Time:            scheduling.Time().String(),
		CreatedAt:       scheduling.CreatedAt(),
	}, nil
}

func FromSchedulingData(data *data.SchedulingData) (*entity.Scheduling, error) {
	scheduling, err := entity.RestoreScheduling(data.ID.Hex(), data.ServiceID.Hex(), data.CustomerID.Hex(), data.ProfessionalID.Hex(), data.EstablishmentID.Hex(), data.Date, data.Time, data.CreatedAt)
	if err != nil {
		fmt.Println("error to restore scheduling from database", err)
		return nil, err
	}
	return scheduling, nil
}
