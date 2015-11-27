package weixin

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/omigo/log"
)

// EncryptMsg 加密报文
func EncryptMsg(msg []byte, aesKey []byte, appId string) (b64Enc string, err error) {
	// 拼接完整报文
	src := SpliceFullMsg(msg, appId)

	// AES CBC 加密报文
	dst, err := AESCBCEncrypt(src, aesKey, aesKey[:aes.BlockSize])
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(dst), nil
}

// DecryptMsg 解密报文
func DecryptMsg(b64Enc string, aesKey []byte, appId string) (msg []byte, err error) {
	log.Tracef("b64Enc: %s", b64Enc)
	enc, err := base64.StdEncoding.DecodeString(b64Enc)
	if err != nil {
		return nil, err
	}

	// AES CBC 解密报文
	src, err := AESCBCDecrypt(enc, aesKey, aesKey[:aes.BlockSize])
	if err != nil {
		return nil, err
	}

	log.Tracef("full message: %s", src)

	_, _, msg, appId2 := ParseFullMsg(src)
	if appId2 != appId {
		return nil, fmt.Errorf("expected appId %s, but %s", appId, appId2)
	}

	return msg, nil
}

// SpliceFullMsg 拼接完整报文，
// AES加密的buf由16个字节的随机字符串、4个字节的msg_len(网络字节序)、msg和$AppId组成，
// 其中msg_len为msg的长度，$AppId为公众帐号的AppId
func SpliceFullMsg(msg []byte, appId string) (fullMsg []byte) {
	// 16个字节的随机字符串
	randBytes := RandBytes(16)

	// 4个字节的msg_len(网络字节序)
	msgLen := len(msg)
	lenBytes := []byte{
		byte(msgLen >> 24 & 0xFF),
		byte(msgLen >> 16 & 0xFF),
		byte(msgLen >> 8 & 0xFF),
		byte(msgLen & 0xFF),
	}

	return bytes.Join([][]byte{randBytes, lenBytes, msg, []byte(appId)}, nil)
}

// ParseFullMsg 从完整报文中解析出消息内容，
// AES加密的buf由16个字节的随机字符串、4个字节的msg_len(网络字节序)、msg和$AppId组成，
// 其中msg_len为msg的长度，$AppId为公众帐号的AppId
func ParseFullMsg(fullMsg []byte) (randBytes []byte, msgLen int, msg []byte, appId string) {
	randBytes = fullMsg[:16]

	msgLen = (int(fullMsg[16]) << 24) |
		(int(fullMsg[17]) << 16) |
		(int(fullMsg[18]) << 8) |
		int(fullMsg[19])
	// log.Tracef("msgLen=[% x]=(%d %d %d %d)=%d", fullMsg[16:20], (int(fullMsg[16]) << 24),
	// 	(int(fullMsg[17]) << 16), (int(fullMsg[18]) << 8), int(fullMsg[19]), msgLen)

	msg = fullMsg[20 : 20+msgLen]

	appId = string(fullMsg[20+msgLen:])

	return
}

// RandBytes 产生 size 个长度的随机字节
func RandBytes(size int) (r []byte) {
	r = make([]byte, size)
	_, err := rand.Read(r)
	if err != nil {
		// 忽略错误，不影响其他逻辑，仅仅打印日志
		log.Warnf("rand read error: %s", err)
	}
	return r
}
