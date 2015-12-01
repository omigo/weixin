package weixin

import "testing"

func TestCreateMenu(t *testing.T) {
	Initialize(originId, appId, appSecret, token, encodingAESKey)

	menu := &Menu{
		Button: []Button{
			Button{
				Name: "扫码",
				SubButton: []Button{
					Button{
						Name: "扫码带提示",
						Type: "scancode_waitmsg",
						Key:  "rselfmenu_0_0",
					},
					Button{
						Name: "扫码推事件",
						Type: "scancode_push",
						Key:  "rselfmenu_0_1",
					},
				},
			},
			Button{
				Name: "发图",
				SubButton: []Button{
					Button{
						Name: "系统拍照发图",
						Type: "pic_sysphoto",
						Key:  "rselfmenu_1_0",
					},
					Button{
						Name: "拍照或者相册发图",
						Type: "pic_photo_or_album",
						Key:  "rselfmenu_1_1",
					},
					Button{
						Name: "微信相册发图",
						Type: "pic_weixin",
						Key:  "rselfmenu_1_2",
					},
				},
			},
			Button{
				Name: "发送位置",
				Type: "location_select",
				Key:  "rselfmenu_2_0",
			},
		},
	}

	err := CreateMenu(menu)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
