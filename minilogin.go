package weixin

import (
	"fmt"
)

const (
	urlMiniLogin = `https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code`
)

type MiniLoginResp struct {
	WXError
	UnionId    string `json:"unionid"`     // 用户唯一标识
	OpenId     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
}

func MiniLogin(appId, appSecret, code string) (*MiniLoginResp, error) {
	resp := &MiniLoginResp{}
	err := Get(fmt.Sprintf(urlMiniLogin, appId, appSecret, code), resp)
	return resp, err
}
