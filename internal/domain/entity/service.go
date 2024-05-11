package entity

import "time"

type Service struct {
	ID              string        `bson:"_id,omitempty" json:"id,omitempty"`
	EstablishmentID string        `bson:"establishmentID" json:"establishmentID"`
	Name            string        `bson:"name" json:"name"`
	Description     string        `bson:"description" json:"description"`
	Price           float64       `bson:"price" json:"price"`
	Duration        time.Duration `bson:"duration" json:"duration"`
	Available       bool          `bson:"available" json:"available"`
}

func NewService(establishmentID, name, description string, price float64, duration time.Duration) (*Service, error) {
	// TODO: validate price, duration....
	return &Service{
		EstablishmentID: establishmentID,
		Name:            name,
		Description:     description,
		Price:           price,
		Duration:        duration,
		Available:       true,
	}, nil
}
