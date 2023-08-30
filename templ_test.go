package templ_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"git.gilsuk.page/gilsuk/templ"
)

func TestByPassStrTempl(t *testing.T) {
	template := "A1b2{}C3d4{} {}e6f7{}"
	err := isTheSame(templ.FromString(template), template)

	if err != nil {
		t.Errorf("%+v", err)
	}
}

func TestByPassRdTempl(t *testing.T) {
	expect := "{} header {} template"
	err := isTheSame(templ.FromFile(".", "testdata", "header.templ").Compose(templ.FromString("B1"), templ.FromString("5d")), expect)

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
		{template: "A1b2{}", value: "C3d4", expect: "A1b2C3d4"},
		{template: "A1b2}", value: "C3d4", expect: "A1b2}"},
		{template: "A1b{2}", value: "C3d4", expect: "A1b{2}"},
		{template: "{}", value: "C3d4", expect: "C3d4"},
		{template: "", value: "C3d4", expect: ""},
		{template: "}", value: "C3d4", expect: "}"},
		{template: "A{}c{}e}", value: "b", expect: "Abc{}e"},
	}

	for i, testCase := range cases {
		err := isTheSame(templ.FromString(testCase.template).Compose(templ.FromString(testCase.value)), testCase.expect)

		if err != nil {
			t.Errorf("msg: %s, at case idx[%d]", err.Error(), i)
		}
	}
}

func TestMultiCompose(t *testing.T) {
	cases := []struct {
		template string
		value1   string
		value2   string
		expect   string
	}{
		{template: "A1b2{}", value1: "C3d4", value2: "", expect: "A1b2C3d4"},
		{template: "A1b2{}{}", value1: "C3d4", value2: "E5", expect: "A1b2C3d4E5"},
		{template: "A1b2{}{}", value1: "", value2: "E5", expect: "A1b2E5"},
	}

	for i, testCase := range cases {
		err := isTheSame(templ.FromString(testCase.template).Compose(templ.FromString(testCase.value1), templ.FromString(testCase.value2)), testCase.expect)

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
