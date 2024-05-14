package usecase

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

type Signin struct {
	accountRepository repository.AccountRepository
}

func NewSignin(accountRepository repository.AccountRepository) *Signin {
	return &Signin{
		accountRepository: accountRepository,
	}
}

func (a *Signin) Execute(ctx context.Context, input dto.SigninInput) (*dto.SigninOutput, error) {
	account, err := a.accountRepository.GetAccountByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if !isValidPassword(account.GetEncryptedPassword(), input.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}
	accessToken := createAccessTokenFromAccount(account)
	return &dto.SigninOutput{
		AccessToken: accessToken,
	}, nil
}

func createAccessTokenFromAccount(account *entity.Account) string {
	now := time.Now()
	expires := now.Add(time.Hour * 1)
	claims := jwt.MapClaims{
		"id":    account.GetID(),
		"email": account.GetEmail(),
		"iat":   jwt.NewNumericDate(now),
		"exp":   jwt.NewNumericDate(expires),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("failed to sign token with secret")
	}
	return tokenStr
}

func isValidPassword(encpw, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encpw), []byte(pw)) == nil
}
