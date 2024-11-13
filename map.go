package types

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Map map[string]interface{}

func (m *Map) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		// Parse the JSON data
		decoder := json.NewDecoder(bytes.NewReader(v))
		decoder.UseNumber()

		if err := decoder.Decode(m); err != nil {
			return err
		}

		return nil
	case string:
		// Parse the JSON string
		decoder := json.NewDecoder(bytes.NewReader([]byte(v)))
		decoder.UseNumber()

		if err := decoder.Decode(m); err != nil {
			return err
		}

		return nil
	case nil:
		*m = nil

		return nil
	default:
		return fmt.Errorf("%T, %w", value, ErrUnsupportedType)
	}
}

func (m Map) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}

	// Convert the map to JSON
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return b, nil
}
