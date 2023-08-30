package templ

import "io"

type fnTempl struct {
	writer func(wr io.Writer) error
}

func (t fnTempl) Compose(templs ...Templ) Templ {
	return fnTempl{}
}

func (t fnTempl) Write(wr io.Writer) error {
	return t.writer(wr)
}
