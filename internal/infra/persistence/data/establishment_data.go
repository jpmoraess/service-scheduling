package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EstablishmentData struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AccountID string             `bson:"accountID"`
	Name      string             `bson:"name"`
	Slug      string             `bson:"slug"`
	CreatedAt time.Time          `bson:"createdAt"`
}
