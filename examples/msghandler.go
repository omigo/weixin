package main

import (
	"encoding/json"

	"github.com/gotips/log"
	"github.com/arstd/weixin"
)

func defaultHandler(msg *weixin.Message) weixin.ReplyMsg {
	log.Debugf("%+v", msg)

	event := weixin.NewRecvEvent(msg)
	js, _ := json.Marshal(event)

	// echo message
	ret := &weixin.ReplyText{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   msg.CreateTime,
		Content:      string(js),
	}

	log.Debugf("replay message: %+v", ret)
	return ret
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
