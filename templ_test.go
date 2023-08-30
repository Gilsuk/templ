package templ_test

import (
	"bytes"
	"fmt"
	"strings"
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

func TestCompose(t *testing.T) {
	cases := []struct {
		template string
		value    string
		expect   string
	}{
		{template: "a1b2{}", value: "C3d4", expect: "A1b2C3d4"},
	}

	for i, testCase := range cases {
		err := isTheSame(templ.FromString(testCase.template).Compose(templ.FromString(testCase.value)), testCase.expect)

		if err != nil {
			t.Errorf("msg: %s, at case idx[%d]", err.Error(), i)
		}
	}
}

func isTheSame(template templ.Templ, expect string) error {

	var res bytes.Buffer
	template.Write(&res)

	bufTempl := strings.NewReader(expect)
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
