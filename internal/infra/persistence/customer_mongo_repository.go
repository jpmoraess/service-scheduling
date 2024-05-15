package persistence

import (
	"context"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
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

type customerData struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	EstablishmentID string             `bson:"establishmentID"`
	Name            string             `bson:"name" json:"name"`
	PhoneNumber     string             `bson:"phoneNumber"`
	Email           string             `bson:"email"`
}

func toCustomerData(customer *entity.Customer) (*customerData, error) {
	return &customerData{
		EstablishmentID: customer.EstablishmentID(),
		Name:            customer.Name(),
		PhoneNumber:     customer.PhoneNumber(),
		Email:           customer.Email(),
	}, nil
}

func fromCustomerData(customerData *customerData) (*entity.Customer, error) {
	customer, err := entity.NewCustomer(customerData.EstablishmentID, customerData.Name, customerData.PhoneNumber, customerData.Email)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerMongoRepository) Save(ctx context.Context, entity *entity.Customer) (*entity.Customer, error) {
	customerData, err := toCustomerData(entity)
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
	var customerData customerData
	if err := c.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&customerData); err != nil {
		return nil, err
	}
	customer, err := fromCustomerData(&customerData)
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
	var customerData customerData
	if err := c.coll.FindOne(ctx, filter).Decode(&customerData); err != nil {
		return nil, err
	}
	customer, err := fromCustomerData(&customerData)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
