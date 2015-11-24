package main

import (
	"net/http"

	"github.com/omigo/log"
	"github.com/omigo/weixin"
)

const (
	originId       = "gh_××××××××××"
	appId          = "wx****************"
	appSecret      = "*******************************"
	token          = "*******************************"
	encodingAESKey = "*******************************************"
)

func main() {
	addr := ":3080"

	weixin.Initialize(originId, appId, appSecret, token, encodingAESKey)
	weixin.Register(weixin.MsgTypeText, echo)
	http.HandleFunc("/weixin", weixin.HandleAccess)

	log.Debugf("server is running at %s", addr)
	http.ListenAndServe(addr, nil)
}

func echo(msg *weixin.Message) (ret *weixin.Message) {
	log.Debugf("%+v", msg)

	// echo message
	ret = &weixin.Message{}
	*ret = *msg
	ret.ToUserName, ret.FromUserName = msg.FromUserName, msg.ToUserName
	ret.MsgId = ""

	log.Debugf("%+v", ret)
	return ret
}
