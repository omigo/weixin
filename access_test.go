package weixin

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arstd/log"
)

const (
	addr = ":3080"
)

func TestAccess(t *testing.T) {
	req, _ := http.NewRequest("GET", testURL, nil)

	w := httptest.NewRecorder()
	HandleAccess(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("http request error with status %d: %s", w.Code, w.Body)
		t.FailNow()
	}

	t.Logf("%s", w.Body)
}

func TestEventTemplateSendJobFinishHandler(t *testing.T) {
	// EventTemplateSendJobFinishHandler 模版消息发送结果通知事件
	EventTemplateSendJobFinishHandler = func(m *EventTemplateSendJobFinish) ReplyMsg {
		log.Debugf("%+v", m)

		// echo message
		ret := &ReplyText{
			ToUserName:   m.FromUserName,
			FromUserName: m.ToUserName,
			CreateTime:   m.CreateTime,
			Content:      fmt.Sprintf("Event=%s, MsgID=%d, Status=%s", m.Event, m.MsgID, m.Status),
		}

		log.Debugf("replay message: %+v", ret)
		return ret
	}

	body := `<xml>
       <ToUserName><![CDATA[gh_7f083739789a]]></ToUserName>
       <FromUserName><![CDATA[oia2TjuEGTNoeX76QEjQNrcURxG8]]></FromUserName>
       <CreateTime>1395658984</CreateTime>
       <MsgType><![CDATA[event]]></MsgType>
       <Event><![CDATA[TEMPLATESENDJOBFINISH]]></Event>
       <MsgID>200163840</MsgID>
       <Status><![CDATA[failed: system failed]]></Status>
       </xml>`
	req, _ := http.NewRequest("POST", testURL, bytes.NewBufferString(body))

	w := httptest.NewRecorder()
	HandleAccess(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("http request error with status %d: %s", w.Code, w.Body)
		t.FailNow()
	}

	t.Logf("%s", w.Body)
}
