package weixin

import (
	"encoding/base64"
	"os"
	"regexp"
	"time"

	"github.com/omigo/log"
)

// TConfig 配置
var (
	OriginId       string // 原始ID
	AppId          string // 应用ID
	AppSecret      string // 应用密钥
	Token          string // 令牌
	EncodingAESKey []byte // 消息加解密密钥
)

// Initialize 配置并初始化
func Initialize(originId, appId, appSecret, token, encodingAESKey string) {
	if matched, err := regexp.MatchString("^gh_[0-9a-f]{12}$", originId); err != nil || !matched {
		log.Fatalf("originId format error: %s", err)
	}
	if matched, err := regexp.MatchString("^wx[0-9a-f]{16}$", appId); err != nil || !matched {
		log.Fatalf("appId format error: %s", err)
	}
	if matched, err := regexp.MatchString("^[0-9a-f]{32}$", appSecret); err != nil || !matched {
		log.Fatalf("appSecret format error: %s", err)
	}
	if matched, err := regexp.MatchString("^[0-9a-zA-Z]{3,32}$", token); err != nil || !matched {
		log.Fatalf("token format error: %s", err)
	}
	if matched, err := regexp.MatchString("^[0-9a-zA-Z]{43}$", encodingAESKey); err != nil || !matched {
		log.Fatalf("encodingAESKey format error: %s", err)
	}

	OriginId = originId   // 原始ID
	AppId = appId         // 应用ID
	AppSecret = appSecret // 应用密钥
	Token = token         // 令牌

	var err error
	EncodingAESKey, err = base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		log.Fatalf("appSecret config error: %s", err)
		os.Exit(1)
	}

	// refresh access token
	RefreshAccessToken(AppId, AppSecret)
	time.Sleep(1 * time.Second) // waiting to refresh token
}
