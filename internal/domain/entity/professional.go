package entity

import (
	"time"
)

type Professional struct {
	id              string
	accountID       string
	establishmentID string
	name            string
	active          bool
	createdAt       time.Time
}

func NewProfessional(accountID, establishmentID, name string) (*Professional, error) {
	return &Professional{
		accountID:       accountID,
		establishmentID: establishmentID,
		name:            name,
		active:          true,
		createdAt:       time.Now(),
	}, nil
}

func RestoreProfessional(id, accountID, establishmentID, name string, createdAt time.Time) (*Professional, error) {
	return &Professional{
		id:              id,
		accountID:       accountID,
		establishmentID: establishmentID,
		name:            name,
		active:          true,
		createdAt:       createdAt,
	}, nil
}

//func (p *Professional) CanScheduleAtTheSpecifiedDateAndTime(date, time time.Time) error {
//	day := p.WorkPlan().GetDayFromWorkPlan(date)
//	if day == nil {
//		return fmt.Errorf("professional does not work on the chosen day")
//	}
//
//	if day.StartTime().Before(day.EndTime()) {
//		if (time.Equal(day.StartTime()) || time.After(day.StartTime())) && (time.Equal(day.EndTime()) || time.Before(day.EndTime())) {
//			return nil
//		}
//	} else {
//		if time.Equal(day.StartTime()) || time.After(day.StartTime()) || time.Equal(day.EndTime()) || time.Before(day.EndTime()) {
//			return nil
//		}
//	}
//
//	return fmt.Errorf("hours outside the professional's scheduling range")
//}

func (p *Professional) SetID(id string) {
	p.id = id
}

func (p *Professional) ID() string {
	return p.id
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

func (p *Professional) Active() bool {
	return p.active
}

func (p *Professional) CreatedAt() time.Time {
	return p.createdAt
}
