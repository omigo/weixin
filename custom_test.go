package weixin

import "testing"

func TestSendCustomMsg(t *testing.T) {
	Initialize(originId, appId, appSecret, token, encodingAESKey)

	ct := &CustText{
		Content: "hello",
	}

	err := SendCustomMsg("ozmLcjnM7vnrXmb3DimFLi0EOiY8", ct)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
