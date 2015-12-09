package weixin

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestUploadTemporaryMaterial(t *testing.T) {
	img, err := os.Open("/Users/migo/Pictures/psb.jpeg")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	mediaId, createAt, err := UploadTemporaryMaterial(MediaTypeImage, img)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v %#v", mediaId, createAt)
}

func TestGetTemporaryMaterial(t *testing.T) {
	mediaId := "Oyn8IWz0DhriA2-7C7DeLPSYdBjRmF2N89rZ1ZL68frE2wmLmp99SXz3rRE0ZJo8"
	filename, content, err := GetTemporaryMaterial(mediaId)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = ioutil.WriteFile(filename, content, 0777)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("write file to %s", filename)
}
