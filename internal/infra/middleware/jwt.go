package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jpmoraess/service-scheduling/internal/app/repository"
)

func JWTAuth(accountRepository repository.AccountRepository) fiber.Handler {
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
		// set the current authenticated account to the context
		accountID := claims["id"].(string)
		account, err := accountRepository.Get(c.Context(), accountID)
		if err != nil {
			return fmt.Errorf("unauthorized")
		}
		c.Context().SetUserValue("account", account)
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
