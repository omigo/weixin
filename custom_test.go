package weixin

import (
	"os"
	"testing"
)

func TestSendCustomMsg(t *testing.T) {
	ct := &CustText{
		Content: "我们都是好孩子",
	}

	err := SendCustomMsg("ozmLcjnM7vnrXmb3DimFLi0EOiY8", ct)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestAddCustom(t *testing.T) {
	err := AddCustom("shangxuejin@gmail.com", "migo", "87654321")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestUploadHeading(t *testing.T) {
	file, err := os.Open("custom_test.go")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = UploadHeading("shangxuejin@gmail.com", file)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestGetCustomList(t *testing.T) {
	accs, err := GetCustomList()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%v", accs)
}
