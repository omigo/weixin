package weixin

// TConfig 配置
type TConfig struct {
	originId       string // 原始ID
	appId          string // 应用ID
	appSecret      string // 应用密钥
	token          string // 令牌
	encodingAESKey string // 消息加解密密钥
}

var config TConfig

// Initialize 配置并初始化
func Initialize(originId, appId, appSecret, token, encodingAESKey string) {
	config = TConfig{originId, appId, appSecret, token, encodingAESKey}

	// refresh access token
}
