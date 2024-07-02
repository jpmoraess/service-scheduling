package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/application/repository"
)

type FindCustomerOutputDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

type FindCustomerUseCase struct {
	customerRepository repository.CustomerRepository
}

func NewFindCustomerUseCase(customerRepository repository.CustomerRepository) *FindCustomerUseCase {
	return &FindCustomerUseCase{customerRepository: customerRepository}
}

func (f *FindCustomerUseCase) Execute(ctx context.Context, page, size int64) ([]*FindCustomerOutputDTO, error) {
	establishmentData, err := getEstablishmentData(ctx)
	if err != nil {
		return nil, err
	}

	customers, err := f.customerRepository.Find(ctx, establishmentData.ID(), page, size)
	if err != nil {
		return nil, err
	}

	output := make([]*FindCustomerOutputDTO, 0, len(customers))
	for _, customer := range customers {
		output = append(output, &FindCustomerOutputDTO{
			ID:          customer.ID(),
			Name:        customer.Name(),
			PhoneNumber: customer.PhoneNumber(),
			Email:       customer.Email(),
		})
	}

	return output, nil
}
