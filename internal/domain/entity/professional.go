package entity

import (
	"fmt"
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Professional struct {
	id              string       //`bson:"_id,omitempty" json:"id,omitempty"`
	accountID       string       //`bson:"accountID" json:"accountID"`
	establishmentID string       //`bson:"establishmentID" json:"establishmentID"`
	name            string       //`bson:"name" json:"name"`
	workPlan        *vo.WorkPlan //`bson:"workPlan" json:"workPlan"`
	active          bool         //`bson:"active" json:"active"`
	createdAt       time.Time    //`bson:"createdAt" json:"createdAt"`
}

func NewProfessional(accountID, establishmentID, name string) (*Professional, error) {
	workPlan, err := vo.DefaultWorkPlan()
	if err != nil {
		return nil, err
	}
	return &Professional{
		accountID:       accountID,
		establishmentID: establishmentID,
		name:            name,
		workPlan:        workPlan,
		active:          true,
		createdAt:       time.Now(),
	}, nil
}

func (p *Professional) CanScheduleAtTheSpecifiedDateAndTime(date, time time.Time) error {
	day := p.GetWorkPlan().GetDayFromWorkPlan(date)
	if day == nil {
		return fmt.Errorf("professional does not work on the chosen day")
	}

	// TODO: verify professional's break

	// check if the time is within range
	if day.GetStartTime().Before(day.GetEndTime()) {
		if (time.Equal(day.GetStartTime()) || time.After(day.GetStartTime())) && (time.Equal(day.GetEndTime()) || time.Before(day.GetEndTime())) {
			return nil
		}
	} else {
		// case where the interval crosses midnight
		if time.Equal(day.GetStartTime()) || time.After(day.GetStartTime()) || time.Equal(day.GetEndTime()) || time.Before(day.GetEndTime()) {
			return nil
		}
	}

	return fmt.Errorf("hours outside the professional's scheduling range")
}

func (a *Professional) SetID(id string) {
	a.id = id
}

func (a *Professional) GetID() string {
	return a.id
}

func (p *Professional) GetAccountID() string {
	return p.accountID
}

func (p *Professional) GetEstablishmentID() string {
	return p.establishmentID
}

func (p *Professional) GetName() string {
	return p.name
}

func (p *Professional) GetWorkPlan() *vo.WorkPlan {
	return p.workPlan
}
