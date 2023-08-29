package templ

import "io"

func FromFile(path string) Templ {
	return rdTempl{}
}

type rdTempl struct {
}

func (t rdTempl) Compose(templs ...Templ) Templ {
	return decoTempl{}
}

func (t rdTempl) Write(wr io.Writer) error {
	return nil
}
