package convert

import (
	"database/sql"
	"time"

	"github.com/worldline-go/types"
)

func NullToPtr[T any](v types.Null[T]) *T {
	if v.Valid {
		return &v.V
	}

	return nil
}

func SQLNullToPtr[T any](v sql.Null[T]) *T {
	if v.Valid {
		return &v.V
	}

	return nil
}

// ////////////////////////////////////////
// Time conversion functions

func TimeToStringPtr(v *time.Time, opts ...OptionTime) *string {
	if v == nil || v.IsZero() {
		return nil
	}

	str := v.Format(apply(opts).TimeFormat)

	return &str
}

func StringToTime(v string, opts ...OptionTime) (time.Time, error) {
	t, err := time.Parse(apply(opts).TimeFormat, v)

	return t, err
}

func StringPtrToTime(v *string, opts ...OptionTime) (time.Time, error) {
	if v == nil {
		return time.Time{}, nil
	}

	return StringToTime(*v, opts...)
}

func StringPtrToTimePtr(v *string, opts ...OptionTime) (*time.Time, error) {
	if v == nil {
		return nil, nil
	}

	t, err := StringToTime(*v, opts...)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ////////////////////////////////////////

func BytesToMap(v []byte) (types.Map[any], error) {
	m := make(types.Map[any])

	if err := m.Scan(v); err != nil {
		return nil, err
	}

	return m, nil
}
