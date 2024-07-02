package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/application/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

type CreateCustomerInputDTO struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

type CreateCustomerOutputDTO struct {
	Name string `json:"name"`
}

type CreateCustomerUseCase struct {
	customerRepository repository.CustomerRepository
}

func NewCreateCustomerUseCase(customerRepository repository.CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{customerRepository: customerRepository}
}

func (c *CreateCustomerUseCase) Execute(ctx context.Context, input CreateCustomerInputDTO) (*CreateCustomerOutputDTO, error) {
	establishmentData, err := getEstablishmentData(ctx)
	if err != nil {
		return nil, err
	}

	customer, err := entity.NewCustomer(establishmentData.ID(), input.Name, input.PhoneNumber, input.Email)
	if err != nil {
		return nil, err
	}

	customer, err = c.customerRepository.Save(ctx, customer)
	if err != nil {
		return nil, err
	}

	return &CreateCustomerOutputDTO{
		Name: customer.Name(),
	}, nil
}
