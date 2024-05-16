package usecase

import (
	"context"
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

func getAuthData(ctx context.Context) (*entity.Account, error) {
	authData, ok := ctx.Value("account").(*entity.Account)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return authData, nil
}
