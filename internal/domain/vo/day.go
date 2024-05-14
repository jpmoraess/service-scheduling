package vo

import (
	"time"
)

type Day struct {
	startTime time.Time //`bson:"startTime" json:"startTime"`
	endTime   time.Time //`bson:"endTime" json:"endTime"`
	aBreak    *Break    //`bson:"break" json:"break"`
}

func NewDay(startTime time.Time, endTime time.Time, aBreak *Break) (*Day, error) {
	// TODO: some validations
	return &Day{
		startTime: startTime,
		endTime:   endTime,
		aBreak:    aBreak,
	}, nil
}

func (d *Day) GetStartTime() time.Time {
	return d.startTime
}

func (d *Day) GetEndTime() time.Time {
	return d.endTime
}

func (d *Day) GetBreak() *Break {
	return d.aBreak
}
