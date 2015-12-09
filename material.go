package weixin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
)

// 素材管理
const (
	MaterialUploadTemporaryURL = "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
	MaterialGetTemporaryURL    = "https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
)

// MediaType 媒体文件类型
type MediaType string

// 微信支持的媒体文件类型
const (
	MediaTypeImage MediaType = "image"
	MediaTypeVoice MediaType = "voice"
	MediaTypeVideo MediaType = "video"
	MediaTypeThumb MediaType = "thumb"
)

// UploadTemporaryMaterial 新增临时素材
func UploadTemporaryMaterial(mtype MediaType, file *os.File) (mediaId string, createAt int, err error) {
	url := fmt.Sprintf(MaterialUploadTemporaryURL, AccessToken(), mtype)
	wapper := &struct {
		Type     MediaType `json:"type"`
		MediaId  string    `json:"media_id"`
		CreateAt int       `json:"created_at"`
	}{}
	err = Upload(url, "media", file, wapper)
	return wapper.MediaId, wapper.CreateAt, err
}

// GetTemporaryMaterial 新增临时素材
func GetTemporaryMaterial(mediaId string) (filename string, body []byte, err error) {
	url := fmt.Sprintf(MaterialGetTemporaryURL, AccessToken(), mediaId)
	resp, err := http.Get(url)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	params := make(map[string]string)
	if cd := resp.Header.Get("Content-Disposition"); cd == "" {
		return "", nil, errors.New("missing Content-Disposition header")
	} else if _, params, err = mime.ParseMediaType(cd); err != nil {
		return "", nil, fmt.Errorf("parse Content-Disposition header fail: %s", err.Error())
	} else if filename = params["filename"]; filename == "" {
		return "", nil, errors.New("no filename in Content-Disposition header")
	}

	// // 取文件名
	// disp := resp.Header.Get("Content-Disposition")
	// re := regexp.MustCompile(`filename="(.+?)"`)
	// matched := re.FindStringSubmatch(disp)
	// if len(matched) != 2 {
	// 	filename = mediaId
	// } else {
	// 	filename = matched[1]
	// }

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}
	if resp.StatusCode != http.StatusOK {
		wxerr := &WeixinError{}
		err = json.Unmarshal(body, wxerr)
		if err != nil {
			return "", nil, err
		}
		return "", nil, wxerr
	}
	return filename, body, nil

}
