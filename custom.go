package weixin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/omigo/log"
)

const (
	customAddURL = "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s"
	customMsgURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"
)

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
func SendCustomMsg(openId string, msg interface{}) (err error) {
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
	ret := fmt.Sprintf(`{"touser":"%s","msgtype":"%s","%s": %s}`,
		openId, msgType, msgType, js)

	url := fmt.Sprintf(customMsgURL, AccessToken())
	log.Debugf("custMsgURL=%s", url)
	resp, err := http.Post(url, "application/json", bytes.NewBufferString(ret))
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debugf("custom message sended: %s", body)

	return nil
}
