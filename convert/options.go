package convert

import "time"

type (
	OptionTime func(o *optionTime)
)

// ///////////////////////////////////////////////////////////////////////////

type defaulter interface {
	Default()
}

func apply[T any, O ~func(*T)](opts []O) *T {
	opt := new(T)
	for _, o := range opts {
		o(opt)
	}

	if d, ok := any(opt).(defaulter); ok {
		d.Default()
	}

	return opt
}

// ///////////////////////////////////////////////////////////////////////////
// funcs of OptionTime

type optionTime struct {
	TimeFormat string
}

func (o *optionTime) Default() {
	if o.TimeFormat == "" {
		o.TimeFormat = time.RFC3339
	}
}

func WithTimeFormat(format string) OptionTime {
	return func(o *optionTime) {
		o.TimeFormat = format
	}
}
