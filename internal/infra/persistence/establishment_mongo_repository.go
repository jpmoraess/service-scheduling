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

type establishmentData struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AccountID string             `bson:"accountID"`
	Name      string             `bson:"name"`
	Slug      string             `bson:"slug"`
}

func toEstablishmentData(establishment *entity.Establishment) (*establishmentData, error) {
	return &establishmentData{
		AccountID: establishment.AccountID(),
		Name:      establishment.Name(),
		Slug:      establishment.Slug(),
	}, nil
}

func fromEstablishmentData(establishmentData *establishmentData) (*entity.Establishment, error) {
	establishment, err := entity.NewEstablishment(establishmentData.AccountID, establishmentData.Name, establishmentData.Slug)
	if err != nil {
		return nil, err
	}
	return establishment, nil
}

func (e *EstablishmentMongoRepository) Save(ctx context.Context, entity *entity.Establishment) (*entity.Establishment, error) {
	establishmentData, err := toEstablishmentData(entity)
	if err != nil {
		return nil, err
	}
	res, err := e.coll.InsertOne(ctx, establishmentData)
	if err != nil {
		return nil, err
	}
	entity.SetID(res.InsertedID.(primitive.ObjectID).Hex())
	return entity, nil
}

func (e *EstablishmentMongoRepository) Get(ctx context.Context, id string) (*entity.Establishment, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var establishmentData establishmentData
	if err := e.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&establishmentData); err != nil {
		return nil, err
	}
	establishment, err := fromEstablishmentData(&establishmentData)
	if err != nil {
		return nil, err
	}
	return establishment, nil
}

func (e *EstablishmentMongoRepository) GetBySlug(ctx context.Context, slug string) (*entity.Establishment, error) {
	var establishmentData establishmentData
	if err := e.coll.FindOne(ctx, bson.M{"slug": slug}).Decode(&establishmentData); err != nil {
		return nil, err
	}
	establishment, err := fromEstablishmentData(&establishmentData)
	if err != nil {
		return nil, err
	}
	return establishment, nil
}

func (e *EstablishmentMongoRepository) GetByAccountID(ctx context.Context, accountID string) (*entity.Establishment, error) {
	var establishmentData establishmentData
	if err := e.coll.FindOne(ctx, bson.M{"accountID": accountID}).Decode(&establishmentData); err != nil {
		return nil, err
	}
	establishment, err := fromEstablishmentData(&establishmentData)
	if err != nil {
		return nil, err
	}
	return establishment, nil
}
