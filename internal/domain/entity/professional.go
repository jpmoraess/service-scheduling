package entity

import (
	"fmt"
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Professional struct {
	id              string
	accountID       string
	establishmentID string
	name            string
	workPlan        *vo.WorkPlan
	active          bool
	createdAt       time.Time
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

func RestoreProfessional(id, accountID, establishmentID, name string, workPlan *vo.WorkPlan, createdAt time.Time) (*Professional, error) {
	return &Professional{
		id:              id,
		accountID:       accountID,
		establishmentID: establishmentID,
		name:            name,
		workPlan:        workPlan,
		active:          true,
		createdAt:       createdAt,
	}, nil
}

func (p *Professional) CanScheduleAtTheSpecifiedDateAndTime(date, time time.Time) error {
	day := p.WorkPlan().GetDayFromWorkPlan(date)
	if day == nil {
		return fmt.Errorf("professional does not work on the chosen day")
	}

	// TODO: verify professional's break

	// check if the time is within range
	if day.StartTime().Before(day.EndTime()) {
		if (time.Equal(day.StartTime()) || time.After(day.StartTime())) && (time.Equal(day.EndTime()) || time.Before(day.EndTime())) {
			return nil
		}
	} else {
		// case where the interval crosses midnight
		if time.Equal(day.StartTime()) || time.After(day.StartTime()) || time.Equal(day.EndTime()) || time.Before(day.EndTime()) {
			return nil
		}
	}

	return fmt.Errorf("hours outside the professional's scheduling range")
}

func (a *Professional) SetID(id string) {
	a.id = id
}

func (a *Professional) ID() string {
	return a.id
}

func (p *Professional) AccountID() string {
	return p.accountID
}

func (p *Professional) EstablishmentID() string {
	return p.establishmentID
}

func (p *Professional) Name() string {
	return p.name
}

func (p *Professional) WorkPlan() *vo.WorkPlan {
	return p.workPlan
}

func (p *Professional) Active() bool {
	return p.active
}

func (p *Professional) CreatedAt() time.Time {
	return p.createdAt
}
