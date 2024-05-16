package mapper

import (
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
)

func ToEstablishmentData(entity *entity.Establishment) (*data.EstablishmentData, error) {
	accountID, err := ObjectIDFromString(entity.AccountID())
	if err != nil {
		return nil, err
	}
	return &data.EstablishmentData{
		AccountID: accountID,
		Name:      entity.Name(),
		Slug:      entity.Slug(),
		CreatedAt: entity.CreatedAt(),
	}, nil
}

func FromEstablishmentData(data *data.EstablishmentData) (*entity.Establishment, error) {
	establishment, err := entity.RestoreEstablishment(data.ID.Hex(), data.AccountID.Hex(), data.Name, data.Slug, data.CreatedAt)
	if err != nil {
		fmt.Println("error to restore establishment from database", err)
		return nil, err
	}
	return establishment, nil
}
