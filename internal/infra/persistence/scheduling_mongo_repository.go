package persistence

import (
	"context"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/mapper"
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

func (s *SchedulingMongoRepository) Save(ctx context.Context, entity *entity.Scheduling) (*entity.Scheduling, error) {
	schedulingData, err := mapper.ToSchedulingData(entity)
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
	var schedulingData data.SchedulingData
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&schedulingData); err != nil {
		return nil, err
	}
	scheduling, err := mapper.FromSchedulingData(&schedulingData)
	if err != nil {
		return nil, err
	}
	return scheduling, nil
}
