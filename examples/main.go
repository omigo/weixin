package main

import (
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
