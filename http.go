package weixin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/arstd/log"
)

type wxResp interface {
	// 似有的 error 方法，保证外部(其他包)定义的 struct 只能内嵌
	// WXError 的才能实现这个方法，才能作为当前包 http 方法的参数
	error() error
}

type WXError struct {
	ErrCode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

func (e *WXError) error() error {
	if e.ErrCode == 0 {
		return nil
	}
	return e
}

func (e *WXError) Error() string {
	return fmt.Sprintf("weixin: (%d)%s", e.ErrCode, e.Errmsg)
}

var client = http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 128,
	},
}

// GetUnmarshal HTTP 工具类, GET 并解析返回的报文，如果有错误，返回 error
func Get(url string, wxr wxResp) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	return parseWXResp(resp, wxr)
}

const contentType = "application/json"

// Post HTTP 工具类, POST 并解析返回的报文，如果有错误，返回 error
func Post(url string, v interface{}, wxr wxResp) (err error) {
	var js []byte
	if _, ok := v.([]byte); !ok {
		js, err = json.Marshal(v)
		if err != nil {
			return err
		}
	}

	resp, err := client.Post(url, contentType, bytes.NewBuffer(js))
	if err != nil {
		return err
	}
	return parseWXResp(resp, wxr)
}

// Upload 工具类, 上传文件
func Upload(url, fieldName string, file *os.File, wxr wxResp, desc ...string) (err error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	//关键的一步操作
	// fw, err := w.CreateFormField(file.Name())
	fw, err := w.CreateFormFile(fieldName, file.Name())
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}
	contentType := w.FormDataContentType()
	if len(desc) > 0 {
		w.WriteField("description", desc[0])
	}
	w.Close()

	log.Debugf("url=%s, fieldName=%s, fileName=%s", url, fieldName, file.Name())
	resp, err := client.Post(url, contentType, buf)
	if err != nil {
		return err
	}

	return parseWXResp(resp, wxr)
}

func Download(url string) (filename string, body []byte, err error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", nil, err
	}

	var params map[string]string
	if cd := resp.Header.Get("Content-Disposition"); cd == "" {
		return "", nil, errors.New("missing Content-Disposition header")
	} else if _, params, err = mime.ParseMediaType(cd); err != nil {
		return "", nil, err
	} else if filename = params["filename"]; filename == "" {
		return "", nil, errors.New("no filename in Content-Disposition header")
	}

	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return filename, body, err
}

func parseWXResp(resp *http.Response, wxr wxResp) error {
	js, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	log.Trace("%s", js)
	if wxr == nil {
		wxr = &WXError{}
	}
	err = json.Unmarshal(js, wxr)
	if err != nil {
		return err
	}

	return wxr.error()
}
