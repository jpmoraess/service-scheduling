package entity

import (
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Service struct {
	id              string
	establishmentID string
	name            string
	description     string
	price           *vo.Money
	duration        time.Duration
	available       bool
	createdAt       time.Time
}

func NewService(establishmentID, name, description string, price *vo.Money, duration time.Duration, available bool) (*Service, error) {
	return &Service{
		establishmentID: establishmentID,
		name:            name,
		description:     description,
		price:           price,
		duration:        duration,
		available:       available,
		createdAt:       time.Now(),
	}, nil
}

func RestoreService(id, establishmentID, name, description string, price float64, duration time.Duration, available bool, createdAt time.Time) (*Service, error) {
	priceValue := vo.NewMoney(price)
	return &Service{
		id:              id,
		establishmentID: establishmentID,
		name:            name,
		description:     description,
		price:           priceValue,
		duration:        duration,
		available:       available,
		createdAt:       createdAt,
	}, nil
}

func (s *Service) SetID(id string) {
	s.id = id
}

func (s *Service) ID() string {
	return s.id
}

func (s *Service) EstablishmentID() string {
	return s.establishmentID
}

func (s *Service) Name() string {
	return s.name
}

func (s *Service) Description() string {
	return s.description
}

func (s *Service) Price() *vo.Money {
	return s.price
}

func (s *Service) Duration() time.Duration {
	return s.duration
}

func (s *Service) Available() bool {
	return s.available
}

func (s *Service) CreatedAt() time.Time {
	return s.createdAt
}
