package templ

import "io"

type fnTempl struct {
}

func (t fnTempl) Compose(templs ...templ) templ {

	return fnTempl{}
}

func (t fnTempl) WriteTo(wr io.Writer) (int64, error) {
	return 0, nil
}
