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
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (p *ProfessionalMongoRepository) Find(ctx context.Context, establishmentID string, page int64, size int64) ([]*entity.Professional, error) {
	establishmentOID, err := util.GetObjectID(establishmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert establishment id: %w", err)
	}
	opts := options.FindOptions{}
	opts.SetSkip((page - 1) * size)
	opts.SetLimit(size)
	cur, err := p.coll.Find(ctx, bson.M{"establishmentID": establishmentOID}, &opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find professional by establishment id: %w", err)
	}
	defer cur.Close(ctx)
	var professionalsData []data.ProfessionalData
	if err := cur.All(ctx, &professionalsData); err != nil {
		return nil, fmt.Errorf("failed to decode professional: %w", err)
	}
	professionals := make([]*entity.Professional, 0, len(professionalsData))
	for _, professionalData := range professionalsData {
		professional, err := mapper.FromProfessionalData(&professionalData)
		if err != nil {
			return nil, fmt.Errorf("failed to map professionl data to entity: %w", err)
		}
		professionals = append(professionals, professional)
	}
	return professionals, nil
}
