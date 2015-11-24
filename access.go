package weixin

import (
	"crypto/sha1"
	"fmt"
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
	if !validateURL(config.token, timestamp, nonce, signature) {
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
