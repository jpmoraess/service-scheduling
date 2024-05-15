package vo

import (
	"fmt"
	"regexp"
	"strings"
)

type Slug struct {
	value string
}

func NewSlug(value string) (*Slug, error) {
	normalizedValue := normalize(value)
	if normalizedValue == "" {
		return nil, fmt.Errorf("invalid slug: %s", value)
	}
	return &Slug{value: normalizedValue}, nil
}

func normalize(value string) string {
	lowerValue := strings.ToLower(strings.TrimSpace(value))
	re := regexp.MustCompile(`[^a-z_]`)
	return re.ReplaceAllString(lowerValue, "")
}

func (s *Slug) Value() string {
	return s.value
}
