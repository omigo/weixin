package weixin

import "github.com/gotips/log"

// EventType 事件类型
type EventType string

// RecvEvent 事件消息
type RecvEvent interface {
	RecvMsg
}

// EventBase 事件基础类
type EventBase struct {
	RecvMsg
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型
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
	case EventTypeQualificationVerifySuccess:
		return NewEventQualificationVerifySuccess(msg)
	case EventTypeQualificationVerifyFail:
		return NewEventQualificationVerifyFail(msg)
	case EventTypeNamingVerifySuccess:
		return NewEventNamingVerifySuccess(msg)
	case EventTypeNamingVerifyFail:
		return NewEventNamingVerifyFail(msg)
	case EventTypeAnnualRenew:
		return NewEventAnnualRenew(msg)
	case EventTypeVerifyExpired:
		return NewEventVerifyExpired(msg)
	default:
		log.Errorf("unexpected receive EventType: %s", msg.Event)
		return nil
	}
}
