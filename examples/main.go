package main

import (
	"fmt"
	"net/http"

	"github.com/omigo/log"
	"github.com/omigo/weixin"
)

func main() {
	addr := ":3080"

	weixin.Initialize(originId, appId, appSecret, token, encodingAESKey)

	weixin.RecvTextHandler = echoMsgText             // 注册文本消息处理器
	weixin.RecvImageHandler = echoMsgImage           // 注册图片消息处理器
	weixin.RecvVoiceHandler = echoMsgVoice           // 注册语音消息处理器
	weixin.RecvVideoHandler = echoMsgVideo           // 注册视频消息处理器
	weixin.RecvShortVideoHandler = echoMsgShortVideo // 注册小视频消息处理器
	weixin.RecvLocationHandler = echoMsgLocation     // 注册位置消息处理器
	weixin.RecvLinkHandler = echoMsgLink             // 注册链接消息处理器
	weixin.RecvDefaultHandler = defaultHandler       // 注册默认处理器

	weixin.EventSubscribeHandler = EventSubscribeHandler     // 注册关注事件处理器
	weixin.EventUnsubscribeHandler = EventUnsubscribeHandler // 注册取消关注事件处理器
	weixin.EventLocationHandler = EventLocationHandler       // 注册上报地理位置事件处理器
	weixin.EventClickHandler = EventClickHandler             // 注册点击自定义菜单事件处理器
	weixin.EventViewHandler = EventViewHandler               // 注册点击菜单跳转链接时的事件处理器
	weixin.EventDefaultHandler = EventDefaultHandler         // 注册默认处理器

	http.HandleFunc("/weixin", weixin.HandleAccess)

	log.Debugf("server is running at %s", addr)
	http.ListenAndServe(addr, nil)
}

func echoMsgText(m *weixin.RecvText) weixin.ReplyMsg {
	log.Debugf("receive message: %+v", m)

	// echo message
	ret := &weixin.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      m.FromUserName + ", " + m.Content,
	}

	log.Debugf("replay message: %+v", ret)
	return ret
}

func echoMsgImage(m *weixin.RecvImage) weixin.ReplyMsg {
	log.Debugf("%+v", m)

	// echo message
	ret := &weixin.ReplyImage{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		PicUrl:       m.PicUrl,
		MediaId:      m.MediaId,
	}

	log.Debugf("%+v", ret)
	return ret
}

func echoMsgVoice(m *weixin.RecvVoice) weixin.ReplyMsg {
	log.Debugf("%+v", m)

	// echo message
	ret := &weixin.ReplyVoice{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		MediaId:      m.MediaId,
	}

	log.Debugf("%+v", ret)
	return ret
}

func echoMsgVideo(m *weixin.RecvVideo) weixin.ReplyMsg {
	log.Debugf("%+v", m)

	// MediaId ???
	ret := &weixin.ReplyVideo{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		MediaId:      m.MediaId,
		Title:        "video",
		Description:  "thist is a test desc...",
	}

	log.Debugf("%+v", ret)
	return ret
}

func echoMsgShortVideo(m *weixin.RecvVideo) weixin.ReplyMsg {
	log.Debugf("%+v", m)

	// MediaId ???
	ret := &weixin.ReplyVideo{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		MediaId:      m.ThumbMediaId,
		Title:        "shortvideo",
		Description:  "thist is a test desc...",
	}

	log.Debugf("%+v", ret)
	return ret
}

func echoMsgLocation(m *weixin.RecvLocation) weixin.ReplyMsg {
	log.Debugf("%+v", m)

	// echo message
	ret := &weixin.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      weixin.AccessToken(),
	}

	log.Debugf("replay message: %+v", ret)
	return ret
}

func echoMsgLink(m *weixin.RecvLink) weixin.ReplyMsg {
	log.Debugf("%+v", m)

	// 回复图文消息

	return nil
}

func defaultHandler(msg *weixin.Message) weixin.ReplyMsg {
	log.Debugf("%+v", msg)

	// echo message
	ret := &weixin.ReplyText{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   msg.CreateTime,
		Content:      "openId: " + msg.FromUserName,
	}

	log.Debugf("replay message: %+v", ret)
	return ret
}

// EventSubscribeHandler 注册关注事件处理器
func EventSubscribeHandler(m *weixin.EventSubscribe) weixin.ReplyMsg {
	log.Debugf("%+v", m)

	// echo message
	ret := &weixin.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      fmt.Sprintf("Event=%s, EventKey=%s, Ticket=%s", m.Event, m.EventKey, m.Ticket),
	}

	log.Debugf("replay message: %+v", ret)
	return ret
}

// EventUnsubscribeHandler 注册取消关注事件处理器
func EventUnsubscribeHandler(m *weixin.EventSubscribe) weixin.ReplyMsg {
	log.Debugf("someone gone")
	return nil
}

// EventLocationHandler 注册上报地理位置事件处理器
func EventLocationHandler(m *weixin.EventLocation) weixin.ReplyMsg {
	log.Debugf("%+v", m)

	// echo message
	ret := &weixin.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content: fmt.Sprintf("Latitude=%s, Longitude=%s, Precision=%s",
			m.Latitude, m.Longitude, m.Precision),
	}

	log.Debugf("replay message: %+v", ret)
	return ret
}

// EventClickHandler 注册点击自定义菜单事件处理器
func EventClickHandler(m *weixin.EventClick) weixin.ReplyMsg {
	log.Debugf("%+v", m)

	// echo message
	ret := &weixin.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      fmt.Sprintf("Event=%s, EventKey=%s", m.Event, m.EventKey),
	}

	log.Debugf("replay message: %+v", ret)
	return ret
}

// EventViewHandler 注册点击菜单跳转链接时的事件处理器
func EventViewHandler(m *weixin.EventView) weixin.ReplyMsg {
	log.Debugf("%+v", m)

	// echo message 貌似用户收不到回复的消息？？？
	ret := &weixin.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      fmt.Sprintf("Event=%s, EventKey=%s", m.Event, m.EventKey),
	}

	log.Debugf("replay message: %+v", ret)
	return ret
}

// EventDefaultHandler 注册默认处理器
func EventDefaultHandler(m *weixin.Message) weixin.ReplyMsg {
	log.Debugf("%+v", m)

	// echo message
	ret := &weixin.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      fmt.Sprintf("Event=%s", m.Event),
	}

	log.Debugf("replay message: %+v", ret)
	return ret
}
