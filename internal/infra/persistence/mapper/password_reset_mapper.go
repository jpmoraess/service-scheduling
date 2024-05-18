package mapper

import (
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/util"
)

func ToPasswordResetData(entity *entity.PasswordReset) (*data.PasswordResetData, error) {
	accountID, err := util.GetObjectID(entity.AccountID())
	if err != nil {
		return nil, err
	}
	return &data.PasswordResetData{
		AccountID: accountID,
		Token:     entity.Token(),
		Expiry:    entity.Expiry(),
	}, nil
}

func FromPasswordResetData(data *data.PasswordResetData) (*entity.PasswordReset, error) {
	passwordReset, err := entity.RestorePasswordReset(data.ID.Hex(), data.AccountID.Hex(), data.Token, data.Expiry)
	if err != nil {
		return nil, err
	}
	return passwordReset, nil
}
