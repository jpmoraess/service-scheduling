package vo

import (
	"errors"
	"time"
)

type Time struct {
	value time.Time
}

func NewTime(timeStr string) (*Time, error) {
	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return nil, errors.New("invalid time format, expected HH:MM")
	}
	return &Time{value: parsedTime}, nil
}

func (t *Time) Value() time.Time {
	return t.value
}

func (t *Time) String() string {
	return t.value.String()
}
