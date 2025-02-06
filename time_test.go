package types

import (
	"testing"
	"time"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name    string
		args    [][]byte
		wantErr bool
	}{
		{
			name: "Test Time UnmarshalJSON",
			args: [][]byte{
				[]byte(`2020-01-01`),
				[]byte(`"2020-01-01 00:00:00"`),
				[]byte(`2020-01-01T00:00:00Z`),
				[]byte(`2006-01-02T15:04:05Z`),
				[]byte(`2020-01-01T00:00:00.000Z`),
				[]byte(`2025-01-31T09:41:17Z`),
				[]byte(`2025-01-31T10:41:55+01:00`),
				[]byte(`2025-01-31T09:43:00.3Z`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tr NullTime
			for _, arg := range tt.args {
				if err := tr.UnmarshalJSON(arg); (err != nil) != tt.wantErr {
					t.Errorf("Time.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
