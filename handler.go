package weixin

import "github.com/omigo/log"

// Handlers 各类消息处理器
var (
	RecvTextHandler       func(*RecvText) ReplyMsg
	RecvImageHandler      func(*RecvImage) ReplyMsg
	RecvVoiceHandler      func(*RecvVoice) ReplyMsg
	RecvVideoHandler      func(*RecvVideo) ReplyMsg
	RecvShortVideoHandler func(*RecvVideo) ReplyMsg
	RecvLocationHandler   func(*RecvLocation) ReplyMsg
	RecvLinkHandler       func(*RecvLink) ReplyMsg
)

// RecvDefaultHandler 如果没有注册某类消息，那么收到这类消息时，使用这个默认处理器
var RecvDefaultHandler = func(msg *Message) (reply ReplyMsg) {
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
	default:
		log.Errorf("unexpected receive MsgType: %s", msg.MsgType)
		return nil
	}

	log.Debugf("unregistered receive message handler: %s", msg.MsgType)
	return RecvDefaultHandler(msg)
}
