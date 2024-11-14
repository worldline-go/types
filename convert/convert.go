package convert

import (
	"database/sql"
	"time"
)

func SQLNullToPtr[T any](v sql.Null[T]) *T {
	if v.Valid {
		return &v.V
	}

	return nil
}

func TimeFormatPtr(v *time.Time, opts ...OptionTime) *string {
	if v == nil || v.IsZero() {
		return nil
	}

	o := apply(opts)

	if o.TimeFormat == "" {
		o.TimeFormat = time.RFC3339
	}

	str := v.Format(o.TimeFormat)

	return &str
}
