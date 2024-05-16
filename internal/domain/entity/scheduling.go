package entity

import (
	"time"
)

type Scheduling struct {
	id              string
	serviceID       string
	customerID      string
	professionalID  string
	establishmentID string
	date            time.Time
	time            time.Time
	createdAt       time.Time
}

func NewScheduling(serviceID, customerID, professionalID, establishmentID string, dateDate, timeTime time.Time) (*Scheduling, error) {
	return &Scheduling{
		serviceID:       serviceID,
		customerID:      customerID,
		professionalID:  professionalID,
		establishmentID: establishmentID,
		date:            dateDate,
		time:            timeTime,
		createdAt:       time.Now(),
	}, nil
}

func RestoreScheduling(id, serviceID, customerID, professionalID, establishmentID string, dateDate, timeTime, createdAt time.Time) (*Scheduling, error) {
	return &Scheduling{
		id:              id,
		serviceID:       serviceID,
		customerID:      customerID,
		professionalID:  professionalID,
		establishmentID: establishmentID,
		date:            dateDate,
		time:            timeTime,
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

func (s *Scheduling) Date() time.Time {
	return s.date
}

func (s *Scheduling) Time() time.Time {
	return s.time
}

func (s *Scheduling) CreatedAt() time.Time {
	return s.time
}
