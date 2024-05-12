package vo

import (
	"time"
)

type Day struct {
	Start time.Time `bson:"start" json:"start"`
	End   time.Time `bson:"end" json:"end"`
	Break *Break    `bson:"break" json:"break"`
}

func NewDay(start time.Time, end time.Time, aBreak *Break) (*Day, error) {
	// TODO: some validations
	return &Day{
		Start: start,
		End:   end,
		Break: aBreak,
	}, nil
}
