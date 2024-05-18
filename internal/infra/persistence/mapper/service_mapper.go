package mapper

import (
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/util"
)

func ToServiceData(service *entity.Service) (*data.ServiceData, error) {
	establishmentID, err := util.GetObjectID(service.EstablishmentID())
	if err != nil {
		return nil, err
	}
	return &data.ServiceData{
		EstablishmentID: establishmentID,
		Name:            service.Name(),
		Description:     service.Description(),
		Price:           service.Price().AmountFloat64(),
		Duration:        service.Duration(),
		Available:       service.Available(),
	}, nil
}

func FromServiceData(data *data.ServiceData) (*entity.Service, error) {
	service, err := entity.RestoreService(data.ID.Hex(), data.EstablishmentID.Hex(), data.Name, data.Description, data.Price, data.Duration, data.Available, data.CreatedAt)
	if err != nil {
		fmt.Println("error to restore service from database", err)
		return nil, err
	}
	return service, nil
}
