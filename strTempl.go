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

		for {
			data, err := rd.ReadBytes('{')

			if err == io.EOF {
				wr.Write(data)
				return nil
			} else if err != nil {
				return err
			}

			b, err := rd.ReadByte()

			if err == io.EOF {
				return nil
			} else if err != nil {
				return err
			}

			if b == '}' {
				wr.Write(data[:len(data)-2])
			} else {
				wr.Write(data)
				wr.Write([]byte{b})
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
