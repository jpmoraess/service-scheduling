package entity

import (
	"time"
)

type Account struct {
	ID                string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string    `bson:"name" json:"name"`
	Email             string    `bson:"email" json:"email"`
	PhoneNumber       string    `bson:"phoneNumber" json:"phoneNumber"`
	EncryptedPassword string    `bson:"encryptedPassword" json:"-"`
	CreatedAt         time.Time `bson:"createdAt" json:"createdAt"`
}

func NewAccount(name, email, phoneNumber, encryptedPassword string) (*Account, error) {
	return &Account{
		Name:              name,
		Email:             email,
		PhoneNumber:       phoneNumber,
		EncryptedPassword: encryptedPassword,
		CreatedAt:         time.Now(),
	}, nil
}
