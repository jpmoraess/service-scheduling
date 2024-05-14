package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type CreateCustomer struct {
	customerRepository      repository.CustomerRepository
	establishmentRepository repository.EstablishmentRepository
}

func NewCreateCustomer(customerRepository repository.CustomerRepository, establishmentRepository repository.EstablishmentRepository) *CreateCustomer {
	return &CreateCustomer{
		customerRepository:      customerRepository,
		establishmentRepository: establishmentRepository,
	}
}

func (c *CreateCustomer) Execute(ctx context.Context, input dto.CreateCustomerInput) error {
	authData, err := getAuthData(ctx)
	if err != nil {
		return err
	}

	establishment, err := c.establishmentRepository.GetByAccountID(ctx, authData.GetID())
	if err != nil {
		return err
	}

	// TODO: validate duplicated customer by establishment and phone number

	customer, err := entity.NewCustomer(establishment.GetID(), input.Name, input.PhoneNumber, input.Email)
	if err != nil {
		return err
	}

	_, err = c.customerRepository.Save(ctx, customer)
	if err != nil {
		return err
	}

	return nil
}
