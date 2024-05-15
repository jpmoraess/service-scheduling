package persistence

import (
	"context"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerMongoRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewCustomerMongoRepository(client *mongo.Client) *CustomerMongoRepository {
	return &CustomerMongoRepository{
		client: client,
		coll:   client.Database(configs.DATABASE_NAME).Collection("customer"),
	}
}

func (c *CustomerMongoRepository) Save(ctx context.Context, entity *entity.Customer) (*entity.Customer, error) {
	customerData, err := mapper.ToCustomerData(entity)
	if err != nil {
		return nil, err
	}
	resp, err := c.coll.InsertOne(ctx, customerData)
	if err != nil {
		return nil, err
	}
	entity.SetID(resp.InsertedID.(primitive.ObjectID).Hex())
	return entity, nil
}

func (c *CustomerMongoRepository) Get(ctx context.Context, id string) (*entity.Customer, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var customerData data.CustomerData
	if err := c.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&customerData); err != nil {
		return nil, err
	}
	customer, err := mapper.FromCustomerData(&customerData)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerMongoRepository) GetByEstablishmentIDAndPhoneNumber(ctx context.Context, establishmentID string, phoneNumber string) (*entity.Customer, error) {
	filter := bson.M{
		"establishmentID": establishmentID,
		"phoneNumber":     phoneNumber,
	}
	var customerData data.CustomerData
	if err := c.coll.FindOne(ctx, filter).Decode(&customerData); err != nil {
		return nil, err
	}
	customer, err := mapper.FromCustomerData(&customerData)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
