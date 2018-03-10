package weixin

import (
	"fmt"
	"net/url"
	"time"
)

// 网页授权获取用户基本信息
const (
	UserWebRedirectURL       = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	UserWebWebTokenURL       = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	UserWebRefreshTokenURL   = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	UserWebGetWebUserInfoURL = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s"
	UserWebCheckWebTokenURL  = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
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
func GenRedirectURL(redirectURL string, scope ScopeSNSAPI, state ...string) string {
	if len(state) == 0 {
		state = []string{time.Now().Format("20060102150405")}
	}
	redirectURL = url.QueryEscape(redirectURL)
	return fmt.Sprintf(UserWebRedirectURL, AppId, redirectURL, scope, state[0])
}

// WebToken 网页授权access_token
type WebToken struct {
	WXError
	AccessToken  string `json:"access_token"`      // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in"`        // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"`     // 用户刷新access_token
	OpenId       string `json:"openid"`            // 用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，也会产生一个用户和公众号唯一的OpenID
	Scope        string `json:"scope"`             // 用户授权的作用域，使用逗号（,）分隔
	UnionId      string `json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。详见：获取用户个人信息（UnionID机制）
}

// GetWebToken 获取网页授权 access_token
func GetWebToken(code string) (token *WebToken, err error) {
	url := fmt.Sprintf(UserWebWebTokenURL, AppId, AppSecret, code)
	token = &WebToken{}
	err = Get(url, token)
	return token, err
}

// RefreshWebToken 刷新网页授权 access_token
func RefreshWebToken(refreshToken string) (token *WebToken, err error) {
	url := fmt.Sprintf(UserWebRefreshTokenURL, AppId, refreshToken)
	token = &WebToken{}
	err = Get(url, token)
	return token, err
}

//WebUserInfo 网页授权获取用户基本信息
type WebUserInfo struct {
	WXError
	OpenId   string `json:"openid"`   // 用户的唯一标识
	NickName string `json:"nickname"` // 用户昵称
	Sex      int    `json:"sex"`      // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Province string `json:"province"` // 用户个人资料填写的省份
	City     string `json:"city"`     // 普通用户个人资料填写的城市
	Country  string `json:"country"`  // 国家，如中国为CN
	// 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，
	// 0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"` // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	// 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。详见：获取用户个人信息（UnionID机制）
	UnionId string `json:"unionid"`
}

// GetWebUserInfo 刷新网页授权 access_token
func GetWebUserInfo(webToken, openId string, lang ...Lang) (info *WebUserInfo, err error) {
	if len(lang) == 0 {
		lang = []Lang{LangZHCN}
	}
	url := fmt.Sprintf(UserWebGetWebUserInfoURL, webToken, openId, lang[0])
	info = &WebUserInfo{}
	err = Get(url, info)
	return info, err
}

// CheckWebToken 检验网页授权凭证（access_token）是否有效
func CheckWebToken(webToken, openId string) (err error) {
	url := fmt.Sprintf(UserWebCheckWebTokenURL, webToken, openId)
	err = Get(url, nil)
	return err
}
