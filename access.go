package weixin

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)
import "github.com/omigo/log"

// HandleAccess 接入微信公众平台开发，并接口来自微信服务器的消息
func HandleAccess(w http.ResponseWriter, r *http.Request) {
	log.Debugf("%s", r.URL)

	q := r.URL.Query()
	signature := q.Get("signature")
	timestamp := q.Get("timestamp")
	nonce := q.Get("nonce")

	// 每次都验证 URL，以判断来源是否合法
	if !ValidateURL(Token, timestamp, nonce, signature) {
		http.Error(w, "validate url error, request not from weixin?", http.StatusUnauthorized)
		return
	}

	switch r.Method {
	// 如果是 GET 请求，表示这是接入验证请求
	case "GET":
		w.Write([]byte(q.Get("echostr")))
	case "POST":
		processMessage(w, r)
	default:
		http.Error(w, "only GET or POST method allowed", http.StatusUnauthorized)
	}
}

// 处理所有来自微信的消息，已经验证过 URL 和 Method 了
func processMessage(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	timestamp := q.Get("timestamp")
	nonce := q.Get("nonce")
	encryptType := q.Get("encrypt_type")
	msgSignature := q.Get("msg_signature")

	// 读取报文
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		http.Error(w, "read body error", http.StatusNotAcceptable)
		return
	}
	r.Body.Close()
	log.Debugf("from weixin: %s", body)

	msg, err := parseBody(encryptType, timestamp, nonce, msgSignature, body)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 处理消息
	reply := HandleMessage(msg)

	// 如果返回为 nil，表示不需要回复，结束
	if reply == nil {
		return
	}

	// 如果返回不为 nil，表示需要回复
	ret, err := packReply(reply, encryptType, timestamp, nonce)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Debugf("to weixin: %s", ret)
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.Write(ret)
}

func parseBody(encryptType, timestamp, nonce, msgSignature string, body []byte) (msg *Message, err error) {
	msg = &Message{}
	// 如果报文被加密了，先要验签解密
	if encryptType == "aes" {
		encMsg := &EncMessage{}
		// 解析加密的 xml
		err = xml.Unmarshal(body, encMsg)
		if err != nil {
			return nil, err
		}
		msg.ToUserName = encMsg.ToUserName
		msg.Encrypt = encMsg.Encrypt

		if !CheckSignature(Token, timestamp, nonce, encMsg.Encrypt, msgSignature) {
			return nil, errors.New("check signature error")
		}

		body, err = DecryptMsg(encMsg.Encrypt, EncodingAESKey, AppId)
		if err != nil {
			return nil, err
		}
		log.Debugf("receive: %s", body)
	}

	// 解析 xml
	err = xml.Unmarshal(body, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func packReply(reply ReplyMsg, encryptType, timestamp, nonce string) (ret []byte, err error) {
	switch reply.(type) {
	case *ReplyText:
		reply.SetMsgType(MsgTypeText)
	case *ReplyImage:
		reply.SetMsgType(MsgTypeImage)
	case *ReplyVoice:
		reply.SetMsgType(MsgTypeVoice)
	case *ReplyVideo:
		reply.SetMsgType(MsgTypeVideo)
	case *ReplyMusic:
		reply.SetMsgType(MsgTypeMusic)
	case *ReplyNews:
		reply.SetMsgType(MsgTypeNews)
	default:
		panic("unexpected custom message type")
	}

	ret, err = xml.MarshalIndent(reply, "", "  ")
	if err != nil {
		return nil, err
	}
	log.Debugf("replay: %s", ret)

	// 如果接收的消息加密了，那么回复的消息也需要签名加密
	if encryptType == "aes" {
		b64Enc, err := EncryptMsg(ret, EncodingAESKey, AppId)
		if err != nil {
			return nil, err
		}
		encMsg := EncMessage{
			Encrypt:      b64Enc,
			MsgSignature: Signature(Token, timestamp, nonce, b64Enc),
			TimeStamp:    timestamp,
			Nonce:        nonce, // 随机数
		}
		ret, err = xml.MarshalIndent(encMsg, "", "    ")
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}
