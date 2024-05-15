package mapper

import (
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
)

func ToCustomerData(entity *entity.Customer) (*data.CustomerData, error) {
	return &data.CustomerData{
		EstablishmentID: entity.EstablishmentID(),
		Name:            entity.Name(),
		PhoneNumber:     entity.PhoneNumber(),
		Email:           entity.Email(),
	}, nil
}

func FromCustomerData(data *data.CustomerData) (*entity.Customer, error) {
	customer, err := entity.RestoreCustomer(data.ID.Hex(), data.EstablishmentID, data.Name, data.PhoneNumber, data.Email)
	if err != nil {
		fmt.Println("error to restore customer from database", err)
		return nil, err
	}
	return customer, nil
}
