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
	}, nil
}

func FromAccountData(accountData *data.AccountData) (*entity.Account, error) {
	accountType, err := vo.ParseAccountTypeFromInt(accountData.AccountType)
	if err != nil {
		fmt.Println("error parse account type from int", err)
		return nil, err
	}
	account, err := entity.RestoreAccount(accountData.ID.Hex(), accountType, accountData.Name, accountData.Email, accountData.PhoneNumber, accountData.EncryptedPassword)
	if err != nil {
		fmt.Println("error to restore account from database", err)
		return nil, err
	}
	return account, nil
}
