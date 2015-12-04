package weixin

import (
	"fmt"
	"time"
)

// 网页授权获取用户基本信息
const (
	UserWebRedirectURL     = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	UserWebJsTokenURL      = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	UserWebRefreshTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
)

// ScopeSNSAPI 应用授权作用域类型
type ScopeSNSAPI string

// 应用授权作用域
const (
	// 不弹出授权页面，直接跳转，只能获取用户openid
	ScopeSNSAPIBase = "snsapi_base"
	// 弹出授权页面，可通过openid拿到昵称、性别、所在地。并且，即使在未关注的情况下，只要用户授权，也能获取其信息
	ScopeSNSAPIUserInfo = "snsapi_userinfo"
)

// GenRedirectURL 生成跳转链接
// state 非必填，默认当前时间 重定向后会带上state参数，开发者可以填写a-zA-Z0-9的参数值，最多128字节
func GenRedirectURL(url string, scope ScopeSNSAPI, state ...string) string {
	if len(state) == 0 {
		state = []string{time.Now().Format("20060102150405")}
	}
	return fmt.Sprintf(UserWebRedirectURL, AppId, url, scope, state[0])
}

// JsToken 网页授权access_token
type JsToken struct {
	WeixinError
	AccessToken  string `json:"access_token"`      // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in"`        // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"`     // 用户刷新access_token
	OpenId       string `json:"openid"`            // 用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，也会产生一个用户和公众号唯一的OpenID
	Scope        string `json:"scope"`             // 用户授权的作用域，使用逗号（,）分隔
	UnionId      string `json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。详见：获取用户个人信息（UnionID机制）
}

// GetJsToken 获取网页授权 access_token
func GetJsToken(code string) (token *JsToken, err error) {
	url := fmt.Sprintf(UserWebJsTokenURL, AppId, AppSecret, code)
	token = &JsToken{}
	err = GetUnmarshal(url, token)
	if err != nil {
		return nil, err
	}
	if token.ErrCode != WeixinErrCodeSuccess {
		return nil, token
	}
	return token, err
}

// RefreshToken 刷新网页授权 access_token
func RefreshToken(refreshToken string) (token *JsToken, err error) {
	url := fmt.Sprintf(UserWebRefreshTokenURL, AppId, refreshToken)
	token = &JsToken{}
	err = GetUnmarshal(url, token)
	if err != nil {
		return nil, err
	}
	if token.ErrCode != WeixinErrCodeSuccess {
		return nil, token
	}
	return token, err
}
