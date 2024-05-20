package usecase

import (
	"context"

	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
)

type FindCustomer struct {
	customerRepository repository.CustomerRepository
}

func NewFindCustomer(customerRepository repository.CustomerRepository) *FindCustomer {
	return &FindCustomer{customerRepository: customerRepository}
}

func (f *FindCustomer) Execute(ctx context.Context, establishmentID string, page, size int64) ([]*dto.CustomerOutput, error) {
	customers, err := f.customerRepository.Find(ctx, establishmentID, page, size)
	if err != nil {
		return nil, err
	}
	output := make([]*dto.CustomerOutput, 0, len(customers))
	for _, customer := range customers {
		output = append(output, &dto.CustomerOutput{
			ID:          customer.ID(),
			Name:        customer.Name(),
			PhoneNumber: customer.PhoneNumber(),
			Email:       customer.Email(),
		})
	}
	return output, nil
}
