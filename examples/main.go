package main

import (
	"net/http"

	"github.com/omigo/log"
	"github.com/omigo/weixin"
)

func main() {
	addr := ":3080"

	weixin.Initialize(originId, appId, appSecret, token, encodingAESKey)
	weixin.MsgTextHandler = echoText   // 注册文本消息处理器
	weixin.MsgImageHandler = echoImage // 注册图片消息处理器
	http.HandleFunc("/weixin", weixin.HandleAccess)

	log.Debugf("server is running at %s", addr)
	http.ListenAndServe(addr, nil)
}

func echoText(msg *weixin.MsgText) interface{} {
	log.Debugf("%+v", msg)

	// echo message
	ret := &weixin.ReplyText{
		MsgType:      msg.MsgType,
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   msg.CreateTime,
		Content:      msg.Content,
	}

	log.Debugf("%+v", ret)
	return ret
}

func echoImage(msg *weixin.MsgImage) interface{} {
	log.Debugf("%+v", msg)

	// echo message
	ret := &weixin.ReplyImage{
		MsgType:      msg.MsgType,
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   msg.CreateTime,
		PicUrl:       msg.PicUrl,
		MediaId:      msg.MediaId,
	}

	log.Debugf("%+v", ret)
	return ret
}
