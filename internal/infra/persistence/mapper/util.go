package mapper

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectIDFromString(idStr string) (primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("error to parse string to ObjectID")
	}
	return oid, err
}
