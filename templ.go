package templ

import "io"

type Templ interface {
	Compose(...Templ) Templ
	Write(io.Writer) (int, error)
}
