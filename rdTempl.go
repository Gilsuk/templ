package templ

import "io"

func FromFile(path string) templ {
	return rdTempl{}
}

type rdTempl struct {
}

func (t rdTempl) Compose(templs ...templ) templ {

	return fnTempl{}
}

func (t rdTempl) WriteTo(wr io.Writer) (int64, error) {
	return 0, nil
}
