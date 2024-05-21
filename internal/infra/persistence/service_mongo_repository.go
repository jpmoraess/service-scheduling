package persistence

import (
	"context"
	"fmt"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/mapper"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/util"
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
		return nil, fmt.Errorf("failed to parse service data from entity: %w", err)
	}

	res, err := s.coll.InsertOne(ctx, serviceData)
	if err != nil {
		return nil, fmt.Errorf("failed to insert service data into database: %w", err)
	}

	entity.SetID(res.InsertedID.(primitive.ObjectID).Hex())
	return entity, nil
}

func (s *ServiceMongoRepository) Get(ctx context.Context, id string) (*entity.Service, error) {
	oid, err := util.GetObjectID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert service id: %w", err)
	}
	var serviceData data.ServiceData
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&serviceData); err != nil {
		return nil, fmt.Errorf("failed to decode service: %w", err)
	}
	service, err := mapper.FromServiceData(&serviceData)
	if err != nil {
		return nil, fmt.Errorf("failed to map service data to entity: %w", err)
	}
	return service, nil
}

func (s *ServiceMongoRepository) FindByEstablishmentID(ctx context.Context, establishmentID string) ([]*entity.Service, error) {
	establishmentOID, err := util.GetObjectID(establishmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert establishment ID: %w", err)
	}

	cur, err := s.coll.Find(ctx, bson.M{"establishmentID": establishmentOID})
	if err != nil {
		return nil, fmt.Errorf("failed to find services: %w", err)
	}
	defer cur.Close(ctx)

	var serviceData []*data.ServiceData
	if err := cur.All(ctx, &serviceData); err != nil {
		return nil, fmt.Errorf("failed to decode services: %w", err)
	}

	services := make([]*entity.Service, 0, len(serviceData))
	for _, data := range serviceData {
		service, err := mapper.FromServiceData(data)
		if err != nil {
			return nil, fmt.Errorf("failed to map service data to entity: %w", err)
		}
		services = append(services, service)
	}

	return services, nil
}
