package middleware

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jpmoraess/service-scheduling/internal/application/repository"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

func JWTAuth(accountRepository repository.AccountRepository, establishmentRepository repository.EstablishmentRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if len(token) == 0 {
			return fmt.Errorf("unauthorized")
		}
		claims, err := validateJWTToken(token)
		if err != nil {
			return err
		}
		expires, err := claims.GetExpirationTime()
		if err != nil {
			return err
		}
		if time.Now().After(expires.Time) {
			return fmt.Errorf("token expired")
		}
		// START: set the current authenticated account and establishment data to the context
		accountID := claims["id"].(string)
		wg := sync.WaitGroup{}
		var account *entity.Account
		var establishment *entity.Establishment
		var accountErr, establishmentErr error

		wg.Add(2)
		go func() {
			defer wg.Done()
			account, accountErr = accountRepository.Get(c.Context(), accountID)
		}()
		go func() {
			defer wg.Done()
			establishment, establishmentErr = establishmentRepository.GetByAccountID(c.Context(), accountID)
		}()
		wg.Wait()

		if accountErr != nil || establishmentErr != nil {
			return fmt.Errorf("unauthorized")
		}
		c.Context().SetUserValue("account", account)
		c.Context().SetUserValue("establishment", establishment)
		// END: set the current authenticated account and establishment data to the context
		return c.Next()
	}
}

func validateJWTToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse JWT token", err)
		return nil, fmt.Errorf("unauthorized")
	}
	if !token.Valid {
		fmt.Println("invalid JWT token", err)
		return nil, fmt.Errorf("unauthorized")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}
