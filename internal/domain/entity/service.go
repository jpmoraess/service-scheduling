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

func (a *Service) GetID() string {
	return a.id
}

func (s *Service) GetEstablishmentID() string {
	return s.establishmentID
}

func (s *Service) GetName() string {
	return s.name
}

func (s *Service) GetDescription() string {
	return s.description
}

func (s *Service) GetPrice() *vo.Money {
	return s.price
}

func (s *Service) GetDuration() time.Duration {
	return s.duration
}

func (s *Service) GetAvailable() bool {
	return s.available
}
