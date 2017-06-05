package weixin

import (
	"io"
	"os"
	"testing"
)

func TestWxAcode(t *testing.T) {
	body, err := Acode("pages/index/index")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create("test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	defer body.Close()

	n, err := io.Copy(f, body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(n)
}
