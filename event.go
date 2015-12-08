package weixin

import "github.com/omigo/log"

// EventType 事件类型
type EventType string

// 微信支持的事件类型
const (
	EventTypeSubscribe             = "subscribe"
	EventTypeUnsubscribe           = "unsubscribe"
	EventTypeLocation              = "LOCATION"
	EventTypeClick                 = "CLICK"
	EventTypeView                  = "VIEW"
	EventTypeTemplateSendJobFinish = "TEMPLATESENDJOBFINISH" // 模版消息发送结果通知事件
	EventTypeScancodePush          = "scancode_push"         // 扫码推事件的事件推送
	EventTypeScancodeWaitmsg       = "scancode_waitmsg"      // 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventTypePicSysphoto           = "pic_sysphoto"          // 弹出系统拍照发图的事件推送
	EventTypePicPhotoOrAlbum       = "pic_photo_or_album"    // 弹出拍照或者相册发图的事件推送
	EventTypePicWeixin             = "pic_weixin"            // 弹出微信相册发图器的事件推送
	EventTypeLocationSelect        = "location_select"       // 弹出地理位置选择器的事件推送
)

// RecvEvent 事件消息
type RecvEvent interface {
	RecvMsg
}

// NewRecvEvent 把通用 struct 转化成相应类型的 struct
func NewRecvEvent(msg *Message) RecvEvent {
	switch msg.Event {
	case EventTypeSubscribe:
		return NewEventSubscribe(msg)
	case EventTypeUnsubscribe:
		return NewEventSubscribe(msg)
	case EventTypeLocation:
		return NewEventLocation(msg)
	case EventTypeClick:
		return NewEventClick(msg)
	case EventTypeView:
		return NewEventView(msg)
	case EventTypeScancodePush:
		return NewEventScancodePush(msg)
	case EventTypeScancodeWaitmsg:
		return NewEventScancodeWaitmsg(msg)
	case EventTypePicSysphoto:
		return NewEventPicSysphoto(msg)
	case EventTypePicPhotoOrAlbum:
		return NewEventPicPhotoOrAlbum(msg)
	case EventTypePicWeixin:
		return NewEventPicWeixin(msg)
	case EventTypeLocationSelect:
		return NewEventLocationSelect(msg)
	default:
		log.Errorf("unexpected receive EventType: %s", msg.Event)
		return nil
	}
}

// EventSubscribe 关注/取消关注事件
type EventSubscribe struct {
	RecvEvent
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型，subscribe(订阅)、unsubscribe(取消订阅)
	// 用户扫描带场景值二维码时，可能推送以下两种事件：
	// 1. 如果用户还未关注公众号，则用户可以关注公众号，关注后微信会将带场景值关注事件推送给开发者。
	//    EventKey	事件KEY值，qrscene_为前缀，后面为二维码的参数值
	// 2. 如果用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者。
	//    EventKey	事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	EventKey string
	Ticket   string // 二维码的ticket，可用来换取二维码图片

}

// NewEventSubscribe 把通用 struct 转化成相应类型的 struct
func NewEventSubscribe(m *Message) *EventSubscribe {
	return &EventSubscribe{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Event:        m.Event,
		EventKey:     m.EventKey,
		Ticket:       m.Ticket,
	}
}

// EventLocation 上报地理位置事件
type EventLocation struct {
	RecvEvent
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型，LOCATION
	Latitude     float64   // 地理位置纬度
	Longitude    float64   // 地理位置经度
	Precision    float64   // 地理位置精度
}

// NewEventLocation 把通用 struct 转化成相应类型的 struct
func NewEventLocation(m *Message) *EventLocation {
	return &EventLocation{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Event:        m.Event,
		Latitude:     m.Latitude,
		Longitude:    m.Longitude,
		Precision:    m.Precision,
	}
}

// EventClick 点击菜单拉取消息时的事件推送
type EventClick struct {
	RecvEvent
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型，CLICK
	EventKey     string    // 事件KEY值，与自定义菜单接口中KEY值对应
}

// NewEventClick 把通用 struct 转化成相应类型的 struct
func NewEventClick(m *Message) *EventClick {
	return &EventClick{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Event:        m.Event,
		EventKey:     m.EventKey,
	}
}

// EventView 点击菜单跳转链接时的事件推送
type EventView struct {
	RecvEvent
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型，VIEW
	EventKey     string    // 事件KEY值，设置的跳转URL
}

// NewEventView 把通用 struct 转化成相应类型的 struct
func NewEventView(m *Message) *EventView {
	return &EventView{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Event:        m.Event,
		EventKey:     m.EventKey,
	}
}

// EventTemplateSendJobFinish 模版消息发送结果通知事件
type EventTemplateSendJobFinish struct {
	RecvEvent
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型，VIEW
	MsgID        int       // 消息id
	Status       string    // 发送状态为成功
}

