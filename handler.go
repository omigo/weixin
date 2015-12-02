package weixin

import "github.com/omigo/log"

// 各类消息处理器
var (
	RecvTextHandler       func(*RecvText) ReplyMsg
	RecvImageHandler      func(*RecvImage) ReplyMsg
	RecvVoiceHandler      func(*RecvVoice) ReplyMsg
	RecvVideoHandler      func(*RecvVideo) ReplyMsg
	RecvShortVideoHandler func(*RecvVideo) ReplyMsg
	RecvLocationHandler   func(*RecvLocation) ReplyMsg
	RecvLinkHandler       func(*RecvLink) ReplyMsg
)

// 各类事件处理器
var (
	EventSubscribeHandler   func(*EventSubscribe) ReplyMsg
	EventUnsubscribeHandler func(*EventSubscribe) ReplyMsg
	EventLocationHandler    func(*EventLocation) ReplyMsg
	EventClickHandler       func(*EventClick) ReplyMsg
	EventViewHandler        func(*EventView) ReplyMsg
)

// RecvDefaultHandler 如果没有注册某类消息处理器，那么收到这类消息时，使用这个默认处理器
var RecvDefaultHandler = func(msg *Message) (reply ReplyMsg) {
	log.Debugf("unregistered receive message handler %s, use RecvDefaultHandler", msg.MsgType)
	return nil
}

// EventDefaultHandler 如果没有注册某类事件处理器，那么收到这类事件时，使用这个默认处理器
var EventDefaultHandler = func(msg *Message) (reply ReplyMsg) {
	log.Debugf("unregistered receive event handler %s, use EventDefaultHandler", msg.Event)
	return nil
}

// HandleMessage 处理各类消息
func HandleMessage(msg *Message) (ret ReplyMsg) {
	log.Debugf("process `%s` message", msg.MsgType)

	switch msg.MsgType {
	case MsgTypeText:
		if RecvTextHandler != nil {
			return RecvTextHandler(NewRecvText(msg))
		}
	case MsgTypeImage:
		if RecvImageHandler != nil {
			return RecvImageHandler(NewRecvImage(msg))
		}
	case MsgTypeVoice:
		if RecvVoiceHandler != nil {
			return RecvVoiceHandler(NewRecvVoice(msg))
		}
	case MsgTypeVideo:
		if RecvVideoHandler != nil {
			return RecvVideoHandler(NewRecvVideo(msg))
		}
	case MsgTypeShortVideo:
		if RecvShortVideoHandler != nil {
			return RecvShortVideoHandler(NewRecvVideo(msg))
		}
	case MsgTypeLocation:
		if RecvLocationHandler != nil {
			return RecvLocationHandler(NewRecvLocation(msg))
		}
	case MsgTypeLink:
		if RecvLinkHandler != nil {
			return RecvLinkHandler(NewRecvLink(msg))
		}
	case MsgTypeEvent:
		return HandleEvent(msg)
	default:
		log.Errorf("unexpected receive MsgType: %s", msg.MsgType)
		return nil
	}

	return RecvDefaultHandler(msg)
}

// HandleEvent 处理各类事件
func HandleEvent(msg *Message) (reply ReplyMsg) {
	log.Debugf("process `%s` event", msg.MsgType)

	switch msg.Event {
	case EventTypeSubscribe:
		if EventSubscribeHandler != nil {
			return EventSubscribeHandler(NewEventSubscribe(msg))
		}
	case EventTypeUnsubscribe:
		if EventUnsubscribeHandler != nil {
			return EventUnsubscribeHandler(NewEventSubscribe(msg))
		}
	case EventTypeLocation:
		if EventLocationHandler != nil {
			return EventLocationHandler(NewEventLocation(msg))
		}
	case EventTypeClick:
		if EventClickHandler != nil {
			return EventClickHandler(NewEventClick(msg))
		}
	case EventTypeView:
		if EventViewHandler != nil {
			return EventViewHandler(NewEventView(msg))
		}
	default:
		log.Errorf("unexpected receive EventType: %s", msg.Event)
		return nil
	}

	return EventDefaultHandler(msg)
}
