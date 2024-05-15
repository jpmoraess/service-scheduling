package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
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

type serviceData struct {
	ID              string        `bson:"_id,omitempty"`
	EstablishmentID string        `bson:"establishmentID"`
	Name            string        `bson:"name"`
	Description     string        `bson:"description"`
	Price           float64       `bson:"price"`
	Duration        time.Duration `bson:"duration"`
	Available       bool          `bson:"available"`
}

func toServiceData(service *entity.Service) (*serviceData, error) {
	return &serviceData{
		EstablishmentID: service.EstablishmentID(),
		Name:            service.Name(),
		Description:     service.Description(),
		Price:           service.Price().AmountFloat64(),
		Duration:        service.Duration(),
		Available:       service.Available(),
	}, nil
}

func fromServiceData(data *serviceData) (*entity.Service, error) {
	service, err := entity.NewService(data.EstablishmentID, data.Name, data.Description, vo.NewMoney(data.Price), data.Duration, data.Available)
	if err != nil {
		fmt.Println("error to create new service", err)
		return nil, err
	}
	return service, nil
}

func (s *ServiceMongoRepository) Save(ctx context.Context, entity *entity.Service) (*entity.Service, error) {
	serviceData, err := toServiceData(entity)
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
	var serviceData serviceData
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&serviceData); err != nil {
		return nil, err
	}
	service, err := fromServiceData(&serviceData)
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
	var serviceData []*serviceData
	if err := resp.All(ctx, &serviceData); err != nil {
		return nil, err
	}
	var services []*entity.Service
	for _, data := range serviceData {
		service, err := fromServiceData(data)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}
