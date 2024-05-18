package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetObjectID(id string) (primitive.ObjectID, error) {
	if id == "" {
		return primitive.NilObjectID, nil
	}
	return primitive.ObjectIDFromHex(id)
}
