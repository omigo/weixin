package weixin

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// 素材管理
const (
	MaterialUploadTemporaryURL = "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
	MaterialGetTemporaryURL    = "https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
	MaterialAddNewsURL         = "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=%s"
	MaterialUploadImg          = "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"
	MaterialUploadNewsURL      = "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s"
	MaterialGetNewsURL         = "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%s"
	MaterialDeleteNewsURL      = "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s"
	MaterialUpdateNewsURL      = "https://api.weixin.qq.com/cgi-bin/material/update_news?access_token=%s"
	MaterialCountURL           = "https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=%s"
	MaterialBatchGetNewsURL    = "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%s"
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
		WXError
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
	return Download(url)
}

// Article 永久图文素材
type Article struct {
	Title            string `json:"title"`              // 标题
	ThumbMediaId     string `json:"thumb_media_id"`     // 图文消息的封面图片素材id（必须是永久mediaID）
	Author           string `json:"author"`             // 作者
	Digest           string `json:"digest"`             // 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空
	ShowCoverPic     int    `json:"show_cover_pic"`     // 是否显示封面，0为false，即不显示，1为true，即显示
	Content          string `json:"content"`            // 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS
	URL              string `json:"url"`                // 图文页的URL
	ContentSourceURL string `json:"content_source_url"` // 图文消息的原文地址，即点击“阅读原文”后的URL
}

// AddNews 新增永久图文素材
func AddNews(as []Article) (mediaId string, err error) {
	if len(as) == 0 {
		return "", errors.New("articles is blank")
	}

	js, err := json.Marshal(as)
	if err != nil {
		return "", err
	}

	body := fmt.Sprintf(`{"articles":%s}`, js)
	url := fmt.Sprintf(MaterialAddNewsURL, AccessToken())
	wrapper := &struct {
		WXError
		MediaId string `json:"media_id"`
	}{}
	err = Post(url, []byte(body), wrapper)
	return wrapper.MediaId, err
}

// UploadImg 上传图文消息内的图片获取URL
func UploadImg(file *os.File) (u string, err error) {
	url := fmt.Sprintf(MaterialUploadImg, AccessToken())
	wapper := &struct {
		WXError
		URL string `json:"url"`
	}{}
	err = Upload(url, "media", file, wapper)
	return wapper.URL, err
}

// UploadNews 新增其他类型永久素材
func UploadNews(mtype MediaType, file *os.File) (mediaId, u string, err error) {
	url := fmt.Sprintf(MaterialUploadNewsURL, AccessToken(), mtype)
	wapper := &struct {
		WXError
		MediaId string `json:"media_id"`
		URL     string `json:"url"`
	}{}
	err = Upload(url, "media", file, wapper)
	return wapper.MediaId, wapper.URL, err
}

// UploadVideo 新增视频类型永久素材
func UploadVideo(mtype MediaType, title, desc string, file *os.File) (mediaId, u string, err error) {
	url := fmt.Sprintf(MaterialUploadNewsURL, AccessToken(), mtype)
	wapper := &struct {
		WXError
		MediaId string `json:"media_id"`
		URL     string `json:"url"`
	}{}
	desc = fmt.Sprintf(`{"title":"%s", "introduction":"%s"}`, title, desc)
	err = Upload(url, "media", file, wapper, desc)
	return wapper.MediaId, wapper.URL, err
}

// News 永久素材
type News struct {
	WXError
	Title       string    `json:"title"` // 图文消息的标题
	Description string    `json:"description"`
	DownURL     string    `json:"down_url"`
	NewsItem    []Article `json:"news_item"`
}

// GetNews 获取临时素材
func GetNews(mediaId string) (ret *News, err error) {
	url := fmt.Sprintf(MaterialGetNewsURL, AccessToken())
	body := fmt.Sprintf(`{"media_id":"%s"}`, mediaId)
	ret = &News{}
	err = Post(url, []byte(body), ret)
	return ret, err
}

// DeleteNews 删除永久素材
func DeleteNews(mediaId string) (err error) {
	url := fmt.Sprintf(MaterialDeleteNewsURL, AccessToken())
	body := fmt.Sprintf(`{"media_id":"%s"}`, mediaId)
	return Post(url, []byte(body), nil)
}

// UpdateNewsReq 修改永久图文素材
type UpdateNewsReq struct {
	MediaId  string    `json:"media_id"` // 要修改的图文消息的id
	Index    string    `json:"index"`    // 要更新的文章在图文消息中的位置（多图文消息时，此字段才有意义），第一篇为0
	Articles []Article `json:"articles"` // 要修改的图文消息的id
}

// UpdateNews 修改永久图文素材
func UpdateNews(news *UpdateNewsReq) (err error) {
	url := fmt.Sprintf(MaterialDeleteNewsURL, AccessToken())
	return Post(url, news, nil)
}

// MaterialCount 素材总数
type MaterialCount struct {
	WXError
	VoiceCount int `json:"voice_count"` // 语音总数量
	VideoCount int `json:"video_count"` // 视频总数量
	ImageCount int `json:"image_count"` // 图片总数量
	NewsCount  int `json:"news_count"`  // 图文总数量
}

// GetMaterialCount 获取素材总数
func GetMaterialCount() (mc *MaterialCount, err error) {
	url := fmt.Sprintf(MaterialCountURL, AccessToken())
	mc = &MaterialCount{}
	err = Get(url, mc)
	return mc, err
}

// NewsList 素材列表
type NewsList struct {
	WXError
	TotalCount string `json:"total_count"` // 该类型的素材的总数
	ItemCount  string `json:"item_count"`  // 本次调用获取的素材的数量
	Item       []struct {
		MediaId    string `json:"media_id"`
		UpdateTime string `json:"update_time"` // 这篇图文消息素材的最后更新时间
		Name       string `json:"name"`        // 文件名称
		URL        string `json:"url"`         // 图文页的URL，或者，当获取的列表是图片素材列表时，该字段是图片的URL
		Content    struct {
			NewsItem []Article `json:"news_item"`
		} `json:"content"` // 本次调用获取的素材的数量
	} `json:"item"` // 多图文消息会在此处有多篇文章
}

// BatchGetNews 获取素材列表
func BatchGetNews(mtype MediaType, offset int, count int) (ret *NewsList, err error) {
	url := fmt.Sprintf(MaterialBatchGetNewsURL, AccessToken())
	body := fmt.Sprintf(`{"type":%s,"offset":%d,"count":%d}`, mtype, offset, count)
	ret = &NewsList{}
	err = Post(url, []byte(body), ret)
	return ret, err
}
