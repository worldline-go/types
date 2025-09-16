package types

import (
	"encoding/json"
	"testing"
)

func TestNull(t *testing.T) {
	type tt struct {
		Name  *string      `json:"name"`
		Value Null[string] `json:"value"`
	}

	type test struct {
		json  []byte
		check func(t *testing.T, n tt)
	}

	tests := []test{
		{
			json: []byte(`{"name": "example", "value": "test"}`),
			check: func(t *testing.T, n tt) {
				if n.Name == nil || *n.Name != "example" {
					t.Errorf("expected Name to be 'example', got %v", n.Name)
				}

				if !n.Value.Valid || n.Value.V != "test" {
					t.Errorf("expected Value to be valid with 'test', got %v", n.Value)
				}
			},
		},
		{
			json: []byte(`{"name": null, "value": null}`),
			check: func(t *testing.T, n tt) {
				if n.Name != nil {
					t.Errorf("expected Name to be nil, got %v", n.Name)
				}

				if !n.Value.ParsedNull {
					t.Errorf("expected Value.ParsedNull to be true, got false")
				}
			},
		},
		{
			json: []byte(`{"name": "example"}`),
			check: func(t *testing.T, n tt) {
				if n.Name == nil || *n.Name != "example" {
					t.Errorf("expected Name to be 'example', got %v", n.Name)
				}

				if n.Value.Valid {
					t.Errorf("expected Value.Valid to be false, got true")
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(string(test.json), func(t *testing.T) {
			n := tt{}

			if err := json.Unmarshal(test.json, &n); err != nil {
				t.Fatal(err)
			}

			test.check(t, n)
		})
	}
}
