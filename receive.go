package weixin

// MsgText 接收文本消息
type MsgText struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间（整型）
	MsgType      MsgType // text
	Content      string  // 文本消息内容
	MsgId        int     // 消息id，64位整型
}

// MsgText 把通用 struct 转化成相应类型的 struct
func (m *Message) MsgText() *MsgText {
	return &MsgText{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		Content:      m.Content,
		MsgId:        m.MsgId,
	}
}

// MsgImage 接收图片消息
type MsgImage struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间（整型）
	MsgType      MsgType // image
	PicUrl       string  // 图片链接
	MediaId      string  // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据
	MsgId        int     // 消息id，64位整型
}

// MsgImage 把通用 struct 转化成相应类型的 struct
func (m *Message) MsgImage() *MsgImage {
	return &MsgImage{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		PicUrl:       m.PicUrl,
		MediaId:      m.MediaId,
		MsgId:        m.MsgId,
	}
}

// MsgVoice 接收视频/小视频消息
type MsgVoice struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间（整型）
	MsgType      MsgType // voice
	MediaId      string  // 语音消息媒体id，可以调用多媒体文件下载接口拉取数据
	Format       string  // 语音格式，如amr，speex等
	Recongnition string  // 语音识别结果，使用UTF8编码
	MsgId        int     // 消息id，64位整型
}

// MsgVoice 把通用 struct 转化成相应类型的 struct
func (m *Message) MsgVoice() *MsgVoice {
	return &MsgVoice{
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

// MsgVideo 接收图片消息
type MsgVideo struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间（整型）
	MsgType      MsgType // video
	MediaId      string  // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据
	ThumbMediaId string  // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据
	MsgId        int     // 消息id，64位整型
}

// MsgVideo 把通用 struct 转化成相应类型的 struct
func (m *Message) MsgVideo() *MsgVideo {
	return &MsgVideo{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		MediaId:      m.MediaId,
		ThumbMediaId: m.ThumbMediaId,
		MsgId:        m.MsgId,
	}
}

// MsgLocation 接收地理位置消息
type MsgLocation struct {
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

// MsgLocation 把通用 struct 转化成相应类型的 struct
func (m *Message) MsgLocation() *MsgLocation {
	return &MsgLocation{
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

// MsgLink 接收链接消息
type MsgLink struct {
	ToUserName   string  // 开发者微信号
	FromUserName string  // 发送方帐号（一个OpenID）
	CreateTime   string  // 消息创建时间（整型）
	MsgType      MsgType // location
	Title        string  // 消息标题
	Description  string  // 消息描述
	Url          string  // 消息链接
	MsgId        int     // 消息id，64位整型
}

// MsgLink 把通用 struct 转化成相应类型的 struct
func (m *Message) MsgLink() *MsgLink {
	return &MsgLink{
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
