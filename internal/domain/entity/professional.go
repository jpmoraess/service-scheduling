package entity

import (
	"fmt"
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Professional struct {
	ID              string       `bson:"_id,omitempty" json:"id,omitempty"`
	AccountID       string       `bson:"accountID" json:"accountID"`
	EstablishmentID string       `bson:"establishmentID" json:"establishmentID"`
	Name            string       `bson:"name" json:"name"`
	WorkPlan        *vo.WorkPlan `bson:"workPlan" json:"workPlan"`
	Active          bool         `bson:"active" json:"active"`
	CreatedAt       time.Time
}

func NewProfessional(accountID, establishmentID, name string) (*Professional, error) {
	workPlan, err := vo.DefaultWorkPlan()
	if err != nil {
		return nil, err
	}
	return &Professional{
		AccountID:       accountID,
		EstablishmentID: establishmentID,
		Name:            name,
		WorkPlan:        workPlan,
		Active:          true,
		CreatedAt:       time.Now(),
	}, nil
}

func (p *Professional) CanScheduleAtTheSpecifiedDateAndTime(date, time time.Time) error {
	day := p.WorkPlan.GetDayFromWorkPlan(date)
	if day == nil {
		return fmt.Errorf("professional does not work on the chosen day")
	}

	// TODO: verify professional's break

	// check if the time is within range
	if day.StartTime.Before(day.EndTime) {
		if (time.Equal(day.StartTime) || time.After(day.StartTime)) && (time.Equal(day.EndTime) || time.Before(day.EndTime)) {
			return nil
		}
	} else {
		// case where the interval crosses midnight
		if time.Equal(day.StartTime) || time.After(day.StartTime) || time.Equal(day.EndTime) || time.Before(day.EndTime) {
			return nil
		}
	}

	return fmt.Errorf("hours outside the professional's scheduling range")
}
