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

type SchedulingData struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ServiceID       primitive.ObjectID `bson:"serviceID"`
	CustomerID      primitive.ObjectID `bson:"customerID"`
	ProfessionalID  primitive.ObjectID `bson:"professionalID"`
	EstablishmentID primitive.ObjectID `bson:"establishmentID"`
	Date            string             `bson:"date"`
	Time            string             `bson:"time"`
	CreatedAt       time.Time          `bson:"createdAt"`
}

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
	oid, err := util.GetObjectID(id)
	var schedulingData SchedulingData
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&schedulingData); err != nil {
		return nil, err
	}
	scheduling, err := fromSchedulingData(&schedulingData)
	if err != nil {
		return nil, err
	}
	return scheduling, nil
}

func toSchedulingData(scheduling *entity.Scheduling) (*SchedulingData, error) {
	serviceID, err := util.GetObjectID(scheduling.ServiceID())
	if err != nil {
		return nil, err
	}
	customerID, err := util.GetObjectID(scheduling.CustomerID())
	if err != nil {
		return nil, err
	}
	professionalID, err := util.GetObjectID(scheduling.ProfessionalID())
	if err != nil {
		return nil, err
	}
	establishmentID, err := util.GetObjectID(scheduling.EstablishmentID())
	if err != nil {
		return nil, err
	}
	return &SchedulingData{
		ServiceID:       serviceID,
		CustomerID:      customerID,
		ProfessionalID:  professionalID,
		EstablishmentID: establishmentID,
		Date:            scheduling.Date().String(),
		Time:            scheduling.Time().String(),
		CreatedAt:       scheduling.CreatedAt(),
	}, nil
}

func fromSchedulingData(data *SchedulingData) (*entity.Scheduling, error) {
	scheduling, err := entity.RestoreScheduling(data.ID.Hex(), data.ServiceID.Hex(), data.CustomerID.Hex(), data.ProfessionalID.Hex(), data.EstablishmentID.Hex(), data.Date, data.Time, data.CreatedAt)
	if err != nil {
		fmt.Println("error to restore scheduling from database", err)
		return nil, err
	}
	return scheduling, nil
}
