package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Map map[string]interface{}

func (m *Map) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		// Parse the JSON data
		err := json.Unmarshal(v, &m)
		if err != nil {
			return err
		}
		return nil
	case string:
		// Parse the JSON string
		err := json.Unmarshal([]byte(v), &m)
		if err != nil {
			return err
		}
		return nil
	case nil:
		*m = nil
		return nil
	default:
		return errors.New("unsupported type for Map scan")
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
