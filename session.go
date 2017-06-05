package weixin

import (
	"fmt"

	"github.com/arstd/log"
	"github.com/arstd/mooo/misc/fetch"
)

const (
	urlSession = `https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code`
)

type SessionResp struct {
	WeixinError
	UnionId    string `json:"unionid"`     // 用户唯一标识
	OpenId     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
}

func Session(appId, appSecret, code string) (*SessionResp, error) {
	resp := &SessionResp{}
	err := fetch.Get(fmt.Sprintf(urlSession, appId, appSecret, code), resp)
	if err != nil {
		return nil, err
	}
	if resp.ErrCode != 0 {
		log.Error(resp)
		return nil, &resp.WeixinError
	}
	return resp, nil
}
