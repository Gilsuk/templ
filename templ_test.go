package templ_test

import (
	"bytes"
	"fmt"
	"testing"

	"git.gilsuk.page/gilsuk/templ"
)

func TestByPass(t *testing.T) {
	template := "A1b2{}C3d4{} {}e6f7{}"
	err := isTheSame(templ.FromString(template), template)

	if err != nil {
		t.Errorf("%+v", err)
	}
}

func isTheSame(template templ.Templ, expect string) error {

	var res bytes.Buffer
	template.WriteTo(&res)

	bufTempl := bytes.NewBufferString(expect)
	lenBufTempl := bufTempl.Len()

	for i := 0; i < lenBufTempl; i++ {
		b, err := bufTempl.ReadByte()
		b2, _ := res.ReadByte()
		if err != nil {
			return err
		}
		if b != b2 {
			return fmt.Errorf("expect: %x, but %x at the index %d", b, b2, i)
		}
	}

	return nil
}
