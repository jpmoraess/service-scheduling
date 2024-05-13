package entity

import (
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Account struct {
	ID                string         `bson:"_id,omitempty" json:"id,omitempty"`
	AccountType       vo.AccountType `bson:"accountType" json:"accountType"`
	Name              string         `bson:"name" json:"name"`
	Email             string         `bson:"email" json:"email"`
	PhoneNumber       string         `bson:"phoneNumber" json:"phoneNumber"`
	EncryptedPassword string         `bson:"encryptedPassword" json:"-"`
	CreatedAt         time.Time      `bson:"createdAt" json:"createdAt"`
}

func NewAccount(accountType vo.AccountType, name, email, phoneNumber, encryptedPassword string) (*Account, error) {
	return &Account{
		AccountType:       accountType,
		Name:              name,
		Email:             email,
		PhoneNumber:       phoneNumber,
		EncryptedPassword: encryptedPassword,
		CreatedAt:         time.Now(),
	}, nil
}
