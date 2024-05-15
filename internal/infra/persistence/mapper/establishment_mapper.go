package mapper

import (
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
)

func ToEstablishmentData(entity *entity.Establishment) (*data.EstablishmentData, error) {
	return &data.EstablishmentData{
		AccountID: entity.AccountID(),
		Name:      entity.Name(),
		Slug:      entity.Slug(),
	}, nil
}

func FromEstablishmentData(data *data.EstablishmentData) (*entity.Establishment, error) {
	establishment, err := entity.RestoreEstablishment(data.ID.Hex(), data.AccountID, data.Name, data.Slug)
	if err != nil {
		fmt.Println("error to restore establishment from database", err)
		return nil, err
	}
	return establishment, nil
}
