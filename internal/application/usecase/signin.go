package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jpmoraess/service-scheduling/internal/application/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

type SigninInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninOutputDTO struct {
	AccessToken string `json:"accessToken"`
}

type SigninUseCase struct {
	accountRepository repository.AccountRepository
}

func NewSigninUseCase(accountRepository repository.AccountRepository) *SigninUseCase {
	return &SigninUseCase{
		accountRepository: accountRepository,
	}
}

func (a *SigninUseCase) Execute(ctx context.Context, input SigninInputDTO) (*SigninOutputDTO, error) {
	account, err := a.accountRepository.GetAccountByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	if !isValidPassword(account.EncryptedPassword(), input.Password) {
		return nil, errors.New("invalid credentials")
	}

	accessToken := createAccessTokenFromAccount(account)

	return &SigninOutputDTO{
		AccessToken: accessToken,
	}, nil
}

func createAccessTokenFromAccount(account *entity.Account) string {
	now := time.Now()
	expires := now.Add(time.Hour * 1)
	claims := jwt.MapClaims{
		"id":    account.ID(),
		"email": account.Email(),
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

func isValidPassword(encryptedPassword, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(pw)) == nil
}
