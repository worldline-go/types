package types

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

var timeFormats = []string{
	time.RFC3339,
	time.RFC3339Nano,
	time.DateTime,
	time.DateOnly,
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05.000000",
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

// /////////////////////////////////

func NewTime(t time.Time) Time {
	return Time{Time: t}
}

func NewTimeNull(t time.Time) Null[Time] {
	return NewNull(Time{Time: t})
}

func NewTimeNullWithValid(t time.Time, valid bool) Null[Time] {
	return NewNullWithValid(Time{Time: t}, valid)
}

func NewTimeNullFromPtr(t *time.Time) Null[Time] {
	if t == nil {
		return NewNullFromPtr[Time](nil)
	}

	return NewNullFromPtr(&Time{Time: *t})
}
