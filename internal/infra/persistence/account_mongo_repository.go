package persistence

import (
	"context"
	"fmt"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/mapper"
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

func (a *AccountMongoRepository) Save(ctx context.Context, entity *entity.Account) (*entity.Account, error) {
	accountData, err := mapper.ToAccountData(entity)
	if err != nil {
		fmt.Println("error parse account data from entity", err)
		return nil, err
	}
	res, err := a.coll.InsertOne(ctx, accountData)
	if err != nil {
		fmt.Println("error to insert account data into database", err)
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
	var accountData data.AccountData
	if err := a.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&accountData); err != nil {
		return nil, err
	}
	account, err := mapper.FromAccountData(&accountData)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a *AccountMongoRepository) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var accountData data.AccountData
	if err := a.coll.FindOne(ctx, bson.M{"email": email}).Decode(&accountData); err != nil {
		return nil, err
	}
	account, err := mapper.FromAccountData(&accountData)
	if err != nil {
		return nil, err
	}
	return account, nil
}
