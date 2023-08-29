package templ

import "io"

func FromString(value string) Templ {
	return strTempl{}
}

type strTempl struct {
}

func (t strTempl) Compose(templs ...Templ) Templ {

	return fnTempl{}
}

func (t strTempl) WriteTo(wr io.Writer) (int64, error) {
	return 0, nil
}
