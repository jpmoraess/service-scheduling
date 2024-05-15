package entity

import (
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Account struct {
	id                string
	accountType       vo.AccountType
	name              string
	email             string
	phoneNumber       string
	encryptedPassword string
}

func NewAccount(accountType vo.AccountType, name, email, phoneNumber, encryptedPassword string) (*Account, error) {
	return &Account{
		accountType:       accountType,
		name:              name,
		email:             email,
		phoneNumber:       phoneNumber,
		encryptedPassword: encryptedPassword,
	}, nil
}

func (a *Account) SetID(id string) {
	a.id = id
}

func (a *Account) ID() string {
	return a.id
}

func (a *Account) Name() string {
	return a.name
}

func (a *Account) Email() string {
	return a.email
}

func (a *Account) PhoneNumber() string {
	return a.phoneNumber
}

func (a *Account) EncryptedPassword() string {
	return a.encryptedPassword
}
