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
	MsgTypeWXCard     MsgType = "wxcard"     // 卡券，客服消息时使用
	MsgTypeEvent      MsgType = "event"      // 事件推送
)

// EncMessage 加密消息体
type EncMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"-"` // 开发者微信号
	Encrypt      string   // 加密的消息报文
	MsgSignature string   // 报文签名
	TimeStamp    string   // 时间戳
	Nonce        string   // 随机数
}

// 接收普通消息

// Message 通用类，接收微信各类消息
type Message struct {
	Encrypt string

	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	FromUserName string   // 发送方帐号（一个OpenID）
	CreateTime   string   // 消息创建时间（整型）
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
	Format      string `xml:",omitempty"` // 语音格式，如amr，speex等
	Recognition string `xml:",omitempty"` // 语音识别结果，使用UTF8编码

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

	// 接收事件推送

	// 关注/取消关注事件
	Event EventType // 事件类型，subscribe(订阅)、unsubscribe(取消订阅)

	// 扫描带参数二维码事件
	// Event	EventType // 事件类型，subscribe
	// 用户扫描带场景值二维码时，可能推送以下两种事件：
	// 1. 如果用户还未关注公众号，则用户可以关注公众号，关注后微信会将带场景值关注事件推送给开发者。
	//    EventKey	事件KEY值，qrscene_为前缀，后面为二维码的参数值
	// 2. 如果用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者。
	//    EventKey	事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	EventKey string // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket   string // 二维码的ticket，可用来换取二维码图片

	// 上报地理位置事件
	// Event EventType	// 事件类型，LOCATION
	Latitude  float64 // 地理位置纬度
	Longitude float64 // 地理位置经度
	Precision float64 // 地理位置精度

	// 点击菜单拉取消息时的事件推送
	// Event    EventType // 事件类型，CLICK
	// EventKey string    // 事件KEY值，与自定义菜单接口中KEY值对应

	// 点击菜单跳转链接时的事件推送
	// Event    EventType // 事件类型，VIEW
	// EventKey string    // 事件KEY值，设置的跳转URL

	// 模版消息发送任务完成后，微信服务器会将是否送达成功作为通知
	// Event         EventType // 事件类型，TEMPLATESENDJOBFINISH
	TplMsgId int    // 消息id
	Status   string // 发送状态为成功

	// 事件类型，scancode_push/scancode_waitmsg
	// Event        EventType // 事件类型，scancode_push
	// EventKey     string    // 事件KEY值，由开发者在创建菜单时设定
	ScanType   string `json:"ScanCodeInfo>ScanType"`   // 扫描类型，一般是qrcode
	ScanResult string `json:"ScanCodeInfo>ScanResult"` // 扫描结果，即二维码对应的字符串信息

	// 事件类型，pic_sysphoto
	// Event      EventType // 事件类型，pic_sysphoto
	// EventKey   string    // 事件KEY值，由开发者在创建菜单时设定
	Count      int      `json:"SendPicsInfo>Count"`                  // 发送的图片数量
	PicMd5Sums []string `json:"SendPicsInfo>PicList>item>PicMd5Sum"` // 图片的MD5值，开发者若需要，可用于验证接收到图片

	LocationX2 float64 `xml:"SendLocationInfo>Location_X,omitempty"` // 地理位置维度
	LocationY2 float64 `xml:"SendLocationInfo>Location_Y,omitempty"` // 地理位置经度
	Scale2     int     `xml:"SendLocationInfo>Scale,omitempty"`      // 精度，可理解为精度或者比例尺、越精细的话 scale越高
	Label2     string  `xml:"SendLocationInfo>Label,omitempty"`      // 地理位置的字符串信息
	Poiname    string  `xml:"SendLocationInfo>Poiname,omitempty"`    // 朋友圈POI的名字，可能为空

	// 事件类型 qualification_verify_success,naming_verify_success,annual_renew,verify_expired
	ExpiredTime int // 有效期 (整形)，指的是时间戳，将于该时间戳认证过期

	// 事件类型，qualification_verify_fail,naming_verify_fail
	FailTime   int    // 失败发生时间 (整形)，时间戳
	FailReason string // 认证失败的原因
}
