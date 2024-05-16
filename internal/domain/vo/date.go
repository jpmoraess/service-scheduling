package vo

import (
	"errors"
	"time"
)

type Date struct {
	value time.Time
}

func NewDate(dateStr string) (*Date, error) {
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, errors.New("invalid date format, expected YYYY-MM-DD")
	}
	return &Date{value: parsedDate}, nil
}

func (d *Date) Value() time.Time {
	return d.value
}

func (d *Date) String() string {
	return d.value.String()
}
