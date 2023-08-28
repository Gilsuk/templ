package templ

import "io"

type templ interface {
	Compose(...templ) templ
	WriteTo(io.Writer) (int64, error)
}
