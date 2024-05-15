package persistence

import (
	"context"
	"fmt"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/mapper"
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
	serviceData, err := mapper.ToServiceData(entity)
	if err != nil {
		fmt.Println("error parse account data from entity", err)
		return nil, err
	}
	res, err := s.coll.InsertOne(ctx, serviceData)
	if err != nil {
		fmt.Println("error to insert service data into database", err)
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
	var serviceData data.ServiceData
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&serviceData); err != nil {
		return nil, err
	}
	service, err := mapper.FromServiceData(&serviceData)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (s *ServiceMongoRepository) FindByEstablishmentID(ctx context.Context, establishmentID string) ([]*entity.Service, error) {
	resp, err := s.coll.Find(ctx, bson.M{"establishmentID": establishmentID})
	if err != nil {
		return nil, err
	}
	var serviceData []*data.ServiceData
	if err := resp.All(ctx, &serviceData); err != nil {
		return nil, err
	}
	var services []*entity.Service
	for _, data := range serviceData {
		service, err := mapper.FromServiceData(data)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}
