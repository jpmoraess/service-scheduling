package persistence

import (
	"context"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkPlanMongoRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewWorkPlanMongoRepository(client *mongo.Client) *WorkPlanMongoRepository {
	return &WorkPlanMongoRepository{
		client: client,
		coll:   client.Database(configs.DATABASE_NAME).Collection("work_plan"),
	}
}

func (w *WorkPlanMongoRepository) Save(ctx context.Context, entity *entity.WorkPlan) (*entity.WorkPlan, error) {
	res, err := w.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	entity.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return entity, nil
}
