package weixin

import "encoding/xml"

// 被动回复用户消息

// ReplyText 回复文本消息
type ReplyText struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	FromUserName string   // 发送方帐号（一个OpenID）
	CreateTime   string   // 消息创建时间（整型）
	MsgType      MsgType  // text
	Content      string   // 文本消息内容
}

// ReplyImage 回复图片消息
type ReplyImage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	FromUserName string   // 发送方帐号（一个OpenID）
	CreateTime   string   // 消息创建时间（整型）
	MsgType      MsgType  // image
	PicUrl       string   // 图片链接
	MediaId      string   `xml:"Image>MediaId"` // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据
}

// ReplyVoice 回复语音消息
type ReplyVoice struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	FromUserName string   // 发送方帐号（一个OpenID）
	CreateTime   string   // 消息创建时间（整型）
	MsgType      MsgType  // voice
	MediaId      string   `xml:"Voice>MediaId"` // 通过素材管理接口上传多媒体文件，得到的id
}

// ReplyVideo 回复视频消息
type ReplyVideo struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	FromUserName string   // 发送方帐号（一个OpenID）
	CreateTime   string   // 消息创建时间（整型）
	MsgType      MsgType  // video
	MediaId      string   `xml:"Video>MediaId"`               // 通过素材管理接口上传多媒体文件，得到的id
	Title        string   `xml:"Video>Title,omitempty"`       // 视频消息的标题
	Description  string   `xml:"Video>Description,omitempty"` // 视频消息的描述
}

// ReplyMusic 回复音乐消息
type ReplyMusic struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	FromUserName string   // 发送方帐号（一个OpenID）
	CreateTime   string   // 消息创建时间（整型）
	MsgType      MsgType  // music
	Title        string   `xml:"Music>Title,omitempty"`        // 音乐标题
	Description  string   `xml:"Music>Description,omitempty"`  // 音乐描述
	MusicURL     string   `xml:"Music>MusicURL,omitempty"`     // 	音乐链接
	HQMusicUrl   string   `xml:"Music>HQMusicUrl,omitempty"`   // 	高质量音乐链接，WIFI环境优先使用该链接播放音乐
	ThumbMediaId string   `xml:"Music>ThumbMediaId,omitempty"` // 	缩略图的媒体id，通过素材管理接口上传多媒体文件，得到的id
}

// ReplyNews 回复图文消息
type ReplyNews struct {
	XMLName      xml.Name       `xml:"xml"`
	ToUserName   string         // 开发者微信号
	FromUserName string         // 发送方帐号（一个OpenID）
	CreateTime   string         // 消息创建时间（整型）
	MsgType      MsgType        // location
	ArticleCount int            // 图文消息个数，限制为10条以内
	Articles     []ReplyArticle `xml:"Articles>item"` // 多条图文消息信息，默认第一个item为大图,注意，如果图文数超过10，则将会无响应
}

// ReplyArticle 图文消息
type ReplyArticle struct {
	Title       string `xml:",omitempty"` // 	图文消息标题
	Description string `xml:",omitempty"` // 	图文消息描述
	PicUrl      string `xml:",omitempty"` // 	图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
	Url         string `xml:",omitempty"` // 	点击图文消息跳转链接
}
