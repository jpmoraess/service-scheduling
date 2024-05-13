package persistence

import (
	"context"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
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
	res, err := s.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	entity.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return entity, nil
}
