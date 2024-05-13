package persistence

import (
	"context"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EstablishmentMongoRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewEstablishmentMongoRepository(client *mongo.Client) *EstablishmentMongoRepository {
	return &EstablishmentMongoRepository{
		client: client,
		coll:   client.Database(configs.DATABASE_NAME).Collection("establishment"),
	}
}

func (e *EstablishmentMongoRepository) Save(ctx context.Context, entity *entity.Establishment) (*entity.Establishment, error) {
	res, err := e.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	entity.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return entity, nil
}

func (e *EstablishmentMongoRepository) GetByAccountID(ctx context.Context, accountID string) (*entity.Establishment, error) {
	var establishment entity.Establishment
	if err := e.coll.FindOne(ctx, bson.M{"accountID": accountID}).Decode(&establishment); err != nil {
		return nil, err
	}
	return &establishment, nil
}
