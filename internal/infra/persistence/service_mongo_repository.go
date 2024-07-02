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

type ServiceData struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	EstablishmentID primitive.ObjectID `bson:"establishmentID"`
	Name            string             `bson:"name"`
	Description     string             `bson:"description"`
	Price           float64            `bson:"price"`
	Duration        time.Duration      `bson:"duration"`
	Available       bool               `bson:"available"`
	CreatedAt       time.Time          `bson:"createdAt"`
}

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
	serviceData, err := toServiceData(entity)
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
	var serviceData ServiceData
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&serviceData); err != nil {
		return nil, fmt.Errorf("failed to decode service: %w", err)
	}
	service, err := fromServiceData(&serviceData)
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
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {

		}
	}(cur, ctx)

	var serviceData []*ServiceData
	if err := cur.All(ctx, &serviceData); err != nil {
		return nil, fmt.Errorf("failed to decode services: %w", err)
	}

	services := make([]*entity.Service, len(serviceData))
	for _, service := range serviceData {
		service, err := fromServiceData(service)
		if err != nil {
			return nil, fmt.Errorf("failed to map service data to entity: %w", err)
		}
		services = append(services, service)
	}

	return services, nil
}

func toServiceData(service *entity.Service) (*ServiceData, error) {
	establishmentID, err := util.GetObjectID(service.EstablishmentID())
	if err != nil {
		return nil, err
	}
	return &ServiceData{
		EstablishmentID: establishmentID,
		Name:            service.Name(),
		Description:     service.Description(),
		Price:           service.Price().AmountFloat64(),
		Duration:        service.Duration(),
		Available:       service.Available(),
	}, nil
}

func fromServiceData(data *ServiceData) (*entity.Service, error) {
	service, err := entity.RestoreService(data.ID.Hex(), data.EstablishmentID.Hex(), data.Name, data.Description, data.Price, data.Duration, data.Available, data.CreatedAt)
	if err != nil {
		fmt.Println("error to restore service from database", err)
		return nil, err
	}
	return service, nil
}
