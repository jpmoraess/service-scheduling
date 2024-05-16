package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerData struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	EstablishmentID string             `bson:"establishmentID"`
	Name            string             `bson:"name" json:"name"`
	PhoneNumber     string             `bson:"phoneNumber"`
	Email           string             `bson:"email"`
	CreatedAt       time.Time          `bson:"createdAt"`
}
