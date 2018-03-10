package weixin

import (
	"fmt"
)

const (
	urlSession = `https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code`
)

type SessionResp struct {
	WXError
	UnionId    string `json:"unionid"`     // 用户唯一标识
	OpenId     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
}

func Session(appId, appSecret, code string) (*SessionResp, error) {
	resp := &SessionResp{}
	err := Get(fmt.Sprintf(urlSession, appId, appSecret, code), resp)
	return resp, err
}
