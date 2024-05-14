package entity

import (
	"time"
)

type Scheduling struct {
	ID              string    `bson:"_id,omitempty" json:"id,omitempty"`
	ServiceID       string    `bson:"serviceID" json:"serviceID"`
	CustomerID      string    `bson:"customerID" json:"customerID"`
	ProfessionalID  string    `bson:"professionalID" json:"professionalID"`
	EstablishmentID string    `bson:"establishmentID" json:"establishmentID"`
	Date            time.Time `bson:"date" json:"date"`
	Time            time.Time `bson:"time" json:"time"`
}

func NewScheduling(service *Service, customer *Customer, professional *Professional, establishment *Establishment, date time.Time, time time.Time) (*Scheduling, error) {
	if err := professional.CanScheduleAtTheSpecifiedDateAndTime(date, time); err != nil {
		return nil, err
	}
	return &Scheduling{
		ServiceID:       service.ID,
		CustomerID:      customer.ID,
		ProfessionalID:  professional.ID,
		EstablishmentID: establishment.ID,
		Date:            date,
		Time:            time,
	}, nil
}
