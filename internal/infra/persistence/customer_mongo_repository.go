package persistence

import (
	"context"

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
	oid, err := util.GetObjectID(id)
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

func (c *CustomerMongoRepository) Find(ctx context.Context, establishmentID string, page int64, size int64) ([]*entity.Customer, error) {
	establishmentOID, err := util.GetObjectID(establishmentID)
	if err != nil {
		return nil, err
	}
	opts := options.FindOptions{}
	opts.SetSkip((page - 1) * size)
	opts.SetLimit(size)
	resp, err := c.coll.Find(ctx, bson.M{"establishmentID": establishmentOID}, &opts)
	if err != nil {
		return nil, err
	}
	var customersData []data.CustomerData
	if err := resp.All(ctx, &customersData); err != nil {
		return nil, err
	}
	customers := make([]*entity.Customer, 0, len(customersData))
	for _, customerData := range customersData {
		customer, err := mapper.FromCustomerData(&customerData)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
