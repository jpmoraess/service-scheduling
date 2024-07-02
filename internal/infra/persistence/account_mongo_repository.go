package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccountData struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	AccountType       int                `bson:"accountType"`
	Name              string             `bson:"name"`
	Email             string             `bson:"email"`
	PhoneNumber       string             `bson:"phoneNumber"`
	EncryptedPassword string             `bson:"encryptedPassword"`
	CreatedAt         time.Time          `bson:"createdAt"`
}

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

func (a *AccountMongoRepository) Save(ctx context.Context, entity *entity.Account) (*entity.Account, error) {
	accountData, err := toAccountData(entity)
	if err != nil {
		return nil, fmt.Errorf("failed to parse account data from entity: %w", err)
	}

	oid, err := util.GetObjectID(entity.ID())
	if err != nil {
		return nil, fmt.Errorf("failed to convert account id: %w", err)
	}
	if oid != primitive.NilObjectID {
		accountData.ID = oid
	} else {
		accountData.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": accountData.ID}
	update := bson.M{"$set": accountData}
	opts := options.Update().SetUpsert(true)

	res, err := a.coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert account data into database: %w", err)
	}

	if res.UpsertedID != nil {
		entity.SetID(res.UpsertedID.(primitive.ObjectID).Hex())
	}
	return entity, nil
}

func (a *AccountMongoRepository) Get(ctx context.Context, id string) (*entity.Account, error) {
	oid, err := util.GetObjectID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert account id: %w", err)
	}
	var accountData AccountData
	if err := a.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&accountData); err != nil {
		return nil, fmt.Errorf("failed to decode account: %w", err)
	}
	account, err := fromAccountData(accountData)
	if err != nil {
		return nil, fmt.Errorf("failed to map account data to entity: %w", err)
	}
	return account, nil
}

func (a *AccountMongoRepository) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var accountData AccountData
	if err := a.coll.FindOne(ctx, bson.M{"email": email}).Decode(&accountData); err != nil {
		return nil, fmt.Errorf("failed to decode account: %w", err)
	}
	account, err := fromAccountData(accountData)
	if err != nil {
		return nil, fmt.Errorf("failed to map account data to entity: %w", err)
	}
	return account, nil
}

func toAccountData(account *entity.Account) (*AccountData, error) {
	return &AccountData{
		AccountType:       account.AccountType().Int(),
		Name:              account.Name(),
		Email:             account.Email(),
		PhoneNumber:       account.PhoneNumber(),
		EncryptedPassword: account.EncryptedPassword(),
		CreatedAt:         account.CreatedAt(),
	}, nil
}

func fromAccountData(data AccountData) (*entity.Account, error) {
	account, err := entity.RestoreAccount(data.ID.Hex(), data.AccountType, data.Name, data.Email, data.PhoneNumber, data.EncryptedPassword, data.CreatedAt)
	if err != nil {
		fmt.Println("error to restore account from database", err)
		return nil, err
	}
	return account, nil
}
