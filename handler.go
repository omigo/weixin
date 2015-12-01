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

// HandleMessage 处理各类消息
func HandleMessage(msg *Message) (ret ReplyMsg) {
	log.Debugf("process `%s` message", msg.MsgType)

	switch msg.MsgType {
	case MsgTypeText:
		if RecvTextHandler == nil {
			log.Warnf("unregister RecvTextHandler: %s", msg.MsgType)
			return nil
		}
		ret = RecvTextHandler(msg.RecvText())
	case MsgTypeImage:
		if RecvImageHandler == nil {
			log.Warnf("unregister RecvImageHandler: %s", msg.MsgType)
			return nil
		}
		ret = RecvImageHandler(msg.RecvImage())
	case MsgTypeVoice:
		if RecvVoiceHandler == nil {
			log.Warnf("unregister RecvVoiceHandler: %s", msg.MsgType)
			return nil
		}
		ret = RecvVoiceHandler(msg.RecvVoice())
	case MsgTypeVideo:
		if RecvVideoHandler == nil {
			log.Warnf("unregister RecvVideoHandler: %s", msg.MsgType)
			return nil
		}
		ret = RecvVideoHandler(msg.RecvVideo())
	case MsgTypeShortVideo:
		if RecvShortVideoHandler == nil {
			log.Warnf("unregister RecvShortVideoHandler: %s", msg.MsgType)
			return nil
		}
		ret = RecvShortVideoHandler(msg.RecvVideo())
	case MsgTypeLocation:
		if RecvLocationHandler == nil {
			log.Warnf("unregister RecvLocationHandler: %s", msg.MsgType)
			return nil
		}
		ret = RecvLocationHandler(msg.RecvLocation())
	case MsgTypeLink:
		if RecvLinkHandler == nil {
			log.Warnf("unregister RecvLinkHandler: %s", msg.MsgType)
			return nil
		}
		ret = RecvLinkHandler(msg.RecvLink())
	default:
		log.Warnf("unexpected RecvType: %s", msg.MsgType)
	}

	return ret
}
