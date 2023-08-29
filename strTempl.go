package templ

import (
	"io"
)

func FromString(value string) Templ {
	return strTempl{value: value}
}

type strTempl struct {
	value string
}

func (t strTempl) Compose(templs ...Templ) Templ {

	return fnTempl{}
}

func (t strTempl) Write(wr io.Writer) (int, error) {
	return wr.Write([]byte(t.value))
}
