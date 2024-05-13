package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

func getAuthData(ctx context.Context) (*entity.Account, error) {
	authData, ok := ctx.Value("account").(*entity.Account)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return authData, nil
}

func parseDateTime(value string, layout string) (time.Time, error) {
	dateTime, err := time.Parse(layout, value)
	if err != nil {
		return time.Now(), err
	}
	return dateTime, nil
}
