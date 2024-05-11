package entity

import "time"

type Professional struct {
	ID              string `bson:"_id,omitempty" json:"id,omitempty"`
	AccountID       string `bson:"accountID" json:"accountID"`
	EstablishmentID string `bson:"establishmentID" json:"establishmentID"`
	Name            string `bson:"name" json:"name"`
	Active          bool   `bson:"active" json:"active"`
	CreatedAt       time.Time
}

func NewProfessional(accountID, establishmentID, name string) (*Professional, error) {
	return &Professional{
		AccountID:       accountID,
		EstablishmentID: establishmentID,
		Name:            name,
		Active:          true,
		CreatedAt:       time.Now(),
	}, nil
}
