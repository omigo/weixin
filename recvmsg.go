package weixin

// RecvText 接收文本消息
type RecvText struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间（整型）
	MsgType      MsgType // text
	Content      string  // 文本消息内容
	MsgId        int     // 消息id，64位整型
}

// RecvText 把通用 struct 转化成相应类型的 struct
func (m *Message) RecvText() *RecvText {
	return &RecvText{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Content:      m.Content,
		MsgId:        m.MsgId,
	}
}

// RecvImage 接收图片消息
type RecvImage struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间（整型）
	MsgType      MsgType // image
	PicUrl       string  // 图片链接
	MediaId      string  // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据
	MsgId        int     // 消息id，64位整型
}

// RecvImage 把通用 struct 转化成相应类型的 struct
func (m *Message) RecvImage() *RecvImage {
	return &RecvImage{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		PicUrl:       m.PicUrl,
		MediaId:      m.MediaId,
		MsgId:        m.MsgId,
	}
}

// RecvVoice 接收视频/小视频消息
type RecvVoice struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间（整型）
	MsgType      MsgType // voice
	MediaId      string  // 语音消息媒体id，可以调用多媒体文件下载接口拉取数据
	Format       string  // 语音格式，如amr，speex等
	Recongnition string  // 语音识别结果，使用UTF8编码
	MsgId        int     // 消息id，64位整型
}

// RecvVoice 把通用 struct 转化成相应类型的 struct
func (m *Message) RecvVoice() *RecvVoice {
	return &RecvVoice{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		MediaId:      m.MediaId,
		Format:       m.Format,
		Recongnition: m.Recongnition,
		MsgId:        m.MsgId,
	}
}

// RecvVideo 接收图片消息
type RecvVideo struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间（整型）
	MsgType      MsgType // video
	MediaId      string  // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据
	ThumbMediaId string  // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据
	MsgId        int     // 消息id，64位整型
}

// RecvVideo 把通用 struct 转化成相应类型的 struct
func (m *Message) RecvVideo() *RecvVideo {
	return &RecvVideo{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		MediaId:      m.MediaId,
		ThumbMediaId: m.ThumbMediaId,
		MsgId:        m.MsgId,
	}
}

// RecvLocation 接收地理位置消息
type RecvLocation struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间（整型）
	MsgType      MsgType // location
	LocationX    float64 `xml:"Location_X,omitempty"` // 地理位置维度
	LocationY    float64 `xml:"Location_Y,omitempty"` // 地理位置经度
	Scale        int     // 地图缩放大小
	Label        string  // 地理位置信息
	MsgId        int     // 消息id，64位整型
}

// RecvLocation 把通用 struct 转化成相应类型的 struct
func (m *Message) RecvLocation() *RecvLocation {
	return &RecvLocation{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		LocationX:    m.LocationX,
		LocationY:    m.LocationY,
		Scale:        m.Scale,
		Label:        m.Label,
		MsgId:        m.MsgId,
	}
}

// RecvLink 接收链接消息
type RecvLink struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间（整型）
	MsgType      MsgType // location
	Title        string  // 消息标题
	Description  string  // 消息描述
	Url          string  // 消息链接
	MsgId        int     // 消息id，64位整型
}

// RecvLink 把通用 struct 转化成相应类型的 struct
func (m *Message) RecvLink() *RecvLink {
	return &RecvLink{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Title:        m.Title,
		Description:  m.Description,
		Url:          m.Url,
		MsgId:        m.MsgId,
	}
}
