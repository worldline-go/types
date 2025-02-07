package types

import (
	"database/sql"
	"database/sql/driver"
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
	return []byte(`"` + t.Time.Format(time.RFC3339) + `"`), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)

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

// Scan implements the [Scanner] interface.
func (t *Time) Scan(value any) error {
	v := sql.NullTime{}
	if err := v.Scan(value); err != nil {
		return err
	}

	if v.Valid {
		t.Time = v.Time

		return nil
	}

	return fmt.Errorf("cannot scan [%v] into Time", value)
}

// Value implements the [driver.Valuer] interface.
func (t Time) Value() (driver.Value, error) {
	return t.Time, nil
}
