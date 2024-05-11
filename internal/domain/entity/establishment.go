package entity

import (
	"time"
)

type Establishment struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	AccountID string    `bson:"accountID" json:"accountID"`
	Name      string    `bson:"name" json:"name"`
	Slug      string    `bson:"slug" json:"slug"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

func NewEstablishment(accountID, name, slug string) (*Establishment, error) {
	return &Establishment{
		AccountID: accountID,
		Name:      name,
		Slug:      slug,
		CreatedAt: time.Now(),
	}, nil
}
