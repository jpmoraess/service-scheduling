package entity

import (
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Account struct {
	id                string
	accountType       vo.AccountType
	name              string
	email             string
	phoneNumber       string
	encryptedPassword string
	createdAt         time.Time
}

func NewAccount(accountType vo.AccountType, name, email, phoneNumber, password string) (*Account, error) {
	pw, err := vo.NewPassword(password)
	if err != nil {
		return nil, err
	}
	return &Account{
		accountType:       accountType,
		name:              name,
		email:             email,
		phoneNumber:       phoneNumber,
		encryptedPassword: pw.EncryptedPassword(),
		createdAt:         time.Now(),
	}, nil
}

func RestoreAccount(id string, accountType int, name, email, phoneNumber, encryptedPassword string, createdAt time.Time) (*Account, error) {
	accountTypeValue, err := vo.AccountTypeFromInt(accountType)
	if err != nil {
		return nil, err
	}
	return &Account{
		id:                id,
		accountType:       accountTypeValue,
		name:              name,
		email:             email,
		phoneNumber:       phoneNumber,
		encryptedPassword: encryptedPassword,
		createdAt:         createdAt,
	}, nil
}

func (a *Account) ResetPassword(newPassword string) error {
	pw, err := vo.NewPassword(newPassword)
	if err != nil {
		return err
	}
	a.encryptedPassword = pw.EncryptedPassword()
	return nil
}

func (a *Account) SetID(id string) {
	a.id = id
}

func (a *Account) ID() string {
	return a.id
}

func (a *Account) AccountType() vo.AccountType {
	return a.accountType
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

func (a *Account) CreatedAt() time.Time {
	return a.createdAt
}
