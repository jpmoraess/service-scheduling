package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type CreateCustomer struct {
	customerRepository repository.CustomerRepository
}

func NewCreateCustomer(customerRepository repository.CustomerRepository) *CreateCustomer {
	return &CreateCustomer{customerRepository: customerRepository}
}

func (c *CreateCustomer) Execute(ctx context.Context, input dto.CreateCustomerInput) error {
	establishmentData, err := getEstablishmentData(ctx)
	if err != nil {
		return err
	}

	// TODO: validate duplicated customer by establishment and phone number
	customer, err := entity.NewCustomer(establishmentData.ID(), input.Name, input.PhoneNumber, input.Email)
	if err != nil {
		return err
	}

	_, err = c.customerRepository.Save(ctx, customer)
	if err != nil {
		return err
	}

	return nil
}
