package weixin

import (
	"encoding/json"
	"fmt"
	"os"
)

// 客服消息接口
const (
	CustomAddURL    = "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s"
	CustomUpdateURL = "https://api.weixin.qq.com/customservice/kfaccount/update?access_token=%s"
	CustomDeleteURL = "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=%s"
	// 设置客服帐号的头像
	CustomHeadingURL = "http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=%s&kf_account=%s"
	CustomListURL    = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=%s"
	CustomMsgURL     = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"
)

// Custom 客服帐号
type Custom struct {
	Account  string `json:"kf_account"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
}

// AddCustom 添加客服帐号
func AddCustom(account, nickname, password string) (err error) {
	url := fmt.Sprintf(CustomAddURL, AccessToken())
	return operateCustomStruct(url, &Custom{account, nickname, password})
}

// UpdateCustom 修改客服帐号
func UpdateCustom(account, nickname, password string) (err error) {
	url := fmt.Sprintf(CustomUpdateURL, AccessToken())
	return operateCustomStruct(url, &Custom{account, nickname, password})
}

// DeleteCustom 删除客服帐号
func DeleteCustom(account, nickname, password string) (err error) {
	url := fmt.Sprintf(CustomDeleteURL, AccessToken())
	return operateCustomStruct(url, &Custom{account, nickname, password})
}

// AddCustomStruct 添加客服帐号
func AddCustomStruct(cust *Custom) (err error) {
	url := fmt.Sprintf(CustomAddURL, AccessToken())
	return operateCustomStruct(url, cust)
}

// UpdateCustomStruct 修改客服帐号
func UpdateCustomStruct(cust *Custom) (err error) {
	url := fmt.Sprintf(CustomAddURL, AccessToken())
	return operateCustomStruct(url, cust)
}

// DeleteCustomStruct 删除客服帐号
func DeleteCustomStruct(cust *Custom) (err error) {
	url := fmt.Sprintf(CustomDeleteURL, AccessToken())
	return operateCustomStruct(url, cust)
}

// operateCustomStruct 客服帐号
func operateCustomStruct(url string, cust *Custom) (err error) {
	return Post(url, cust, nil)
}

// UploadHeading 设置客服帐号的头像
func UploadHeading(account string, file *os.File) (err error) {
	url := fmt.Sprintf(CustomHeadingURL, AccessToken(), account)
	return Upload(url, file.Name(), file, nil)
}

// CustomList 客服列表
type CustomList struct {
	WXError
	List []*Account `json:"kf_list"`
}

// Account 客服账号
type Account struct {
	Account    string `json:"kf_account"`
	NickName   string `json:"kf_nick"`
	AccountId  string `json:"kf_id"`
	HeadImgURL string `json:"kf_headimgurl"`
}

// GetCustomList 获取所有客服账号
func GetCustomList() (accs []*Account, err error) {
	url := fmt.Sprintf(CustomListURL, AccessToken())
	list := &CustomList{}
	err = Get(url, list)
	return list.List, err
}

// CustMsg 客服消息接口
type CustMsg interface{}

// CustText 客服接口发送文本消息
type CustText struct {
	Content string `json:"content"` //	文本消息内容
}

// CustImage 客服接口发送图片消息
type CustImage struct {
	MediaId string `json:"media_id"` //	发送的图片/语音/视频的媒体ID
}

// CustVoice 客服接口发送语音消息
type CustVoice struct {
	MediaId string `json:"media_id"` //	发送的图片/语音/视频的媒体ID
}

// CustVideo 客服接口发送视频消息
type CustVideo struct {
	MediaId      string `json:"media_id"`       //	发送的图片/语音/视频的媒体ID
	ThumbMediaId string `json:"thumb_media_id"` //	缩略图的媒体ID
	Title        string `json:"title"`          //	图文消息/视频消息/音乐消息的标题
	Description  string `json:"description"`    //	图文消息/视频消息/音乐消息的描述
}

// CustMusic 客服接口发送音乐消息
type CustMusic struct {
	Title        string `json:"title"`          //	图文消息/视频消息/音乐消息的标题
	Description  string `json:"description"`    //	图文消息/视频消息/音乐消息的描述
	MusicURL     string `json:"musicurl"`       //	音乐链接
	HQMusicURL   string `json:"hqmusicurl"`     //	高品质音乐链接，wifi环境优先使用该链接播放音乐
	ThumbMediaId string `json:"thumb_media_id"` //	缩略图的媒体ID
}

// CustNewsArticle 客服接口发送图文消息。图文消息条数限制在10条以内，注意，如果图文数超过10，则将会无响应
type CustNewsArticle struct {
	Title       string `json:"title"`       //	图文消息/视频消息/音乐消息的标题
	Description string `json:"description"` //	图文消息/视频消息/音乐消息的描述
	URL         string `json:"url"`         //	图文消息被点击后跳转的链接
	PicURL      string `json:"picurl"`      //	图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图640*320，小图80*80
}

// CustWXCard 客服接口发送卡券
type CustWXCard struct {
	CardId  string `json:"card_id"`
	CardExt string `json:"card_ext"`
}

// SendCustomMsg 客服接口-发消息
func SendCustomMsg(openId string, msg CustMsg) (err error) {
	msgType := MsgTypeText
	switch msg.(type) {
	case *CustText:
		msgType = MsgTypeText
	case *CustImage:
		msgType = MsgTypeImage
	case *CustVoice:
		msgType = MsgTypeVoice
	case *CustVideo:
		msgType = MsgTypeVideo
	case *CustMusic:
		msgType = MsgTypeMusic
	case *CustNewsArticle:
		msgType = MsgTypeNews
	case *CustWXCard:
		msgType = MsgTypeWXCard
	default:
		panic("unexpected custom message type")
	}

	js, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	jsonStr := fmt.Sprintf(`{"touser":"%s","msgtype":"%s","%s": %s}`,
		openId, msgType, msgType, js)

	url := fmt.Sprintf(CustomMsgURL, AccessToken())
	return Post(url, []byte(jsonStr), nil)
}
