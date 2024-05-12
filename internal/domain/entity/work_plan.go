package entity

import (
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type WorkPlan struct {
	ID             string  `bson:"_id,omitempty" json:"id,omitempty"`
	ProfessionalID string  `bson:"professionalID" json:"professionalID"`
	Monday         *vo.Day `bson:"monday" json:"monday"`
	Tuesday        *vo.Day `bson:"tuesday" json:"tuesday"`
	Wednesday      *vo.Day `bson:"wednesday" json:"wednesday"`
	Thursday       *vo.Day `bson:"thursday" json:"thursday"`
	Friday         *vo.Day `bson:"friday" json:"friday"`
	Saturday       *vo.Day `bson:"saturday" json:"saturday"`
	Sunday         *vo.Day `bson:"sunday" json:"sunday"`
}

func NewWorkPlan(monday, tuesday, wednesday, thursday, friday, saturday, sunday *vo.Day, professionalID string) (*WorkPlan, error) {
	return &WorkPlan{
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

func DefaultWorkPlan() (*WorkPlan, error) {
	day, err := vo.NewDay(time.Now(), time.Now(), nil)
	if err != nil {
		return nil, err
	}
	return &WorkPlan{
		Monday:    day,
		Tuesday:   day,
		Wednesday: day,
		Thursday:  day,
		Friday:    day,
		Saturday:  day,
	}, nil
}
