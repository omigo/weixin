package weixin

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/omigo/log"
)

// Post 工具类, POST json 并返回 error
func Post(url string, js []byte) (err error) {
	wxerr := &WeixinError{}
	err = PostUnmarshal(url, js, wxerr)
	if err != nil {
		return err
	}

	if wxerr.ErrCode == WeixinErrCodeSuccess {
		return nil
	}

	// if wxerr.ErrCode == WeixinErrCodeSystemBusy {
	//
	// }
	log.Errorf("weixin error %d: %s", wxerr.ErrCode, wxerr.ErrMsg)
	return wxerr
}

// PostMarshalUnmarshal 工具类, POST 编组并解析返回的报文，返回 error
func PostMarshalUnmarshal(url string, v interface{}, ret interface{}) (err error) {
	js, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return PostUnmarshal(url, js, ret)
}

// PostUnmarshal 工具类, POST json 并解析返回的报文，返回 error
func PostUnmarshal(url string, js []byte, ret interface{}) (err error) {
	log.Debugf("url=%s, body=%s", url, js)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(js))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Debugf("response: %s", body)

	err = json.Unmarshal(body, ret)
	if err != nil {
		return err
	}

	return nil
}

// Upload 工具类, 上传文件
func Upload(url string, file *os.File) (err error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	//关键的一步操作
	fw, err := w.CreateFormField(file.Name())
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}
	contentType := w.FormDataContentType()
	w.Close()

	resp, err := http.Post(url, contentType, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Debugf("response: %s", body)

	wxerr := &WeixinError{}
	err = json.Unmarshal(body, wxerr)
	if err != nil {
		return err
	}

	if wxerr.ErrCode == WeixinErrCodeSuccess {
		return nil
	}

	// if wxerr.ErrCode == WeixinErrCodeSystemBusy {
	//
	// }

	log.Errorf("weixin error %d: %s", wxerr.ErrCode, wxerr.ErrMsg)
	return wxerr
}
