package templ

import (
	"bufio"
	"io"
	"strings"
)

func FromString(value string) Templ {
	return strTempl{value: value}
}

type strTempl struct {
	value string
}

func (t strTempl) Compose(templs ...Templ) Templ {
	return fnTempl{func(wr io.Writer) error {
		rd := bufio.NewReader(strings.NewReader(t.value))

		idx := 0
		for {
			data, err := rd.ReadBytes('}')

			if err != nil && err != io.EOF {
				return err
			}

			if len(data) < 2 {
				wr.Write(data)
				return nil
			}

			if data[len(data)-2] == '{' && len(templs) > idx {
				wr.Write(data[:len(data)-2])
				err := templs[idx].Write(wr)
				idx++
				if err != nil {
					return nil
				}
			} else {
				wr.Write(data)
			}

			if err == io.EOF {
				return nil
			}
		}
	}}
}

func (t strTempl) Write(wr io.Writer) error {
	n, err := wr.Write([]byte(t.value))
	if n != len(t.value) {
		return err
	}
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}
