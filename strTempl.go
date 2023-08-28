package templ

import "io"

func FromString(value string) templ {
	return strTempl{}
}

type strTempl struct {
}

func (t strTempl) Compose(templs ...templ) templ {

	return fnTempl{}
}

func (t strTempl) WriteTo(wr io.Writer) (int64, error) {
	return 0, nil
}
