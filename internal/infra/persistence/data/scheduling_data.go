package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SchedulingData struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ServiceID       primitive.ObjectID `bson:"serviceID"`
	CustomerID      primitive.ObjectID `bson:"customerID"`
	ProfessionalID  primitive.ObjectID `bson:"professionalID"`
	EstablishmentID primitive.ObjectID `bson:"establishmentID"`
	Date            string             `bson:"date"`
	Time            string             `bson:"time"`
	CreatedAt       time.Time          `bson:"createdAt"`
}
