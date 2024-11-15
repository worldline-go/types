package convert

import "time"

type (
	Option     optioner
	OptionTime optioner
)

// ///////////////////////////////////////////////////////////////////////////
// interface holder

type optioner interface {
	Apply(opt *options)
}

type optionHandler struct {
	apply func(*options)
}

func (o *optionHandler) Apply(opt *options) {
	o.apply(opt)
}

func newOptionHandler(apply func(*options)) *optionHandler {
	return &optionHandler{apply: apply}
}

// ///////////////////////////////////////////////////////////////////////////
// main options

type options struct {
	TimeFormat string
}

func apply[T optioner](opts []T) *options {
	o := &options{}

	for _, opt := range opts {
		opt.Apply(o)
	}

	if o.TimeFormat == "" {
		o.TimeFormat = time.RFC3339
	}

	return o
}

// ///////////////////////////////////////////////////////////////////////////
// funcs of OptionTime

func WithTimeFormat(format string) OptionTime {
	return newOptionHandler(func(o *options) {
		o.TimeFormat = format
	})
}
