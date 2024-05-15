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

// TODO: remover dessa classe
type breakData struct {
	StartTime time.Time `bson:"startTime"`
	EndTime   time.Time `bson:"endTime"`
}

type dayData struct {
	StartTime time.Time  `bson:"startTime"`
	EndTime   time.Time  `bson:"endTime"`
	ABreak    *breakData `bson:"aBreak"`
}

type workPlanData struct {
	Monday    *dayData `bson:"monday"`
	Tuesday   *dayData `bson:"tuesday"`
	Wednesday *dayData `bson:"wednesday"`
	Thursday  *dayData `bson:"thursday"`
	Friday    *dayData `bson:"friday"`
	Saturday  *dayData `bson:"saturday"`
	Sunday    *dayData `bson:"sunday"`
}

type professionalData struct {
	ID              string        `bson:"_id,omitempty"`
	AccountID       string        `bson:"accountID"`
	EstablishmentID string        `bson:"establishmentID"`
	Name            string        `bson:"name"`
	WorkPlan        *workPlanData `bson:"workPlan"`
	Active          bool          `bson:"active"`
}

func toProfessionalData(professional *entity.Professional) (*professionalData, error) {
	// TODO: finalizar mapeamento e remover dessa classe

	//workPlan := professional.WorkPlan()

	workPlanData := &workPlanData{
		Monday:    &dayData{},
		Tuesday:   &dayData{},
		Wednesday: &dayData{},
		Thursday:  &dayData{},
		Friday:    &dayData{},
		Saturday:  &dayData{},
		Sunday:    &dayData{},
	}

	return &professionalData{
		ID:              professional.ID(),
		AccountID:       professional.AccountID(),
		EstablishmentID: professional.EstablishmentID(),
		Name:            professional.Name(),
		WorkPlan:        workPlanData,
		Active:          professional.Active(),
	}, nil
}

func fromProfessionalData(data *professionalData) (*entity.Professional, error) {
	professional, err := entity.NewProfessional(data.AccountID, data.EstablishmentID, data.Name)
	if err != nil {
		return nil, err
	}
	return professional, nil
}

func (p *ProfessionalMongoRepository) Save(ctx context.Context, entity *entity.Professional) (*entity.Professional, error) {
	professionalData, err := toProfessionalData(entity)
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
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var professionalData professionalData
	if err := p.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&professionalData); err != nil {
		return nil, err
	}
	professional, err := fromProfessionalData(&professionalData)
	if err != nil {
		return nil, err
	}
	return professional, nil
}
