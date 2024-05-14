package vo

import "time"

type Break struct {
	StartTime time.Time `bson:"startTime" json:"startTime"`
	EndTime   time.Time `bson:"endTime" json:"endTime"`
}

func NewBreak(startTime, endTime time.Time) (*Break, error) {
	// apply some validations
	return &Break{
		StartTime: startTime,
		EndTime:   endTime,
	}, nil
}
