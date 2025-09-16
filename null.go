package types

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

type Null[T any] struct {
	sql.Null[T]
	// ParsedNull is a helper field to distinguish between a null value and an omitted field during JSON unmarshalling.
	//  - if the field is present in the JSON (even if it's null), ParsedNull will be true.
	ParsedNull bool `json:"-"`
}

// NewNull creates a Null[T] with the given value and Valid set to true.
func NewNull[T any](v T) Null[T] {
	return Null[T]{Null: sql.Null[T]{V: v, Valid: true}}
}

// NewNullWithValid creates a Null[T] with the given value and Valid set to the specified boolean.
func NewNullWithValid[T any](v T, valid bool) Null[T] {
	return Null[T]{Null: sql.Null[T]{V: v, Valid: valid}}
}

// NewNullFromPtr creates a Null[T] from a pointer to T.
//   - If the pointer is nil, it returns a Null[T] with Valid set to false and V set to the zero value of T.
func NewNullFromPtr[T any](v *T) Null[T] {
	if v == nil {
		var zero T

		return Null[T]{Null: sql.Null[T]{V: zero, Valid: false}}
	}

	return Null[T]{Null: sql.Null[T]{V: *v, Valid: true}}
}

// Ptr returns a pointer to the inner value if Valid is true, otherwise it returns nil.
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
		n.V, n.Valid, n.ParsedNull = *new(T), false, true

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

func (n *Null[T]) Scan(value any) error {
	if err := n.Null.Scan(value); err != nil {
		return err
	}

	if !n.Valid {
		n.ParsedNull = true
	}

	return nil
}
