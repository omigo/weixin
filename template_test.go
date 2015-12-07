package weixin

import "testing"

func TestSetIndustry(t *testing.T) {
	err := SetIndustry(1, 4)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestAddTemplate(t *testing.T) {
	templateId, err := AddTemplate("TM00015")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("templateId=%s", templateId)
}

func TestSendTemplateMsg(t *testing.T) {
	tm := &TemplateMsg{
		ToUser:     "odSYLjwG_jAujPH-XJfDseBjuggo",
		TemplateId: "peyNhQaS0BmNk7_ynbfS9aao323c_7Kz0p3qSknu-o0",
		URL:        "http://show.money",
		Data: TemplateData{
			First: KeywordPair{
				Value: "这是一个测试模板消息",
				Color: "#00ff00",
			},
			Keyword1: KeywordPair{
				Value: "这是一个测试模板消息",
				Color: "#0ff000",
			},
			Keyword2: KeywordPair{
				Value: "这是一个测试模板消息",
				Color: "#ff0000",
			},
			Keyword3: KeywordPair{
				Value: "这是一个测试模板消息",
				Color: "#0000ff",
			},
			Remark: KeywordPair{
				Value: "这是一个测试模板消息",
				Color: "#000ff0",
			},
		},
	}

	msgId, err := SendTemplateMsg(tm)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("msgId=%d", msgId)
}
