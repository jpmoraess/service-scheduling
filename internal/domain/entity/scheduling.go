package entity

import (
	"time"
)

type Scheduling struct {
	id              string    //`bson:"_id,omitempty" json:"id,omitempty"`
	serviceID       string    //`bson:"serviceID" json:"serviceID"`
	customerID      string    //`bson:"customerID" json:"customerID"`
	professionalID  string    //`bson:"professionalID" json:"professionalID"`
	establishmentID string    //`bson:"establishmentID" json:"establishmentID"`
	date            time.Time //`bson:"date" json:"date"`
	time            time.Time //`bson:"time" json:"time"`
	createdAt       time.Time //`bson:"createdAt" json:"createdAt"`
}

func NewScheduling(service *Service, customer *Customer, professional *Professional, establishment *Establishment, dateDate time.Time, timeTime time.Time) (*Scheduling, error) {
	if err := professional.CanScheduleAtTheSpecifiedDateAndTime(dateDate, timeTime); err != nil {
		return nil, err
	}
	return &Scheduling{
		serviceID:       service.ID(),
		customerID:      customer.ID(),
		professionalID:  professional.ID(),
		establishmentID: establishment.ID(),
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
