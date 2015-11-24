package weixin

import "github.com/omigo/log"

// MessageHandler 处理各类消息
type MessageHandler func(*Message) interface{}

// Handlers 各类消息处理器
var (
	MsgTextHandler     func(*MsgText) interface{}
	MsgImageHandler    func(*MsgImage) interface{}
	MsgVoiceHandler    func(*MsgVoice) interface{}
	MsgVideoHandler    func(*MsgVideo) interface{}
	MsgLocationHandler func(*MsgLocation) interface{}
	MsgLinkHandler     func(*MsgLink) interface{}
)

func processMessage(msg *Message) (ret interface{}) {
	log.Debugf("process %s message", msg.MsgType)

	switch msg.MsgType {
	case MsgTypeText:
		if MsgTextHandler == nil {
			log.Warnf("unregister MsgType: %s", msg.MsgType)
			return nil
		}
		ret = MsgTextHandler(msg.MsgText())
	case MsgTypeImage:
		if MsgImageHandler == nil {
			log.Warnf("unregister MsgType: %s", msg.MsgType)
			return nil
		}
		ret = MsgImageHandler(msg.MsgImage())
	case MsgTypeVoice:
		if MsgVoiceHandler == nil {
			log.Warnf("unregister MsgType: %s", msg.MsgType)
			return nil
		}
		ret = MsgVoiceHandler(msg.MsgVoice())
	case MsgTypeVideo, MsgTypeShortVideo:
		if MsgVideoHandler == nil {
			log.Warnf("unregister MsgType: %s", msg.MsgType)
			return nil
		}
		ret = MsgVideoHandler(msg.MsgVideo())
	case MsgTypeLocation:
		if MsgLocationHandler == nil {
			log.Warnf("unregister MsgType: %s", msg.MsgType)
			return nil
		}
		ret = MsgLocationHandler(msg.MsgLocation())
	case MsgTypeLink:
		if MsgLinkHandler == nil {
			log.Warnf("unregister MsgType: %s", msg.MsgType)
			return nil
		}
		ret = MsgLinkHandler(msg.MsgLink())
	default:
		log.Warnf("unexpected MsgType: %s", msg.MsgType)
	}

	return ret
}
