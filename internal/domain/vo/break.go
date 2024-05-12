package vo

import "time"

type Break struct {
	Start time.Time `bson:"start" json:"start"`
	End   time.Time `bson:"end" json:"end"`
}

func NewBreak(start, end time.Time) (*Break, error) {
	// apply some validations
	return &Break{
		Start: start,
		End:   end,
	}, nil
}
