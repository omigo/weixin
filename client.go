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
func Post(url string, jsonBytes []byte) (err error) {
	log.Debugf("url=%s, body=%s", url, jsonBytes)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
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
