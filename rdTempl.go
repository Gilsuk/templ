package templ

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

func FromFile(elem ...string) Templ {
	return rdTempl{filePath: filepath.Join(elem...)}
}

type rdTempl struct {
	filePath string
}

func (t rdTempl) Compose(templs ...Templ) Templ {
	return fnTempl{func(wr io.Writer) error {
		file, err := os.OpenFile(t.filePath, os.O_RDONLY, 0444)
		if err != nil {
			return err
		}
		defer file.Close()
		rd := bufio.NewReader(file)

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

func (t rdTempl) Write(wr io.Writer) error {
	file, err := os.OpenFile(t.filePath, os.O_RDONLY, 0444)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = bufio.NewReader(file).WriteTo(wr)
	return err
}
