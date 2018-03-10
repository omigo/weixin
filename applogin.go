package weixin

import (
	"fmt"
)

// 文档地址： https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1419317853&token=&lang=zh_CN
const (
	getTokenByCodeUrl = "https://api.weixin.qq.com/sns/oauth2/access_token?" +
		"appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	getUserInfobyTokenUrl = "https://api.weixin.qq.com/sns/userinfo?" +
		"access_token=%s&openid=%s"
)

type WXToken struct {
	WXError
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionId      string `json:"unionid"`
}

func GetTokenByCode(appId, appSecret, code string) (*WXToken, error) {
	u := fmt.Sprintf(getTokenByCodeUrl, appId, appSecret, code)
	resp := new(WXToken)
	err := Get(u, resp)
	return resp, err
}

type WXUserInfo struct {
	WXError
	OpenId     string `json:"openid"`
	UnionId    string `json:"unionid"`
	Nickname   string `json:"nickname"`
	Sex        int8   `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	HeadimgUrl string `json:"headimgurl"`
}

func GetUserInfoByToken(accessToken, openId string) (*WXUserInfo, error) {
	u := fmt.Sprintf(getUserInfobyTokenUrl, accessToken, openId)
	resp := new(WXUserInfo)
	err := Get(u, resp)
	return resp, err
}
