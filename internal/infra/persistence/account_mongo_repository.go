package persistence

import (
	"context"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountMongoRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewAccountMongoRepository(client *mongo.Client) *AccountMongoRepository {
	return &AccountMongoRepository{
		client: client,
		coll:   client.Database(configs.DATABASE_NAME).Collection("account"),
	}
}

type accountData struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	AccountType       int                `bson:"accountType"`
	Name              string             `bson:"name"`
	Email             string             `bson:"email"`
	PhoneNumber       string             `bson:"phoneNumber"`
	EncryptedPassword string             `bson:"encryptedPassword"`
}

func toAccountData(account *entity.Account) (*accountData, error) {
	return &accountData{
		AccountType:       1,
		Name:              account.Name(),
		Email:             account.Email(),
		PhoneNumber:       account.PhoneNumber(),
		EncryptedPassword: account.EncryptedPassword(),
	}, nil
}

func fromAccountData(accountData *accountData) (*entity.Account, error) {
	accountType, err := vo.ParseAccountTypeFromInt(accountData.AccountType)
	if err != nil {
		return nil, err
	}
	account, err := entity.NewAccount(accountType, accountData.Name, accountData.Email, accountData.PhoneNumber, accountData.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a *AccountMongoRepository) Save(ctx context.Context, entity *entity.Account) (*entity.Account, error) {
	accountData, err := toAccountData(entity)
	if err != nil {
		return nil, err
	}
	res, err := a.coll.InsertOne(ctx, accountData)
	if err != nil {
		return nil, err
	}
	entity.SetID(res.InsertedID.(primitive.ObjectID).Hex())
	return entity, nil
}

func (a *AccountMongoRepository) GetAccountByID(ctx context.Context, id string) (*entity.Account, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var accountData accountData
	if err := a.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&accountData); err != nil {
		return nil, err
	}
	account, err := fromAccountData(&accountData)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a *AccountMongoRepository) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var accountData accountData
	if err := a.coll.FindOne(ctx, bson.M{"email": email}).Decode(&accountData); err != nil {
		return nil, err
	}
	account, err := fromAccountData(&accountData)
	if err != nil {
		return nil, err
	}
	return account, nil
}
