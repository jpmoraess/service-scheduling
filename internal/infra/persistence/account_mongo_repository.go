package persistence

import (
	"context"
	"fmt"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/mapper"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	var accountData data.AccountData
	if err := a.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&accountData); err != nil {
		return nil, fmt.Errorf("failed to decode account: %w", err)
	}
	account, err := mapper.FromAccountData(&accountData)
	if err != nil {
		return nil, fmt.Errorf("failed to map account data to entity: %w", err)
	}
	return account, nil
}

func (a *AccountMongoRepository) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var accountData data.AccountData
	if err := a.coll.FindOne(ctx, bson.M{"email": email}).Decode(&accountData); err != nil {
		return nil, fmt.Errorf("failed to decode account: %w", err)
	}
	account, err := mapper.FromAccountData(&accountData)
	if err != nil {
		return nil, fmt.Errorf("failed to map account data to entity: %w", err)
	}
	return account, nil
}
