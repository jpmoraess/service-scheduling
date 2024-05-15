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

func NewScheduling(serviceID string, customerID string, professionalID string, establishmentID string, dateDate time.Time, timeTime time.Time) (*Scheduling, error) {
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

func (a *Scheduling) SetID(id string) {
	a.id = id
}

func (a *Scheduling) ID() string {
	return a.id
}

func (a *Scheduling) ServiceID() string {
	return a.serviceID
}

func (a *Scheduling) CustomerID() string {
	return a.customerID
}

func (a *Scheduling) ProfessionalID() string {
	return a.professionalID
}

func (a *Scheduling) EstablishmentID() string {
	return a.establishmentID
}
