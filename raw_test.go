package types

import (
	"encoding/json"
	"testing"
)

func TestRawJSON(t *testing.T) {
	raw := RawJSON(`{"key":"value","price":123.65}`)
	m, err := raw.ToMap()
	if err != nil {
		t.Error(err)
	}

	if m["key"] != "value" {
		t.Error("invalid value")
	}

	if m["price"] != json.Number("123.65") {
		t.Error("invalid value")
	}

	b, err := raw.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	if string(b) != `{"key":"value","price":123.65}` {
		t.Error("invalid json")
	}
}

func TestRawJSON_Nested(t *testing.T) {
	type post struct {
		Title   string  `json:"title"`
		Details RawJSON `json:"details"`
	}

	p := post{
		Title:   "Title",
		Details: RawJSON(`{"key":"value","price":123.65}`),
	}

	b, err := json.Marshal(p)
	if err != nil {
		t.Error(err)
	}

	if string(b) != `{"title":"Title","details":{"key":"value","price":123.65}}` {
		t.Error("invalid json")
	}
}

func TestRawJson_Null(t *testing.T) {
	type post struct {
		Title   string  `json:"title"`
		Details RawJSON `json:"details"`
	}

	p := post{
		Title: "Title",
	}

	b, err := json.Marshal(p)
	if err != nil {
		t.Error(err)
	}

	if string(b) != `{"title":"Title","details":null}` {
		t.Error("invalid json")
	}
}
