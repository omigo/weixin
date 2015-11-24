package weixin

import "encoding/xml"

// MsgType 消息类型
type MsgType string

// 微信支持的消息类型
const (
	MsgTypeText       MsgType = "text"       // 文本消息
	MsgTypeImage      MsgType = "image"      // 图片消息
	MsgTypeVoice      MsgType = "voice"      // 语音消息
	MsgTypeVideo      MsgType = "video"      // 视频消息
	MsgTypeShortVideo MsgType = "shortvideo" // 小视频消息
	MsgTypeLocation   MsgType = "location"   // 地理位置消息
	MsgTypeLink       MsgType = "link"       // 链接消息
	MsgTypeMusic      MsgType = "music"      // 音乐消息
	MsgTypeNews       MsgType = "news"       // 图文消息
)

// 接收普通消息

// Message 通用类，接收微信各类消息
type Message struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	FromUserName string   // 发送方帐号（一个OpenID）
	CreateTime   string   // 消息创建时间 （整型）
	MsgId        int      // 消息id，64位整型

	// text-文本消息，image-图片消息，voice-语音消息，
	// video-视频消息，shortvideo-小视频消息，
	// location-地理位置消息，link-链接消息，
	// music-音乐，news-图文消息
	MsgType MsgType

	// text-文本消息
	Content string `xml:",omitempty"` // 文本消息内容

	// image-图片消息
	PicUrl  string `xml:",omitempty"` // 图片链接
	MediaId string `xml:",omitempty"` // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据

	// voice-语音消息
	// MediaId string  `xml:",omitempty"`// 语音消息媒体id，可以调用多媒体文件下载接口拉取数据
	Format       string `xml:",omitempty"` // 语音格式，如amr，speex等
	Recongnition string `xml:",omitempty"` // 语音识别结果，使用UTF8编码

	// video-视频消息，shortvideo-小视频消息
	// MediaId string `xml:",omitempty"` // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据
	ThumbMediaId string `xml:",omitempty"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据

	// location-地理位置消息
	LocationX float64 `xml:"Location_X,omitempty"` // 地理位置维度
	LocationY float64 `xml:"Location_Y,omitempty"` // 地理位置经度
	Scale     int     `xml:",omitempty"`           // 地图缩放大小
	Label     string  `xml:",omitempty"`           // 地理位置信息

	// link-链接消息
	Title       string `xml:",omitempty"` // 消息标题
	Description string `xml:",omitempty"` // 消息描述
	Url         string `xml:",omitempty"` // 消息链接
}

// MsgText 接收文本消息
type MsgText struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间 （整型）
	MsgType      MsgType // text
	Content      string  // 文本消息内容
	MsgId        int     // 消息id，64位整型
}

// MsgText 把通用 struct 转化成相应类型的 struct
func (m *Message) MsgText() *MsgText {
	return &MsgText{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Content:      m.Content,
		MsgId:        m.MsgId,
	}
}

// MsgImage 接收图片消息
type MsgImage struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间 （整型）
	MsgType      MsgType // image
	PicUrl       string  // 图片链接
	MediaId      string  // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据
	MsgId        int     // 消息id，64位整型
}

// MsgImage 把通用 struct 转化成相应类型的 struct
func (m *Message) MsgImage() *MsgImage {
	return &MsgImage{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		PicUrl:       m.PicUrl,
		MediaId:      m.MediaId,
		MsgId:        m.MsgId,
	}
}

// MsgVoice 接收视频/小视频消息
type MsgVoice struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间 （整型）
	MsgType      MsgType // voice
	MediaId      string  // 语音消息媒体id，可以调用多媒体文件下载接口拉取数据
	Format       string  // 语音格式，如amr，speex等
	Recongnition string  // 语音识别结果，使用UTF8编码
	MsgId        int     // 消息id，64位整型
}

// MsgVoice 把通用 struct 转化成相应类型的 struct
func (m *Message) MsgVoice() *MsgVoice {
	return &MsgVoice{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		MediaId:      m.MediaId,
		Format:       m.Format,
		Recongnition: m.Recongnition,
		MsgId:        m.MsgId,
	}
}

// MsgVideo 接收图片消息
type MsgVideo struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间 （整型）
	MsgType      MsgType // video
	MediaId      string  // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据
	ThumbMediaId string  // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据
	MsgId        int     // 消息id，64位整型
}

// MsgVideo 把通用 struct 转化成相应类型的 struct
func (m *Message) MsgVideo() *MsgVideo {
	return &MsgVideo{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		MediaId:      m.MediaId,
		ThumbMediaId: m.ThumbMediaId,
		MsgId:        m.MsgId,
	}
}

// MsgLocation 接收地理位置消息
type MsgLocation struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间 （整型）
	MsgType      MsgType // location
	LocationX    float64 `xml:"Location_X,omitempty"` // 地理位置维度
	LocationY    float64 `xml:"Location_Y,omitempty"` // 地理位置经度
	Scale        int     // 地图缩放大小
	Label        string  // 地理位置信息
	MsgId        int     // 消息id，64位整型
}

// MsgLocation 把通用 struct 转化成相应类型的 struct
func (m *Message) MsgLocation() *MsgLocation {
	return &MsgLocation{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		LocationX:    m.LocationX,
		LocationY:    m.LocationY,
		Scale:        m.Scale,
		Label:        m.Label,
		MsgId:        m.MsgId,
	}
}

// MsgLink 接收链接消息
type MsgLink struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间 （整型）
	MsgType      MsgType // location
	Title        string  // 消息标题
	Description  string  // 消息描述
	Url          string  // 消息链接
	MsgId        int     // 消息id，64位整型
}

// MsgLink 把通用 struct 转化成相应类型的 struct
func (m *Message) MsgLink() *MsgLink {
	return &MsgLink{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Title:        m.Title,
		Description:  m.Description,
		Url:          m.Url,
		MsgId:        m.MsgId,
	}
}

// 被动回复用户消息

// ReplyText 回复文本消息
type ReplyText struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	FromUserName string   // 发送方帐号（一个OpenID）
	CreateTime   string   // 消息创建时间 （整型）
	MsgType      MsgType  // text
	Content      string   // 文本消息内容
}

// ReplyImage 回复图片消息
type ReplyImage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	FromUserName string   // 发送方帐号（一个OpenID）
	CreateTime   string   // 消息创建时间 （整型）
	MsgType      MsgType  // image
	PicUrl       string   // 图片链接
	MediaId      string   `xml:"Image>MediaId"` // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据
}

// ReplyVoice 回复语音消息
type ReplyVoice struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	FromUserName string   // 发送方帐号（一个OpenID）
	CreateTime   string   // 消息创建时间 （整型）
	MsgType      MsgType  // voice
	MediaId      string   `xml:"Voice>MediaId"` // 通过素材管理接口上传多媒体文件，得到的id
}

// ReplyVideo 回复视频消息
type ReplyVideo struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	FromUserName string   // 发送方帐号（一个OpenID）
	CreateTime   string   // 消息创建时间 （整型）
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
	CreateTime   string         // 消息创建时间 （整型）
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
