package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PasswordResetData struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AccountID primitive.ObjectID `bson:"accountID"`
	Token     string             `bson:"token"`
	Expiry    time.Time          `bson:"expiry"`
}
