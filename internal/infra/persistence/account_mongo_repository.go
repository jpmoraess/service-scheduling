package persistence

import (
	"context"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
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
	res, err := a.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	entity.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return entity, nil
}

func (a *AccountMongoRepository) GetAccountByID(ctx context.Context, id string) (*entity.Account, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var account entity.Account
	if err := a.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&account); err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *AccountMongoRepository) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var account entity.Account
	if err := a.coll.FindOne(ctx, bson.M{"email": email}).Decode(&account); err != nil {
		return nil, err
	}
	return &account, nil
}
