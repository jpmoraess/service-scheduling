package mapper

import (
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
)

func ToAccountData(account *entity.Account) (*data.AccountData, error) {
	return &data.AccountData{
		AccountType:       1, //TODO: FIX FIX
		Name:              account.Name(),
		Email:             account.Email(),
		PhoneNumber:       account.PhoneNumber(),
		EncryptedPassword: account.EncryptedPassword(),
		CreatedAt:         account.CreatedAt(),
	}, nil
}

func FromAccountData(data *data.AccountData) (*entity.Account, error) {
	accountType, err := vo.ParseAccountTypeFromInt(data.AccountType)
	if err != nil {
		fmt.Println("error parse account type from int", err)
		return nil, err
	}
	account, err := entity.RestoreAccount(data.ID.Hex(), accountType, data.Name, data.Email, data.PhoneNumber, data.EncryptedPassword, data.CreatedAt)
	if err != nil {
		fmt.Println("error to restore account from database", err)
		return nil, err
	}
	return account, nil
}
