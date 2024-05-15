package vo

import "time"

type Break struct {
	startTime time.Time
	endTime   time.Time
}

func NewBreak(startTime, endTime time.Time) (*Break, error) {
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
