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

type CustomerData struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	EstablishmentID primitive.ObjectID `bson:"establishmentID"`
	Name            string             `bson:"name" json:"name"`
	PhoneNumber     string             `bson:"phoneNumber"`
	Email           string             `bson:"email"`
	CreatedAt       time.Time          `bson:"createdAt"`
}

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
	customerData, err := toCustomerData(entity)
	if err != nil {
		return nil, fmt.Errorf("failed to parse customer data from entity: %w", err)
	}
	resp, err := c.coll.InsertOne(ctx, customerData)
	if err != nil {
		return nil, fmt.Errorf("failed to insert customer data into database: %w", err)
	}
	entity.SetID(resp.InsertedID.(primitive.ObjectID).Hex())
	return entity, nil
}

func (c *CustomerMongoRepository) Get(ctx context.Context, id string) (*entity.Customer, error) {
	oid, err := util.GetObjectID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert customer id: %w", err)
	}
	var customerData CustomerData
	if err := c.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&customerData); err != nil {
		return nil, fmt.Errorf("failed to decode customer: %w", err)
	}
	customer, err := fromCustomerData(&customerData)
	if err != nil {
		return nil, fmt.Errorf("failed to map customer data to entity: %w", err)
	}
	return customer, nil
}

func (c *CustomerMongoRepository) GetByEstablishmentIDAndPhoneNumber(ctx context.Context, establishmentID string, phoneNumber string) (*entity.Customer, error) {
	filter := bson.M{
		"establishmentID": establishmentID,
		"phoneNumber":     phoneNumber,
	}
	var customerData CustomerData
	if err := c.coll.FindOne(ctx, filter).Decode(&customerData); err != nil {
		return nil, fmt.Errorf("failed to decode customer: %w", err)
	}
	customer, err := fromCustomerData(&customerData)
	if err != nil {
		return nil, fmt.Errorf("failed to map customer data to entity: %w", err)
	}
	return customer, nil
}

func (c *CustomerMongoRepository) Find(ctx context.Context, establishmentID string, page int64, size int64) ([]*entity.Customer, error) {
	establishmentOID, err := util.GetObjectID(establishmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert establishment id: %w", err)
	}
	opts := options.FindOptions{}
	opts.SetSkip((page - 1) * size)
	opts.SetLimit(size)
	cur, err := c.coll.Find(ctx, bson.M{"establishmentID": establishmentOID}, &opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find customers by establishment id: %w", err)
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {

		}
	}(cur, ctx)
	var customersData []*CustomerData
	if err := cur.All(ctx, &customersData); err != nil {
		return nil, fmt.Errorf("failed to decode customer: %w", err)
	}
	customers := make([]*entity.Customer, len(customersData))
	for _, customerData := range customersData {
		customer, err := fromCustomerData(customerData)
		if err != nil {
			return nil, fmt.Errorf("failed to map customer data to entity: %w", err)
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func toCustomerData(entity *entity.Customer) (*CustomerData, error) {
	establishmentID, err := util.GetObjectID(entity.EstablishmentID())
	if err != nil {
		return nil, err
	}

	return &CustomerData{
		EstablishmentID: establishmentID,
		Name:            entity.Name(),
		PhoneNumber:     entity.PhoneNumber(),
		Email:           entity.Email(),
		CreatedAt:       entity.CreatedAt(),
	}, nil
}

func fromCustomerData(data *CustomerData) (*entity.Customer, error) {
	customer, err := entity.RestoreCustomer(data.ID.Hex(), data.EstablishmentID.Hex(), data.Name, data.PhoneNumber, data.Email, data.CreatedAt)
	if err != nil {
		fmt.Println("error to restore customer from database", err)
		return nil, err
	}
	return customer, nil
}
