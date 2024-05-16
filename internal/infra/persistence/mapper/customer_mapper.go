package mapper

import (
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
)

func ToCustomerData(entity *entity.Customer) (*data.CustomerData, error) {
	establishmentID, err := ObjectIDFromString(entity.EstablishmentID())
	if err != nil {
		return nil, err
	}

	return &data.CustomerData{
		EstablishmentID: establishmentID,
		Name:            entity.Name(),
		PhoneNumber:     entity.PhoneNumber(),
		Email:           entity.Email(),
		CreatedAt:       entity.CreatedAt(),
	}, nil
}

func FromCustomerData(data *data.CustomerData) (*entity.Customer, error) {
	customer, err := entity.RestoreCustomer(data.ID.Hex(), data.EstablishmentID.Hex(), data.Name, data.PhoneNumber, data.Email, data.CreatedAt)
	if err != nil {
		fmt.Println("error to restore customer from database", err)
		return nil, err
	}
	return customer, nil
}
