package persistence

import (
	"context"
	"fmt"
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/util"
	"time"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BreakData struct {
	StartTime time.Time `bson:"startTime"`
	EndTime   time.Time `bson:"endTime"`
}

type DayData struct {
	StartTime time.Time  `bson:"startTime"`
	EndTime   time.Time  `bson:"endTime"`
	Break     *BreakData `bson:"break"`
}

type WorkPlanData struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ProfessionalID primitive.ObjectID `bson:"professionalID"`
	Monday         *DayData           `bson:"monday"`
	Tuesday        *DayData           `bson:"tuesday"`
	Wednesday      *DayData           `bson:"wednesday"`
	Thursday       *DayData           `bson:"thursday"`
	Friday         *DayData           `bson:"friday"`
	Saturday       *DayData           `bson:"saturday"`
	Sunday         *DayData           `bson:"sunday"`
}

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
	workPlanData, err := toWorkPlanData(entity)
	if err != nil {
		return nil, fmt.Errorf("failed to parse work plan data from entity: %w", err)
	}

	res, err := w.coll.InsertOne(ctx, workPlanData)
	if err != nil {
		return nil, fmt.Errorf("failed to insert work plan data into database: %w", err)
	}

	entity.SetID(res.InsertedID.(primitive.ObjectID).Hex())
	return entity, nil
}

func toWorkPlan(w *WorkPlanData) (*entity.WorkPlan, error) {
	monday, err := toDay(w.Monday)
	if err != nil {
		return nil, err
	}
	tuesday, err := toDay(w.Tuesday)
	if err != nil {
		return nil, err
	}
	wednesday, err := toDay(w.Wednesday)
	if err != nil {
		return nil, err
	}
	thursday, err := toDay(w.Thursday)
	if err != nil {
		return nil, err
	}
	friday, err := toDay(w.Friday)
	if err != nil {
		return nil, err
	}
	saturday, err := toDay(w.Saturday)
	if err != nil {
		return nil, err
	}
	sunday, err := toDay(w.Sunday)
	if err != nil {
		return nil, err
	}
	workPlan, err := entity.RestoreWorkPlan(w.ID.Hex(), w.ProfessionalID.Hex(), monday, tuesday, wednesday, thursday, friday, saturday, sunday)
	if err != nil {
		return nil, err
	}
	return workPlan, nil
}

func toWorkPlanData(workPlan *entity.WorkPlan) (*WorkPlanData, error) {
	professionalID, err := util.GetObjectID(workPlan.ProfessionalID())
	if err != nil {
		return nil, err
	}
	monday, err := toDayData(workPlan.Monday())
	if err != nil {
		return nil, err
	}
	tuesday, err := toDayData(workPlan.Tuesday())
	if err != nil {
		return nil, err
	}
	wednesday, err := toDayData(workPlan.Wednesday())
	if err != nil {
		return nil, err
	}
	thursday, err := toDayData(workPlan.Thursday())
	if err != nil {
		return nil, err
	}
	friday, err := toDayData(workPlan.Friday())
	if err != nil {
		return nil, err
	}
	saturday, err := toDayData(workPlan.Saturday())
	if err != nil {
		return nil, err
	}
	sunday, err := toDayData(workPlan.Sunday())
	if err != nil {
		return nil, err
	}

	return &WorkPlanData{
		ProfessionalID: professionalID,
		Monday:         monday,
		Tuesday:        tuesday,
		Wednesday:      wednesday,
		Thursday:       thursday,
		Friday:         friday,
		Saturday:       saturday,
		Sunday:         sunday,
	}, nil
}

func toDay(d *DayData) (*vo.Day, error) {
	aBreak, err := toBreak(d.Break)
	if err != nil {
		return nil, err
	}
	day, err := vo.NewDay(d.StartTime, d.EndTime, aBreak)
	if err != nil {
		return nil, err
	}
	return day, nil
}

func toDayData(day *vo.Day) (*DayData, error) {
	if day == nil {
		return nil, nil
	}
	breakVO := day.Break()
	var breakData *BreakData
	if breakVO != nil {
		b, err := toBreakData(breakVO)
		if err != nil {
			return nil, err
		}
		breakData = b
	}
	return &DayData{
		StartTime: day.StartTime(),
		EndTime:   day.EndTime(),
		Break:     breakData,
	}, nil
}

func toBreak(b *BreakData) (*vo.Break, error) {
	if b == nil {
		return nil, nil
	}
	return vo.NewBreak(b.StartTime, b.EndTime)
}

func toBreakData(aBreak *vo.Break) (*BreakData, error) {
	return &BreakData{
		StartTime: aBreak.StartTime(),
		EndTime:   aBreak.EndTime(),
	}, nil
}
