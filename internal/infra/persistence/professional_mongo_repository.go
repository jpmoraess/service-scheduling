package persistence

import (
	"context"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfessionalMongoRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewProfessionalMongoRepository(client *mongo.Client) *ProfessionalMongoRepository {
	return &ProfessionalMongoRepository{
		client: client,
		coll:   client.Database(configs.DATABASE_NAME).Collection("professional"),
	}
}

func (p *ProfessionalMongoRepository) Save(ctx context.Context, entity *entity.Professional) (*entity.Professional, error) {
	res, err := p.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	entity.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return entity, nil
}
