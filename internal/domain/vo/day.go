package vo

import (
	"time"
)

type Day struct {
	StartTime time.Time `bson:"startTime" json:"startTime"`
	EndTime   time.Time `bson:"endTime" json:"endTime"`
	Break     *Break    `bson:"break" json:"break"`
}

func NewDay(startTime time.Time, endTime time.Time, aBreak *Break) (*Day, error) {
	// TODO: some validations
	return &Day{
		StartTime: startTime,
		EndTime:   endTime,
		Break:     aBreak,
	}, nil
}
