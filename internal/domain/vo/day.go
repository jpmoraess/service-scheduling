package vo

import (
	"time"
)

type Day struct {
	startTime time.Time
	endTime   time.Time
	aBreak    *Break
}

func NewDay(startTime time.Time, endTime time.Time, aBreak *Break) (*Day, error) {
	// TODO: some validations
	return &Day{
		startTime: startTime,
		endTime:   endTime,
		aBreak:    aBreak,
	}, nil
}

func (d *Day) StartTime() time.Time {
	return d.startTime
}

func (d *Day) EndTime() time.Time {
	return d.endTime
}

func (d *Day) Break() *Break {
	return d.aBreak
}
