package weixin

import "fmt"

// 链接转换
const (
	ShortURLLongURL2ShortURL = "https://api.weixin.qq.com/cgi-bin/shorturl?access_token=%s"
)

// LongURL2Short 长链接转短链接接口
func LongURL2Short(longURL string) (shortURL string, err error) {
	js := fmt.Sprintf(`{"action":"long2short","long_url":"%s"}`, longURL)
	wapper := &struct {
		WXError
		ShortURL string `json:"short_url"`
	}{}
	url := fmt.Sprintf(ShortURLLongURL2ShortURL, AccessToken())
	err = Post(url, []byte(js), wapper)
	return wapper.ShortURL, err
}
