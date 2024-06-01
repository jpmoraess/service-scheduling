package usecase

import (
	"context"
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
)

func getAccountData(ctx context.Context) (*entity.Account, error) {
	accountData, ok := ctx.Value("account").(*entity.Account)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return accountData, nil
}

func getEstablishmentData(ctx context.Context) (*entity.Establishment, error) {
	establishmentData, ok := ctx.Value("establishment").(*entity.Establishment)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return establishmentData, nil
}
