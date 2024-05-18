package persistence

import (
	"context"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/mapper"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/util"
	"go.mongodb.org/mongo-driver/bson"
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
	professionalData, err := mapper.ToProfessionalData(entity)
	if err != nil {
		return nil, err
	}
	res, err := p.coll.InsertOne(ctx, professionalData)
	if err != nil {
		return nil, err
	}
	entity.SetID(res.InsertedID.(primitive.ObjectID).Hex())
	return entity, nil
}

func (p *ProfessionalMongoRepository) Get(ctx context.Context, id string) (*entity.Professional, error) {
	oid, err := util.GetObjectID(id)
	if err != nil {
		return nil, err
	}
	var professionalData data.ProfessionalData
	if err := p.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&professionalData); err != nil {
		return nil, err
	}
	professional, err := mapper.FromProfessionalData(&professionalData)
	if err != nil {
		return nil, err
	}
	return professional, nil
}
