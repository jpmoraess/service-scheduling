package persistence

import (
	"context"
	"time"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SchedulingMongoRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewSchedulingMongoRepository(client *mongo.Client) *SchedulingMongoRepository {
	return &SchedulingMongoRepository{
		client: client,
		coll:   client.Database(configs.DATABASE_NAME).Collection("scheduling"),
	}
}

type schedulingData struct {
	ID              string    `bson:"_id,omitempty"`
	ServiceID       string    `bson:"serviceID"`
	CustomerID      string    `bson:"customerID"`
	ProfessionalID  string    `bson:"professionalID"`
	EstablishmentID string    `bson:"establishmentID"`
	Date            time.Time `bson:"date"`
	Time            time.Time `bson:"time"`
}

func toSchedulingData(scheduling *entity.Scheduling) (*schedulingData, error) {
	return &schedulingData{
		ServiceID:       scheduling.ID(),
		CustomerID:      scheduling.CustomerID(),
		ProfessionalID:  scheduling.ProfessionalID(),
		EstablishmentID: scheduling.EstablishmentID(),
	}, nil
}

func fromSchedulingData(data *schedulingData) (*entity.Scheduling, error) {
	scheduling, err := entity.NewScheduling(data.ServiceID, data.CustomerID, data.ProfessionalID, data.EstablishmentID, data.Date, data.Time)
	if err != nil {
		return nil, err
	}
	return scheduling, nil
}

func (s *SchedulingMongoRepository) Save(ctx context.Context, entity *entity.Scheduling) (*entity.Scheduling, error) {
	schedulingData, err := toSchedulingData(entity)
	if err != nil {
		return nil, err
	}
	res, err := s.coll.InsertOne(ctx, schedulingData)
	if err != nil {
		return nil, err
	}
	entity.SetID(res.InsertedID.(primitive.ObjectID).Hex())
	return entity, nil
}

func (s *SchedulingMongoRepository) Get(ctx context.Context, id string) (*entity.Scheduling, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var schedulingData schedulingData
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&schedulingData); err != nil {
		return nil, err
	}
	scheduling, err := fromSchedulingData(&schedulingData)
	if err != nil {
		return nil, err
	}
	return scheduling, nil
}
