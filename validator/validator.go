package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

var V *validator.Validate

func init() {
	V = New()
}

func New() *validator.Validate {
	v := validator.New()

	// Register decimal.Decimal
	v.RegisterCustomTypeFunc(validateDecimal, decimal.Decimal{})
	v.RegisterCustomTypeFunc(validateNullDecimal, decimal.NullDecimal{})

	return v
}

func validateDecimal(field reflect.Value) interface{} {
	if dec, ok := field.Interface().(decimal.Decimal); ok {
		v, _ := dec.Float64()
		return v
	}

	return nil
}

func validateNullDecimal(field reflect.Value) interface{} {
	if dec, ok := field.Interface().(decimal.NullDecimal); ok {
		if dec.Valid {
			v, _ := dec.Decimal.Float64()
			return v
		}
	}

	return nil
}
