package weixin

import "testing"

func TestCreateMenu(t *testing.T) {
	Initialize(originId, appId, appSecret, token, encodingAESKey)

	buttons := []Button{
		Button{
			Name: "扫码",
			SubButton: []Button{
				Button{
					Name: "扫码带提示",
					Type: MenuTypeScancodeWaitmsg,
					Key:  "rselfmenu_0_0",
				},
				Button{
					Name: "扫码推事件",
					Type: MenuTypeScancodePush,
					Key:  "rselfmenu_0_1",
				},
			},
		},
		Button{
			Name: "发图",
			SubButton: []Button{
				Button{
					Name: "系统拍照发图",
					Type: MenuTypePicSysphoto,
					Key:  "rselfmenu_1_0",
				},
				Button{
					Name: "拍照或者相册发图",
					Type: MenuTypePicPhotoOrAlbum,
					Key:  "rselfmenu_1_1",
				},
				Button{
					Name: "微信相册发图",
					Type: MenuTypePicWeixin,
					Key:  "rselfmenu_1_2",
				},
			},
		},
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
	}

	err := CreateMenu(buttons)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestGetMenu(t *testing.T) {
	Initialize(originId, appId, appSecret, token, encodingAESKey)

	buttons, err := GetMenu()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%#v", buttons)
}

func TestGetMenuInfo(t *testing.T) {
	Initialize(originId, appId, appSecret, token, encodingAESKey)

	mi, err := GetMenuInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%v", mi)
}
