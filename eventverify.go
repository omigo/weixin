package weixin

// 微信支持的事件类型
const (
	EventTypeQualificationVerifySuccess EventType = "qualification_verify_success" // 资质认证成功
	EventTypeQualificationVerifyFail    EventType = "qualification_verify_fail"    // 资质认证失败
	EventTypeNamingVerifySuccess        EventType = "naming_verify_success"        // 名称认证成功（即命名成功）
	EventTypeNamingVerifyFail           EventType = "naming_verify_fail"           // 名称认证失败
	EventTypeAnnualRenew                EventType = "annual_renew"                 // 年审通知
	EventTypeVerifyExpired              EventType = "verify_expired"               // 认证过期失效通知
)

// EventQualificationVerifySuccess 资质认证成功（此时立即获得接口权限）
type EventQualificationVerifySuccess struct {
	EventBase       // 事件类型 qualification_verify_success
	ExpiredTime int // 有效期 (整形)，指的是时间戳，将于该时间戳认证过期
}

// NewEventQualificationVerifySuccess 把通用 struct 转化成相应类型的 struct
func NewEventQualificationVerifySuccess(m *Message) *EventQualificationVerifySuccess {
	return &EventQualificationVerifySuccess{
		EventBase:   EventBase{nil, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Event},
		ExpiredTime: m.ExpiredTime,
	}
}

// EventQualificationVerifyFail 资质认证失败
type EventQualificationVerifyFail struct {
	EventBase         // 事件类型 qualification_verify_fail
	FailTime   int    // 失败发生时间 (整形)，时间戳
	FailReason string // 认证失败的原因
}

// NewEventQualificationVerifyFail 把通用 struct 转化成相应类型的 struct
func NewEventQualificationVerifyFail(m *Message) *EventQualificationVerifyFail {
	return &EventQualificationVerifyFail{
		EventBase:  EventBase{nil, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Event},
		FailTime:   m.FailTime,
		FailReason: m.FailReason,
	}
}

// EventNamingVerifySuccess 名称认证成功（即命名成功）
type EventNamingVerifySuccess struct {
	EventBase       // 事件类型 naming_verify_success
	ExpiredTime int // 有效期 (整形)，指的是时间戳，将于该时间戳认证过期
}

// NewEventNamingVerifySuccess 把通用 struct 转化成相应类型的 struct
func NewEventNamingVerifySuccess(m *Message) *EventNamingVerifySuccess {
	return &EventNamingVerifySuccess{
		EventBase:   EventBase{nil, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Event},
		ExpiredTime: m.ExpiredTime,
	}
}

// EventNamingVerifyFail 名称认证失败（这时虽然客户端不打勾，但仍有接口权限）
type EventNamingVerifyFail struct {
	EventBase         // 事件类型 naming_verify_fail
	FailTime   int    // 失败发生时间 (整形)，时间戳
	FailReason string // 认证失败的原因
}

// NewEventNamingVerifyFail 把通用 struct 转化成相应类型的 struct
func NewEventNamingVerifyFail(m *Message) *EventNamingVerifyFail {
	return &EventNamingVerifyFail{
		EventBase:  EventBase{nil, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Event},
		FailTime:   m.FailTime,
		FailReason: m.FailReason,
	}
}

// EventAnnualRenew 年审通知
type EventAnnualRenew struct {
	EventBase       // 事件类型 annual_renew
	ExpiredTime int // 有效期 (整形)，指的是时间戳，将于该时间戳认证过期，需尽快年审
}

// NewEventAnnualRenew 把通用 struct 转化成相应类型的 struct
func NewEventAnnualRenew(m *Message) *EventAnnualRenew {
	return &EventAnnualRenew{
		EventBase:   EventBase{nil, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Event},
		ExpiredTime: m.ExpiredTime,
	}
}

// EventVerifyExpired 认证过期失效通知
type EventVerifyExpired struct {
	EventBase       // 事件类型 verify_expired
	ExpiredTime int // 有效期 (整形)，指的是时间戳，表示已于该时间戳认证过期，需要重新发起微信认证
}

// NewEventVerifyExpired 把通用 struct 转化成相应类型的 struct
func NewEventVerifyExpired(m *Message) *EventVerifyExpired {
	return &EventVerifyExpired{
		EventBase:   EventBase{nil, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Event},
		ExpiredTime: m.ExpiredTime,
	}
}
