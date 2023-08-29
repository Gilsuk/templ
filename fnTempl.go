package templ

import "io"

type fnTempl struct {
}

func (t fnTempl) Compose(templs ...Templ) Templ {

	return fnTempl{}
}

func (t fnTempl) Write(wr io.Writer) (int, error) {
	return 0, nil
}
