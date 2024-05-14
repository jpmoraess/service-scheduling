package persistence

import (
	"context"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceMongoRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewServiceMongoRepository(client *mongo.Client) *ServiceMongoRepository {
	return &ServiceMongoRepository{
		client: client,
		coll:   client.Database(configs.DATABASE_NAME).Collection("service"),
	}
}

func (s *ServiceMongoRepository) Save(ctx context.Context, entity *entity.Service) (*entity.Service, error) {
	res, err := s.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	entity.SetID(res.InsertedID.(primitive.ObjectID).Hex())
	return entity, nil
}

func (s *ServiceMongoRepository) Get(ctx context.Context, id string) (*entity.Service, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var service entity.Service
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&service); err != nil {
		return nil, err
	}
	return &service, nil
}

func (s *ServiceMongoRepository) FindByEstablishmentID(ctx context.Context, establishmentID string) ([]*entity.Service, error) {
	resp, err := s.coll.Find(ctx, bson.M{"establishmentID": establishmentID})
	if err != nil {
		return nil, err
	}
	var services []*entity.Service
	if err := resp.All(ctx, &services); err != nil {
		return nil, err
	}
	return services, nil
}
