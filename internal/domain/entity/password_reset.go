package entity

import (
	"time"

	"github.com/google/uuid"
)

type PasswordReset struct {
	id        string
	accountID string
	token     string
	expiry    time.Time
}

func NewPasswordReset(accountID string) (*PasswordReset, error) {
	return &PasswordReset{
		accountID: accountID,
		token:     generateToken(),
		expiry:    time.Now().Add(time.Minute * 30),
	}, nil
}

func RestorePasswordReset(id, accountID, token string, expiry time.Time) (*PasswordReset, error) {

	return &PasswordReset{
		id:        id,
		accountID: accountID,
		token:     token,
		expiry:    expiry,
	}, nil
}

func generateToken() string {
	return uuid.New().String()
}

func (p *PasswordReset) IsExpiryTimeValid() bool {
	return time.Now().Before(p.expiry)
}

func (p *PasswordReset) ID() string {
	return p.id
}

func (p *PasswordReset) AccountID() string {
	return p.accountID
}

func (p *PasswordReset) Token() string {
	return p.token
}

func (p *PasswordReset) Expiry() time.Time {
	return p.expiry
}
