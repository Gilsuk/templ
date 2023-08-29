package templ

import "io"

func FromFile(path string) Templ {
	return rdTempl{}
}

type rdTempl struct {
}

func (t rdTempl) Compose(templs ...Templ) Templ {

	return fnTempl{}
}

func (t rdTempl) WriteTo(wr io.Writer) (int64, error) {
	return 0, nil
}
