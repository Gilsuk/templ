package templ

import "io"

type decoTempl struct {
	deco Templ
	args []Templ
}

func (t decoTempl) Compose(templs ...Templ) Templ {
	return decoTempl{deco: t, args: templs}
}

func (t decoTempl) Write(wr io.Writer) error {
	return nil
}
