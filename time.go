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

	t.Time, _ = v.Time, v.Valid

	return nil
}

// Value implements the [driver.Valuer] interface.
func (t Time) Value() (driver.Value, error) {
	return t.Time, nil
}

// /////////////////////////////////////////////////////////////////////////////

type NullTime struct {
	time.Time
	Valid bool
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}

	return []byte(`"` + t.Time.Format(time.RFC3339) + `"`), nil
}

func (t *NullTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)

	if s == "null" || s == "" {
		t.Time = time.Time{}

		return nil
	}

	for _, format := range timeFormats {
		tt, err := time.Parse(format, s)
		if err == nil {
			t.Time = tt
			t.Valid = true

			return nil
		}
	}

	return fmt.Errorf("invalid time format: %s", s)
}

func (t *NullTime) Parse(s string) error {
	for _, format := range timeFormats {
		tt, err := time.Parse(format, s)
		if err == nil {
			t.Time = tt
			t.Valid = true

			return nil
		}
	}

	return fmt.Errorf("invalid time format: %s", s)
}

// String returns the time in RFC3339 format.
func (t NullTime) String() string {
	return t.Time.Format(time.RFC3339)
}

// Scan implements the [Scanner] interface.
func (t *NullTime) Scan(value any) error {
	v := sql.NullTime{}
	if err := v.Scan(value); err != nil {
		return err
	}

	t.Time, t.Valid = v.Time, v.Valid

	return nil
}

// Value implements the [driver.Valuer] interface.
func (t NullTime) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}

	return t.Time, nil
}
