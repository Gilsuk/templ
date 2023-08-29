package templ

import "io"

type Templ interface {
	Compose(...Templ) Templ
	WriteTo(io.Writer) (int64, error)
}