// NewEventTemplateSendJobFinish 把通用 struct 转化成相应类型的 struct
func NewEventTemplateSendJobFinish(m *Message) *EventTemplateSendJobFinish {
	return &EventTemplateSendJobFinish{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Event:        m.Event,
		MsgID:        m.TplMsgId,
		Status:       m.Status,
	}
}

// EventScancodePush 扫码推事件的事件推送
type EventScancodePush struct {
	RecvEvent
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型，scancode_push
	EventKey     string    // 事件KEY值，由开发者在创建菜单时设定
	ScanType     string    // 扫描类型，一般是qrcode
	ScanResult   string    // 扫描结果，即二维码对应的字符串信息
}

// NewEventScancodePush 把通用 struct 转化成相应类型的 struct
func NewEventScancodePush(m *Message) *EventScancodePush {
	return &EventScancodePush{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Event:        m.Event,
		ScanType:     m.ScanType,
		ScanResult:   m.ScanResult,
	}
}

// EventScancodeWaitmsg 扫码推事件且弹出“消息接收中”提示框的事件推送
type EventScancodeWaitmsg struct {
	RecvEvent
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型，scancode_waitmsg
	EventKey     string    // 事件KEY值，由开发者在创建菜单时设定
	ScanType     string    // 扫描类型，一般是qrcode
	ScanResult   string    // 扫描结果，即二维码对应的字符串信息
}

// NewEventScancodeWaitmsg 把通用 struct 转化成相应类型的 struct
func NewEventScancodeWaitmsg(m *Message) *EventScancodeWaitmsg {
	return &EventScancodeWaitmsg{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Event:        m.Event,
		ScanType:     m.ScanType,
		ScanResult:   m.ScanResult,
	}
}

// EventPicSysphoto 弹出系统拍照发图的事件推送
type EventPicSysphoto struct {
	RecvEvent
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型，pic_sysphoto
	EventKey     string    // 事件KEY值，由开发者在创建菜单时设定
	Count        int       // 发送的图片数量
	PicMd5Sums   []string  // 图片的MD5值，开发者若需要，可用于验证接收到图片
}

// NewEventPicSysphoto 把通用 struct 转化成相应类型的 struct
func NewEventPicSysphoto(m *Message) *EventPicSysphoto {
	return &EventPicSysphoto{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Event:        m.Event,
		Count:        m.Count,
		PicMd5Sums:   m.PicMd5Sums,
	}
}

// EventPicPhotoOrAlbum 弹出拍照或者相册发图的事件推送
type EventPicPhotoOrAlbum struct {
	RecvEvent
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型，pic_sysphoto
	EventKey     string    // 事件KEY值，由开发者在创建菜单时设定
	Count        int       // 发送的图片数量
	PicMd5Sums   []string  // 图片的MD5值，开发者若需要，可用于验证接收到图片
}

// NewEventPicPhotoOrAlbum 把通用 struct 转化成相应类型的 struct
func NewEventPicPhotoOrAlbum(m *Message) *EventPicPhotoOrAlbum {
	return &EventPicPhotoOrAlbum{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Event:        m.Event,
		Count:        m.Count,
		PicMd5Sums:   m.PicMd5Sums,
	}
}

// EventPicWeixin 弹出微信相册发图器的事件推送
type EventPicWeixin struct {
	RecvEvent
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型，pic_sysphoto
	EventKey     string    // 事件KEY值，由开发者在创建菜单时设定
	Count        int       // 发送的图片数量
	PicMd5Sums   []string  // 图片的MD5值，开发者若需要，可用于验证接收到图片
}

// NewEventPicWeixin 把通用 struct 转化成相应类型的 struct
func NewEventPicWeixin(m *Message) *EventPicWeixin {
	return &EventPicWeixin{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Event:        m.Event,
		Count:        m.Count,
		PicMd5Sums:   m.PicMd5Sums,
	}
}

// EventLocationSelect 弹出地理位置选择器的事件推送
type EventLocationSelect struct {
	RecvEvent
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型，pic_sysphoto
	EventKey     string    // 事件KEY值，由开发者在创建菜单时设定
	LocationX    float64   // 地理位置维度
	LocationY    float64   // 地理位置经度
	Scale        int       // 精度，可理解为精度或者比例尺、越精细的话 scale越高
	Label        string    // 地理位置的字符串信息
	Poiname      string    // 朋友圈POI的名字，可能为空
}

// NewEventLocationSelect 把通用 struct 转化成相应类型的 struct
func NewEventLocationSelect(m *Message) *EventLocationSelect {
	return &EventLocationSelect{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Event:        m.Event,
		LocationX:    m.LocationX2,
		LocationY:    m.LocationY2,
		Scale:        m.Scale2,
		Label:        m.Label2,
		Poiname:      m.Poiname,
	}
}
