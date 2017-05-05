package main

import (
	"net/http"

	"github.com/arstd/log"
	"github.com/arstd/weixin"
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
	// 模版消息发送结果通知事件
	weixin.EventTemplateSendJobFinishHandler = EventTemplateSendJobFinishHandler
	weixin.EventDefaultHandler = EventDefaultHandler // 注册默认处理器

	http.HandleFunc("/weixin", weixin.HandleAccess)
	http.Handle("/", http.FileServer(http.Dir("examples/static")))
	// http.Handle("/admin/", http.StripPrefix("/admin/", http.FileServer(http.Dir("admin"))))

	log.Debugf("server is running at %s", addr)
	http.ListenAndServe(addr, nil)
}
