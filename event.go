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
	EventTypeTemplateSendJobFinish = "TEMPLATESENDJOBFINISH"
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
