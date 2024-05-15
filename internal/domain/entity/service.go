package entity

import (
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Service struct {
	id              string        //`bson:"_id,omitempty" json:"id,omitempty"`
	establishmentID string        //`bson:"establishmentID" json:"establishmentID"`
	name            string        //`bson:"name" json:"name"`
	description     string        //`bson:"description" json:"description"`
	price           *vo.Money     //`bson:"price" json:"price"`
	duration        time.Duration //`bson:"duration" json:"duration"`
	available       bool          //`bson:"available" json:"available"`
}

func NewService(establishmentID, name, description string, price *vo.Money, duration time.Duration) (*Service, error) {
	// TODO: validate price, duration....
	return &Service{
		establishmentID: establishmentID,
		name:            name,
		description:     description,
		price:           price,
		duration:        duration,
		available:       true,
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
