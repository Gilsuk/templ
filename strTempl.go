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
	return decoTempl{}
}

func (t strTempl) Write(wr io.Writer) error {
	n, err := wr.Write([]byte(t.value))
	if n != len(t.value) {
		return err
	}
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}
