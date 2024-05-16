package entity

import (
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Scheduling struct {
	id              string
	serviceID       string
	customerID      string
	professionalID  string
	establishmentID string
	date            *vo.Date
	time            *vo.Time
	createdAt       time.Time
}

func NewScheduling(serviceID, customerID, professionalID, establishmentID, dateStr, timeStr string) (*Scheduling, error) {
	date, err := vo.NewDate(dateStr)
	if err != nil {
		return nil, err
	}
	timeVal, err := vo.NewTime(timeStr)
	if err != nil {
		return nil, err
	}
	return &Scheduling{
		serviceID:       serviceID,
		customerID:      customerID,
		professionalID:  professionalID,
		establishmentID: establishmentID,
		date:            date,
		time:            timeVal,
		createdAt:       time.Now(),
	}, nil
}

func RestoreScheduling(id, serviceID, customerID, professionalID, establishmentID, dateStr, timeStr string, createdAt time.Time) (*Scheduling, error) {
	date, err := vo.NewDate(dateStr)
	if err != nil {
		return nil, err
	}
	timeVal, err := vo.NewTime(timeStr)
	if err != nil {
		return nil, err
	}
	return &Scheduling{
		id:              id,
		serviceID:       serviceID,
		customerID:      customerID,
		professionalID:  professionalID,
		establishmentID: establishmentID,
		date:            date,
		time:            timeVal,
		createdAt:       createdAt,
	}, nil
}

func (s *Scheduling) SetID(id string) {
	s.id = id
}

func (s *Scheduling) ID() string {
	return s.id
}

func (s *Scheduling) ServiceID() string {
	return s.serviceID
}

func (s *Scheduling) CustomerID() string {
	return s.customerID
}

func (s *Scheduling) ProfessionalID() string {
	return s.professionalID
}

func (s *Scheduling) EstablishmentID() string {
	return s.establishmentID
}

func (s *Scheduling) Date() *vo.Date {
	return s.date
}

func (s *Scheduling) Time() *vo.Time {
	return s.time
}

func (s *Scheduling) CreatedAt() time.Time {
	return s.createdAt
}
