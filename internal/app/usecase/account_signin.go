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

type AccountSignin struct {
	accountRepository repository.AccountRepository
}

func NewAccountSignin(accountRepository repository.AccountRepository) *AccountSignin {
	return &AccountSignin{
		accountRepository: accountRepository,
	}
}

func (a *AccountSignin) Execute(ctx context.Context, input dto.AccountSigninInput) (*dto.AccountSigninOutput, error) {
	account, err := a.accountRepository.GetAccountByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if !isValidPassword(account.EncryptedPassword, input.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}
	accessToken := createAccessTokenFromAccount(account)
	return &dto.AccountSigninOutput{
		AccessToken: accessToken,
	}, nil
}

func createAccessTokenFromAccount(account *entity.Account) string {
	now := time.Now()
	expires := now.Add(time.Hour * 1)
	claims := jwt.MapClaims{
		"id":    account.ID,
		"email": account.Email,
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
