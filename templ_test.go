package templ_test

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gilsuk/templ"
)

func TestByPassStrTempl(t *testing.T) {
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

func TestRdTempl(t *testing.T) {
	file, _ := os.OpenFile(filepath.Join(".", "testdata", "result.templ"), os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()

	templ.FromFile(".", "testdata", "outer.templ").
		Compose(templ.FromFile(".", "testdata", "content1.templ"), templ.FromFile(".", "testdata", "inner.templ").
			Compose(templ.FromFile(".", "testdata", "content2.templ"))).
		Write(file)

	expect, _ := os.Open(filepath.Join(".", "testdata", "expect.templ"))
	defer expect.Close()

	resRd := bufio.NewScanner(file)
	expRd := bufio.NewScanner(expect)

	for resRd.Scan() {
		if !expRd.Scan() {
			t.Errorf("%s is different with %s", filepath.Join(".", "testdata", "result.templ"), filepath.Join(".", "testdata", "expect.templ"))
		}

		if resRd.Text() != expRd.Text() {
			t.Errorf("%s is different with %s", filepath.Join(".", "testdata", "result.templ"), filepath.Join(".", "testdata", "expect.templ"))
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
