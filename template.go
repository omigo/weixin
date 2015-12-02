package weixin

import "fmt"

// 模板消息接口
const (
	TamplateSetIndustryURL     = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=%s"
	TemplateAddTemplateURL     = "https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=%s"
	TemplateSendTemplateMsgURL = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"
)

// SetIndustry 设置所属行业
/*
行业代码查询

主行业	副行业	代码
IT科技	互联网/电子商务	1
IT科技	IT软件与服务	2
IT科技	IT硬件与设备	3
IT科技	电子技术	4
IT科技	通信与运营商	5
IT科技	网络游戏	6
金融业	银行	7
金融业	基金|理财|信托	8
金融业	保险	9
餐饮	餐饮	10
酒店旅游	酒店	11
酒店旅游	旅游	12
运输与仓储	快递	13
运输与仓储	物流	14
运输与仓储	仓储	15
教育	培训	16
教育	院校	17
政府与公共事业	学术科研	18
政府与公共事业	交警	19
政府与公共事业	博物馆	20
政府与公共事业	公共事业|非盈利机构	21
医药护理	医药医疗	22
医药护理	护理美容	23
医药护理	保健与卫生	24
交通工具	汽车相关	25
交通工具	摩托车相关	26
交通工具	火车相关	27
交通工具	飞机相关	28
房地产	建筑	29
房地产	物业	30
消费品	消费品	31
商业服务	法律	32
商业服务	会展	33
商业服务	中介服务	34
商业服务	认证	35
商业服务	审计	36
文体娱乐	传媒	37
文体娱乐	体育	38
文体娱乐	娱乐休闲	39
印刷	印刷	40
其它	其它	41
*/
func SetIndustry(primary, secondary int) error {
	url := fmt.Sprintf(TamplateSetIndustryURL, AccessToken())
	body := fmt.Sprintf(`{"industry_id1":"%d","industry_id2":"%d"}`, primary, secondary)

	return Post(url, []byte(body))
}

// TemplateId 模板ID with Error
type TemplateId struct {
	WeixinError
	TemplateId string `json:"template_id"`
}

// AddTemplate 添加模板，获得模板ID
func AddTemplate(shortId string) (templateId string, err error) {
	url := fmt.Sprintf(TemplateAddTemplateURL, AccessToken())
	js := fmt.Sprintf(`{"template_id_short":"%s"}`, shortId)

	t := &TemplateId{}
	err = PostUnmarshal(url, []byte(js), t)
	if err != nil {
		return "", err
	}

	return t.TemplateId, nil
}

// TemplateMsg 模板消息
type TemplateMsg struct {
	ToUser     string       `json:"touser"`
	TemplateId string       `json:"template_id"`
	URL        string       `json:"url,omitempty"`
	Data       TemplateData `json:"data"`
}

// TemplateData 模板消息内容
type TemplateData struct {
	First    KeywordPair `json:"first"`
	Keyword1 KeywordPair `json:"keyword1"`
	Keyword2 KeywordPair `json:"keyword2,omitempty"`
	Keyword3 KeywordPair `json:"keyword3,omitempty"`
	Keyword4 KeywordPair `json:"keyword4,omitempty"`
	Keyword5 KeywordPair `json:"keyword5,omitempty"`
	Remark   KeywordPair `json:"value"`
}

// KeywordPair 模板消息内容值
type KeywordPair struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// TempplateMsgId 模版消息 Id
type TempplateMsgId struct {
	WeixinError
	MsgId int `json:"msgid"`
}

// SendTemplateMsg 发送模板消息，返回消息 Id
func SendTemplateMsg(m *TemplateMsg) (msgId int, err error) {
	url := fmt.Sprintf(TemplateSendTemplateMsgURL, AccessToken())

	mid := &TempplateMsgId{}
	err = PostMarshalUnmarshal(url, m, mid)
	if err != nil {
		return 0, err
	}

	return mid.MsgId, nil
}
