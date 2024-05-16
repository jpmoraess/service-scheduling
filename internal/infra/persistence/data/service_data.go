package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceData struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	EstablishmentID primitive.ObjectID `bson:"establishmentID"`
	Name            string             `bson:"name"`
	Description     string             `bson:"description"`
	Price           float64            `bson:"price"`
	Duration        time.Duration      `bson:"duration"`
	Available       bool               `bson:"available"`
	CreatedAt       time.Time          `bson:"createdAt"`
}
