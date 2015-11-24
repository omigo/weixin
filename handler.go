package weixin

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/omigo/log"
)

// HandleMessage 处理所有来自微信的消息
func HandleMessage(w http.ResponseWriter, r *http.Request) {
	// 读取报文
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		http.Error(w, "read body error", http.StatusNotAcceptable)
		return
	}
	log.Debugf("receive: %s", body)

	// 解析 xml
	msg := &Message{}
	err = xml.Unmarshal(body, msg)
	if err != nil {
		log.Error(err)
		http.Error(w, "unmarshal xml error", http.StatusBadRequest)
		return
	}

	// 处理消息
	reply := processMessage(msg)

	// 如果返回为 nil，表示不需要回复，结束
	if reply == nil {
		return
	}

	// 如果返回不为 nil，表示需要回复
	rbody, err := xml.MarshalIndent(reply, "", "  ")
	if err != nil {
		log.Error(err)
		http.Error(w, "system error", http.StatusInternalServerError)
		return
	}

	log.Debugf("replay: %s", rbody)
	w.Write(rbody)
}

// MessageHandler 处理各类消息
type MessageHandler func(*Message) *Message

// handlers 注册各类消息的处理器
var handlers = make(map[MsgType]MessageHandler)

// Register 注册一个消息处理器
func Register(msgType MsgType, handler MessageHandler) {
	handlers[msgType] = handler
}

func processMessage(msg *Message) (ret *Message) {
	h, ok := handlers[msg.MsgType]
	if !ok {
		log.Warnf("unregister MsgType: %s", msg.MsgType)
		return nil
	}

	ret = h(msg)

	return ret
}
