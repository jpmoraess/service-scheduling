package entity

import (
	"errors"
	"time"
)

type AccountType int

const (
	OwnerType AccountType = iota + 1
	ProfessionalType
)

func (at AccountType) String() string {
	switch at {
	case OwnerType:
		return "Owner"
	case ProfessionalType:
		return "Professional"
	default:
		return "Unknown"
	}
}

func ParseAccountTypeFromString(s string) (AccountType, error) {
	switch s {
	case "OwnerType":
		return OwnerType, nil
	case "ProfessionalType":
		return ProfessionalType, nil
	default:
		return -1, errors.New("invalid AccountType")
	}
}

func ParseAccountTypeFromInt(s int) (AccountType, error) {
	switch s {
	case 1:
		return OwnerType, nil
	case 2:
		return ProfessionalType, nil
	default:
		return -1, errors.New("invalid AccountType")
	}
}

type Account struct {
	ID                string      `bson:"_id,omitempty" json:"id,omitempty"`
	AccountType       AccountType `bson:"accountType" json:"accountType"`
	Name              string      `bson:"name" json:"name"`
	Email             string      `bson:"email" json:"email"`
	PhoneNumber       string      `bson:"phoneNumber" json:"phoneNumber"`
	EncryptedPassword string      `bson:"encryptedPassword" json:"-"`
	CreatedAt         time.Time   `bson:"createdAt" json:"createdAt"`
}

func NewAccount(accountType AccountType, name, email, phoneNumber, encryptedPassword string) (*Account, error) {
	return &Account{
		AccountType:       accountType,
		Name:              name,
		Email:             email,
		PhoneNumber:       phoneNumber,
		EncryptedPassword: encryptedPassword,
		CreatedAt:         time.Now(),
	}, nil
}
