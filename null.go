package types

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

type Null[T any] struct {
	sql.Null[T]
}

func NewNull[T any](v T) Null[T] {
	return Null[T]{Null: sql.Null[T]{V: v, Valid: true}}
}

func NewNullWithValid[T any](v T, valid bool) Null[T] {
	return Null[T]{Null: sql.Null[T]{V: v, Valid: valid}}
}

func NewNullFromPtr[T any](v *T) Null[T] {
	if v == nil {
		var zero T

		return Null[T]{Null: sql.Null[T]{V: zero, Valid: false}}
	}

	return Null[T]{Null: sql.Null[T]{V: *v, Valid: true}}
}

func (n Null[T]) Ptr() *T {
	if !n.Valid {
		return nil
	}

	return &n.V
}

// ValueOrZero returns the inner value if valid, otherwise zero.
func (n Null[T]) ValueOrZero() T {
	if !n.Valid {
		var zero T

		return zero
	}

	return n.V
}

func (n Null[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.V)
}

func (n *Null[T]) UnmarshalJSON(data []byte) error {
	if data == nil || bytes.Equal(data, []byte("null")) {
		n.V, n.Valid = *new(T), false

		return nil
	}

	// Parse the JSON data
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	if err := decoder.Decode(&n.V); err != nil {
		return err
	}

	n.Valid = true

	return nil
}
