package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var timeFormats = []string{
	time.RFC3339,     // RFC3339
	time.RFC3339Nano, // RFC3339 with nanoseconds
	time.DateTime,    // DateTime
	time.DateOnly,    // Date only
}

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format(time.RFC3339))
}

func (t *Time) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)

	if s == "null" || s == "" {
		t.Time = time.Time{}

		return nil
	}

	for _, format := range timeFormats {
		tt, err := time.Parse(format, s)
		if err == nil {
			t.Time = tt

			return nil
		}
	}

	return fmt.Errorf("invalid time format: %s", s)
}

func (t *Time) Parse(s string) error {
	for _, format := range timeFormats {
		tt, err := time.Parse(format, s)
		if err == nil {
			t.Time = tt

			return nil
		}
	}

	return fmt.Errorf("invalid time format: %s", s)
}

// String returns the time in RFC3339 format.
func (t Time) String() string {
	return t.Time.Format(time.RFC3339)
}
