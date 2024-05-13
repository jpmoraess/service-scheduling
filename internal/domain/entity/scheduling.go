package entity

import "time"

type Scheduling struct {
	ID              string    `bson:"_id,omitempty" json:"id,omitempty"`
	ServiceID       string    `bson:"serviceID" json:"serviceID"`
	CustomerID      string    `bson:"customerID" json:"customerID"`
	ProfessionalID  string    `bson:"professionalID" json:"professionalID"`
	EstablishmentID string    `bson:"establishmentID" json:"establishmentID"`
	Date            time.Time `bson:"date" json:"date"`
	Time            time.Time `bson:"time" json:"time"`
}

func NewScheduling(serviceID, customerID, professionalID, establishmentID string, date, time time.Time) (*Scheduling, error) {
	return &Scheduling{
		ServiceID:       serviceID,
		CustomerID:      customerID,
		ProfessionalID:  professionalID,
		EstablishmentID: establishmentID,
		Date:            date,
		Time:            time,
	}, nil
}
