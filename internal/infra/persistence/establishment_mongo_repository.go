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
)

type EstablishmentData struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AccountID primitive.ObjectID `bson:"accountID"`
	Name      string             `bson:"name"`
	Slug      string             `bson:"slug"`
	CreatedAt time.Time          `bson:"createdAt"`
}

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
	oid, err := util.GetObjectID(id)
	if err != nil {
		return nil, err
	}
	var establishmentData EstablishmentData
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
	var establishmentData EstablishmentData
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
	oid, err := util.GetObjectID(accountID)
	if err != nil {
		return nil, err
	}
	var establishmentData EstablishmentData
	if err := e.coll.FindOne(ctx, bson.M{"accountID": oid}).Decode(&establishmentData); err != nil {
		return nil, err
	}
	establishment, err := fromEstablishmentData(&establishmentData)
	if err != nil {
		return nil, err
	}
	return establishment, nil
}

func toEstablishmentData(entity *entity.Establishment) (*EstablishmentData, error) {
	accountID, err := util.GetObjectID(entity.AccountID())
	if err != nil {
		return nil, err
	}
	return &EstablishmentData{
		AccountID: accountID,
		Name:      entity.Name(),
		Slug:      entity.Slug(),
		CreatedAt: entity.CreatedAt(),
	}, nil
}

func fromEstablishmentData(data *EstablishmentData) (*entity.Establishment, error) {
	establishment, err := entity.RestoreEstablishment(data.ID.Hex(), data.AccountID.Hex(), data.Name, data.Slug, data.CreatedAt)
	if err != nil {
		fmt.Println("error to restore establishment from database", err)
		return nil, err
	}
	return establishment, nil
}
