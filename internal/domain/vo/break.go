package vo

import "time"

type Break struct {
	startTime time.Time `bson:"startTime" json:"startTime"`
	endTime   time.Time `bson:"endTime" json:"endTime"`
}

func NewBreak(startTime, endTime time.Time) (*Break, error) {
	// apply some validations
	return &Break{
		startTime: startTime,
		endTime:   endTime,
	}, nil
}

func (b *Break) StartTime() time.Time {
	return b.startTime
}

func (b *Break) EndTime() time.Time {
	return b.endTime
}
