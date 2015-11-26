package weixin

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
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
	if !validateURL(Token, timestamp, nonce, signature) {
		http.Error(w, "validate url error, request not from weixin?", http.StatusUnauthorized)
		return
	}

	switch r.Method {
	// 如果是 GET 请求，表示这是接入验证请求
	case "GET":
		w.Write([]byte(q.Get("echostr")))
	case "POST":
		HandleMessage(w, r)
	default:
		http.Error(w, "only GET or POST method allowed", http.StatusUnauthorized)
	}
}

func validateURL(token, timestamp, nonce, signature string) bool {
	tmpArr := []string{token, timestamp, nonce}
	sort.Strings(tmpArr)

	tmpStr := strings.Join(tmpArr, "")
	actual := fmt.Sprintf("%x", sha1.Sum([]byte(tmpStr)))

	log.Tracef("%s %s", tmpArr, actual)
	return actual == signature
}

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
