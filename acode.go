package weixin

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/arstd/log"
)

const (
	urlWxAcode       = "https://api.weixin.qq.com/wxa/getwxacode?access_token=%s"
	jsWxAcodeReqBody = `{"path":"%s","width":430,"auto_color":true}`
)
const (
	urlWxAcodeUnlimit       = "http://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s"
	jsWxAcodeUnlimitReqBody = `{"scene":"%s","width":430,"auto_color":true}`
)

func Acode(path string) (io.ReadCloser, error) {
	atoken := AccessToken()

	url := fmt.Sprintf(urlWxAcode, atoken)
	in := fmt.Sprintf(jsWxAcodeReqBody, path)
	resp, err := http.Post(url, "application/json", bytes.NewBufferString(in))
	if err != nil {
		return nil, err
	}
	log.Debug(resp.Header.Get("Content-Type"))

	return resp.Body, nil
}

func AcodeUnlimit(scene string) ([]byte, error) {
	atoken := AccessToken()

	url := fmt.Sprintf(urlWxAcodeUnlimit, atoken)
	in := fmt.Sprintf(jsWxAcodeUnlimitReqBody, scene)
	resp, err := http.Post(url, "application/json", bytes.NewBufferString(in))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
