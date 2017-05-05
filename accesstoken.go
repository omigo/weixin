package weixin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/arstd/log"
)

// tick := time.Tick(7 * time.Second)
const refreshTimeout = 30 * time.Minute
const tokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

type accessToken struct {
	AccessToken string `json:"access_token"` //	获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   //	凭证有效时间，单位：秒
	mutex       sync.RWMutex
}

// AccessToken 取最新的 access_token，必须使用这个接口取，内部已经加锁
var AccessToken func() string

// RefreshAccessToken 定时刷新 access_token
func RefreshAccessToken(appId, appSecret string) {
	// 内部变量，外部不可以调用
	var _token = &accessToken{}

	AccessToken = func() string {
		_token.mutex.RLock()
		defer _token.mutex.RUnlock()

		return _token.AccessToken
	}

	go func() {
		url := fmt.Sprintf(tokenURL, appId, appSecret)

		tick := time.Tick(refreshTimeout)
		for {
			new := refresh(url)

			log.Debugf("old access token %+v", _token)
			log.Debugf("new access token %+v", new)

			_token.mutex.Lock()
			_token.AccessToken = new.AccessToken
			_token.ExpiresIn = new.ExpiresIn
			_token.mutex.Unlock()

			<-tick // 等待下一个时钟周期到来
		}
	}()
}

func refresh(url string, ns ...int) (new *accessToken) {
	n := 0
	if len(ns) > 0 {
		n = ns[0]
	}

	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			time.Sleep(3 * time.Minute)
			if n < 9 {
				n++
				new = refresh(url, n)
			}
		}
	}()

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resp.Body.Close()

	new = &accessToken{}
	err = json.Unmarshal(body, new)
	if err != nil {
		return
	}

	return new
}
