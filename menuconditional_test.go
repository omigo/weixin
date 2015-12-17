package weixin

import "testing"

func TestCreateConditionalMenu(t *testing.T) {
	cm := &ConditionalMenu{
		Button: []Button{
			Button{
				Name: "测试",
				SubButton: []Button{
					Button{
						Name: "腾讯",
						Type: MenuTypeView,
						URL:  "http://qq.com",
					},
					Button{
						Name: "发送位置",
						Type: MenuTypeLocationSelect,
						Key:  "rselfmenu_2_0",
					},
				},
			},
		},
		MatchRule: MatchRule{
			GroupId: 102, // 用户分组id，可通过用户分组管理接口获取
			// Sex:     1,   // 性别：男（1）女（2），不填则不做匹配
			// // 客户端版本，当前只具体到系统型号：IOS(1), Android(2),Others(3)，不填则不做匹配
			// ClientPlatformType: 1,
			// // country、province、city组成地区信息，将按照country、province、city的顺序进行验证，
			// // 要符合地区信息表的内容。地区信息从大到小验证，小的可以不填，即若填写了省份信息，则国家信
			// // 息也必填并且匹配，城市信息可以不填。 例如 “中国 广东省 广州市”、“中国 广东省”都是合法
			// // 的地域信息，而“中国 广州市”则不合法，因为填写了城市信息但没有填写省份信息。
			// // 地区信息表：http://mp.weixin.qq.com/wiki/static/assets/870a3c2a14e97b3e74fde5e88fa47717.zip
			// Country:  "中国",  // 国家信息，是用户在微信中设置的地区，具体请参考地区信息表
			// Province: "广东省", // 省份信息，是用户在微信中设置的地区，具体请参考地区信息表
			// City:     "广州市", // 城市信息，是用户在微信中设置的地区，具体请参考地区信息表
		},
	}

	menuId, err := CreateConditionalMenu(cm)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v", menuId)
}

func TestDeleteConditionalMenu(t *testing.T) {
	menuId := 401859640
	err := DeleteConditionalMenu(menuId)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestTryMatchConditionalMenu(t *testing.T) {
	buttons, err := TryMatchConditionalMenu("ozmLcjnM7vnrXmb3DimFLi0EOiY8")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v", buttons)
}
