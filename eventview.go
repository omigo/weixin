package weixin

// 微信支持的事件类型
const (
	EventTypeScancodePush    EventType = "scancode_push"      // 扫码推事件的事件推送
	EventTypeScancodeWaitmsg EventType = "scancode_waitmsg"   // 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventTypePicSysphoto     EventType = "pic_sysphoto"       // 弹出系统拍照发图的事件推送
	EventTypePicPhotoOrAlbum EventType = "pic_photo_or_album" // 弹出拍照或者相册发图的事件推送
	EventTypePicWeixin       EventType = "pic_weixin"         // 弹出微信相册发图器的事件推送
	EventTypeLocationSelect  EventType = "location_select"    // 弹出地理位置选择器的事件推送
)

// EventScancodePush 扫码推事件的事件推送
type EventScancodePush struct {
	EventBase         // 事件类型 scancode_push
	EventKey   string // 事件KEY值，由开发者在创建菜单时设定
	ScanType   string // 扫描类型，一般是qrcode
	ScanResult string // 扫描结果，即二维码对应的字符串信息
}

// NewEventScancodePush 把通用 struct 转化成相应类型的 struct
func NewEventScancodePush(m *Message) *EventScancodePush {
	return &EventScancodePush{
		EventBase:  EventBase{nil, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Event},
		ScanType:   m.ScanType,
		ScanResult: m.ScanResult,
	}
}

// EventScancodeWaitmsg 扫码推事件且弹出“消息接收中”提示框的事件推送
type EventScancodeWaitmsg struct {
	EventBase         // 事件类型 scancode_waitmsg
	EventKey   string // 事件KEY值，由开发者在创建菜单时设定
	ScanType   string // 扫描类型，一般是qrcode
	ScanResult string // 扫描结果，即二维码对应的字符串信息
}

// NewEventScancodeWaitmsg 把通用 struct 转化成相应类型的 struct
func NewEventScancodeWaitmsg(m *Message) *EventScancodeWaitmsg {
	return &EventScancodeWaitmsg{
		EventBase:  EventBase{nil, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Event},
		ScanType:   m.ScanType,
		ScanResult: m.ScanResult,
	}
}

// EventPicSysphoto 弹出系统拍照发图的事件推送
type EventPicSysphoto struct {
	EventBase           // 事件类型 pic_sysphoto
	EventKey   string   // 事件KEY值，由开发者在创建菜单时设定
	Count      int      // 发送的图片数量
	PicMd5Sums []string // 图片的MD5值，开发者若需要，可用于验证接收到图片
}

// NewEventPicSysphoto 把通用 struct 转化成相应类型的 struct
func NewEventPicSysphoto(m *Message) *EventPicSysphoto {
	return &EventPicSysphoto{
		EventBase:  EventBase{nil, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Event},
		Count:      m.Count,
		PicMd5Sums: m.PicMd5Sums,
	}
}

// EventPicPhotoOrAlbum 弹出拍照或者相册发图的事件推送
type EventPicPhotoOrAlbum struct {
	EventBase           // 事件类型 pic_sysphoto
	EventKey   string   // 事件KEY值，由开发者在创建菜单时设定
	Count      int      // 发送的图片数量
	PicMd5Sums []string // 图片的MD5值，开发者若需要，可用于验证接收到图片
}

// NewEventPicPhotoOrAlbum 把通用 struct 转化成相应类型的 struct
func NewEventPicPhotoOrAlbum(m *Message) *EventPicPhotoOrAlbum {
	return &EventPicPhotoOrAlbum{
		EventBase:  EventBase{nil, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Event},
		Count:      m.Count,
		PicMd5Sums: m.PicMd5Sums,
	}
}

// EventPicWeixin 弹出微信相册发图器的事件推送
type EventPicWeixin struct {
	EventBase           // 事件类型 pic_weixin
	EventKey   string   // 事件KEY值，由开发者在创建菜单时设定
	Count      int      // 发送的图片数量
	PicMd5Sums []string // 图片的MD5值，开发者若需要，可用于验证接收到图片
}

// NewEventPicWeixin 把通用 struct 转化成相应类型的 struct
func NewEventPicWeixin(m *Message) *EventPicWeixin {
	return &EventPicWeixin{
		EventBase:  EventBase{nil, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Event},
		Count:      m.Count,
		PicMd5Sums: m.PicMd5Sums,
	}
}

// EventLocationSelect 弹出地理位置选择器的事件推送
type EventLocationSelect struct {
	EventBase         // 事件类型 pic_sysphoto
	EventKey  string  // 事件KEY值，由开发者在创建菜单时设定
	LocationX float64 // 地理位置维度
	LocationY float64 // 地理位置经度
	Scale     int     // 精度，可理解为精度或者比例尺、越精细的话 scale越高
	Label     string  // 地理位置的字符串信息
	Poiname   string  // 朋友圈POI的名字，可能为空
}

// NewEventLocationSelect 把通用 struct 转化成相应类型的 struct
func NewEventLocationSelect(m *Message) *EventLocationSelect {
	return &EventLocationSelect{
		EventBase: EventBase{nil, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Event},
		EventKey:  m.EventKey,
		LocationX: m.LocationX2,
		LocationY: m.LocationY2,
		Scale:     m.Scale2,
		Label:     m.Label2,
		Poiname:   m.Poiname,
	}
}
