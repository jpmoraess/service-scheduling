package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SchedulingData struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ServiceID       string             `bson:"serviceID"`
	CustomerID      string             `bson:"customerID"`
	ProfessionalID  string             `bson:"professionalID"`
	EstablishmentID string             `bson:"establishmentID"`
	Date            time.Time          `bson:"date"`
	Time            time.Time          `bson:"time"`
}
