package vo

import "golang.org/x/crypto/bcrypt"

type Password struct {
	encryptedPassword string
}

func NewPassword(password string) (*Password, error) {
	// TODO: validate lenght, caracters...
	encryptedPasswod, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, err
	}
	return &Password{encryptedPassword: string(encryptedPasswod)}, nil
}

func (p *Password) EncryptedPassword() string {
	return p.encryptedPassword
}
