package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccountData struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	AccountType       int                `bson:"accountType"`
	Name              string             `bson:"name"`
	Email             string             `bson:"email"`
	PhoneNumber       string             `bson:"phoneNumber"`
	EncryptedPassword string             `bson:"encryptedPassword"`
}
