package main

import (
	"net/http"

	"github.com/omigo/log"
	"github.com/omigo/weixin"
)

func main() {
	addr := ":3080"

	weixin.Initialize(originId, appId, appSecret, token, encodingAESKey)

	weixin.MsgTextHandler = echoMsgText             // 注册文本消息处理器
	weixin.MsgImageHandler = echoMsgImage           // 注册图片消息处理器
	weixin.MsgVoiceHandler = echoMsgVoice           // 注册语音消息处理器
	weixin.MsgVideoHandler = echoMsgVideo           // 注册视频消息处理器
	weixin.MsgShortVideoHandler = echoMsgShortVideo // 注册小视频消息处理器
	weixin.MsgLocationHandler = echoMsgLocation     // 注册位置消息处理器
	weixin.MsgLinkHandler = echoMsgLink             // 注册链接消息处理器

	http.HandleFunc("/weixin", weixin.HandleAccess)

	log.Debugf("server is running at %s", addr)
	http.ListenAndServe(addr, nil)
}

func echoMsgText(m *weixin.MsgText) interface{} {
	log.Debugf("%+v", m)

	// echo message
	ret := &weixin.ReplyText{
		MsgType:      m.MsgType,
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      m.Content,
	}

	log.Debugf("%+v", ret)
	return ret
}

func echoMsgImage(m *weixin.MsgImage) interface{} {
	log.Debugf("%+v", m)

	// echo message
	ret := &weixin.ReplyImage{
		MsgType:      m.MsgType,
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		PicUrl:       m.PicUrl,
		MediaId:      m.MediaId,
	}

	log.Debugf("%+v", ret)
	return ret
}

func echoMsgVoice(m *weixin.MsgVoice) interface{} {
	log.Debugf("%+v", m)

	// echo message
	ret := &weixin.ReplyVoice{
		MsgType:      m.MsgType,
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		MediaId:      m.MediaId,
	}

	log.Debugf("%+v", ret)
	return ret
}

func echoMsgVideo(m *weixin.MsgVideo) interface{} {
	log.Debugf("%+v", m)

	// MediaId ???
	ret := &weixin.ReplyVideo{
		MsgType:      m.MsgType,
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

func echoMsgShortVideo(m *weixin.MsgVideo) interface{} {
	log.Debugf("%+v", m)

	// MediaId ???
	ret := &weixin.ReplyVideo{
		MsgType:      weixin.MsgTypeVideo,
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

func echoMsgLocation(m *weixin.MsgLocation) interface{} {
	log.Debugf("%+v", m)

	// 回复音乐消息

	return nil
}

func echoMsgLink(m *weixin.MsgLink) interface{} {
	log.Debugf("%+v", m)

	// 回复图文消息

	return nil
}
