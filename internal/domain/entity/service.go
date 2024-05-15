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
}

func NewService(establishmentID, name, description string, price *vo.Money, duration time.Duration, available bool) (*Service, error) {
	// TODO: validate price, duration....
	return &Service{
		establishmentID: establishmentID,
		name:            name,
		description:     description,
		price:           price,
		duration:        duration,
		available:       available,
	}, nil
}

func (a *Service) SetID(id string) {
	a.id = id
}

func (a *Service) ID() string {
	return a.id
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
