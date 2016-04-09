package weixin

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"

	"github.com/gotips/log"
)

// ValidateURL 验证 URL 以判断来源是否合法
func ValidateURL(token, timestamp, nonce, signature string) bool {
	tmpArr := []string{token, timestamp, nonce}
	sort.Strings(tmpArr)

	tmpStr := strings.Join(tmpArr, "")
	actual := fmt.Sprintf("%x", sha1.Sum([]byte(tmpStr)))

	log.Tracef("%s %s", tmpArr, actual)
	return actual == signature
}

// Signature 对加密的报文计算签名
func Signature(token, timestamp, nonce, encrypt string) string {
	tmpArr := []string{token, timestamp, nonce, encrypt}
	sort.Strings(tmpArr)

	tmpStr := strings.Join(tmpArr, "")
	actual := fmt.Sprintf("%x", sha1.Sum([]byte(tmpStr)))

	log.Tracef("%s %s", tmpArr, actual)
	return actual
}

// CheckSignature 验证加密的报文的签名
func CheckSignature(token, timestamp, nonce, encrypt, sign string) bool {
	return Signature(token, timestamp, nonce, encrypt) == sign
}
