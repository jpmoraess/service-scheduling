package entity

import (
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Account struct {
	id                string         //`bson:"_id,omitempty" json:"id,omitempty"`
	accountType       vo.AccountType //`bson:"accountType" json:"accountType"`
	name              string         //`bson:"name" json:"name"`
	email             string         //`bson:"email" json:"email"`
	phoneNumber       string         //`bson:"phoneNumber" json:"phoneNumber"`
	encryptedPassword string         //`bson:"encryptedPassword" json:"-"`
	createdAt         time.Time      //`bson:"createdAt" json:"createdAt"`
}

func NewAccount(accountType vo.AccountType, name, email, phoneNumber, encryptedPassword string) (*Account, error) {
	return &Account{
		accountType:       accountType,
		name:              name,
		email:             email,
		phoneNumber:       phoneNumber,
		encryptedPassword: encryptedPassword,
		createdAt:         time.Now(),
	}, nil
}

func (a *Account) SetID(id string) {
	a.id = id
}

func (a *Account) ID() string {
	return a.id
}

func (a *Account) Email() string {
	return a.email
}

func (a *Account) EncryptedPassword() string {
	return a.encryptedPassword
}
