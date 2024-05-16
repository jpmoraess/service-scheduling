package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BreakData struct {
	StartTime time.Time `bson:"startTime"`
	EndTime   time.Time `bson:"endTime"`
}

type DayData struct {
	StartTime time.Time  `bson:"startTime"`
	EndTime   time.Time  `bson:"endTime"`
	Break     *BreakData `bson:"break"`
}

type WorkPlanData struct {
	Monday    *DayData `bson:"monday"`
	Tuesday   *DayData `bson:"tuesday"`
	Wednesday *DayData `bson:"wednesday"`
	Thursday  *DayData `bson:"thursday"`
	Friday    *DayData `bson:"friday"`
	Saturday  *DayData `bson:"saturday"`
	Sunday    *DayData `bson:"sunday"`
}

type ProfessionalData struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	AccountID       string             `bson:"accountID"`
	EstablishmentID string             `bson:"establishmentID"`
	Name            string             `bson:"name"`
	WorkPlan        *WorkPlanData      `bson:"workPlan"`
	Active          bool               `bson:"active"`
	CreatedAt       time.Time          `bson:"createdAt"`
}
