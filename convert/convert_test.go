package convert

import (
	"reflect"
	"testing"
	"time"
)

func TestTimeFormatPtr(t *testing.T) {
	timeNow := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	type args struct {
		v    *time.Time
		opts []OptionTime
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantNil bool
	}{
		{
			name: "nil time",
			args: args{
				v:    nil,
				opts: nil,
			},
			wantNil: true,
		},
		{
			name: "zero time",
			args: args{
				v:    &time.Time{},
				opts: nil,
			},
			wantNil: true,
		},
		{
			name: "dateonly format",
			args: args{
				v:    &timeNow,
				opts: []OptionTime{WithTimeFormat(time.DateOnly)},
			},
			want: "2024-01-01",
		},
		{
			name: "default format",
			args: args{
				v:    &timeNow,
				opts: nil,
			},
			want: "2024-01-01T00:00:00Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TimeToStringPtr(tt.args.v, tt.args.opts...)

			if tt.wantNil {
				if got != nil {
					t.Errorf("TimeFormatPtr() = %v, want nil", got)
				}
				return
			}

			if got == nil {
				if !tt.wantNil {
					t.Errorf("TimeFormatPtr() = nil, want %v", tt.want)
				}

				return
			}

			if *got != tt.want {
				t.Errorf("TimeFormatPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPtrToZero(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		var vPtr *string
		v := PtrToZero(vPtr)
		if v != "" {
			t.Error("expected empty string")
		}
	})

	t.Run("empty time", func(t *testing.T) {
		var vPtr *time.Time
		v := PtrToZero(vPtr)
		if !v.IsZero() {
			t.Error("expected zero time")
		}
	})

	t.Run("empty int", func(t *testing.T) {
		var vPtr *int
		v := PtrToZero(vPtr)
		if v != 0 {
			t.Error("expected zero int")
		}
	})

	t.Run("non-empty string", func(t *testing.T) {
		vPtr := Ptr("hello")
		v := PtrToZero(vPtr)
		if v != "hello" {
			t.Error("expected 'hello'")
		}
	})
}
func TestRawTo_CustomStruct(t *testing.T) {
	type Custom struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Valid bool   `json:"valid"`
	}

	jsonData := []byte(`{"id":42,"name":"test","valid":true}`)

	got, err := RawTo[Custom](jsonData)
	if err != nil {
		t.Fatalf("RawTo() error = %v", err)
	}

	want := &Custom{ID: 42, Name: "test", Valid: true}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RawTo() = %+v, want %+v", got, want)
	}
}

func TestRawTo_InvalidJSON(t *testing.T) {
	type Custom struct {
		ID int `json:"id"`
	}
	invalidJSON := []byte(`{"id":`)

	_, err := RawTo[Custom](invalidJSON)
	if err == nil {
		t.Error("expected error for invalid JSON, got nil")
	}
}

func TestPtr(t *testing.T) {
	val := 123
	ptr := Ptr(val)
	if ptr == nil || *ptr != val {
		t.Errorf("Ptr() = %v, want %v", ptr, val)
	}
}

func TestIgnoreErr(t *testing.T) {
	val := "hello"
	got := IgnoreErr(val, nil)
	if got != val {
		t.Errorf("IgnoreErr() = %v, want %v", got, val)
	}
}
