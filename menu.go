package weixin

import (
	"errors"
	"fmt"
)

// 自定义菜单
const (
	MenuCreateURL = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
	MenuGetURL    = "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s"
	MenuDeleteURL = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s"
	MenuInfoURL   = "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=%s"
)

/*
MenuType 菜单类型

请注意：
	1、自定义菜单最多包括3个一级菜单，每个一级菜单最多包含5个二级菜单。
	2、一级菜单最多4个汉字，二级菜单最多7个汉字，多出来的部分将会以“...”代替。
	3、创建自定义菜单后，由于微信客户端缓存，需要24小时微信客户端才会展现出来。测试时可以尝试取消关注公众账号后再次关注，则可以看到创建后的效果。
	自定义菜单接口可实现多种类型按钮，如下：

1、click：点击推事件
	用户点击click类型按钮后，微信服务器会通过消息接口推送消息类型为event	的结构给开发者（参考消息接口指南），并且带上按钮中开发者填写的key值，开发者可以通过自定义的key值与用户进行交互；
2、view：跳转URL
	用户点击view类型按钮后，微信客户端将会打开开发者在按钮中填写的网页URL，可与网页授权获取用户基本信息接口结合，获得用户基本信息。
3、scancode_push：扫码推事件
	用户点击按钮后，微信客户端将调起扫一扫工具，完成扫码操作后显示扫描结果（如果是URL，将进入URL），且会将扫码的结果传给开发者，开发者可以下发消息。
4、scancode_waitmsg：扫码推事件且弹出“消息接收中”提示框
	用户点击按钮后，微信客户端将调起扫一扫工具，完成扫码操作后，将扫码的结果传给开发者，同时收起扫一扫工具，然后弹出“消息接收中”提示框，随后可能会收到开发者下发的消息。
5、pic_sysphoto：弹出系统拍照发图
	用户点击按钮后，微信客户端将调起系统相机，完成拍照操作后，会将拍摄的相片发送给开发者，并推送事件给开发者，同时收起系统相机，随后可能会收到开发者下发的消息。
6、pic_photo_or_album：弹出拍照或者相册发图
	用户点击按钮后，微信客户端将弹出选择器供用户选择“拍照”或者“从手机相册选择”。用户选择后即走其他两种流程。
7、pic_weixin：弹出微信相册发图器
	用户点击按钮后，微信客户端将调起微信相册，完成选择操作后，将选择的相片发送给开发者的服务器，并推送事件给开发者，同时收起相册，随后可能会收到开发者下发的消息。
8、location_select：弹出地理位置选择器
	用户点击按钮后，微信客户端将调起地理位置选择工具，完成选择操作后，将选择的地理位置发送给开发者的服务器，同时收起位置选择工具，随后可能会收到开发者下发的消息。
9、media_id：下发消息（除文本消息）
	用户点击media_id类型按钮后，微信服务器会将开发者填写的永久素材id对应的素材下发给用户，永久素材类型可以是图片、音频、视频、图文消息。请注意：永久素材id必须是在“素材管理/新增永久素材”接口上传后获得的合法id。
10、view_limited：跳转图文消息URL
	用户点击view_limited类型按钮后，微信客户端将打开开发者在按钮中填写的永久素材id对应的图文消息URL，永久素材类型只支持图文消息。请注意：永久素材id必须是在“素材管理/新增永久素材”接口上传后获得的合法id。

请注意，3到8的所有事件，仅支持微信iPhone5.4.1以上版本，和Android5.4以上版本的微信用户，旧版本微信用户点击后将没有回应，开发者也不能正常接收到事件推送。9和10，是专门给第三方平台旗下未微信认证（具体而言，是资质认证未通过）的订阅号准备的事件类型，它们是没有事件推送的，能力相对受限，其他类型的公众号不必使用。
*/
type MenuType string

// 各种菜单类型
const (
	MenuTypeClick           = "click"              // 点击推事件
	MenuTypeView            = "view"               // 点击推事件
	MenuTypeScancodePush    = "scancode_push"      // 扫码推事件
	MenuTypeScancodeWaitmsg = "scancode_waitmsg"   // 扫码推事件且弹出“消息接收中”提示框
	MenuTypePicSysphoto     = "pic_sysphoto"       // 弹出系统拍照发图
	MenuTypePicPhotoOrAlbum = "pic_photo_or_album" // 弹出拍照或者相册发图
	MenuTypePicWeixin       = "pic_weixin"         // 弹出微信相册发图器
	MenuTypeLocationSelect  = "location_select"    // 弹出地理位置选择器
	MenuTypeMediaId         = "media_id"           // 下发消息（除文本消息）
	MenuTypeViewLimited     = "view_limited"       // 跳转图文消息URL

)

// AllMenu 自定义菜单
type AllMenu struct {
	WXError
	Menu struct {
		Button []Button `json:"button"`
		MenuId int      `json:"menuid"` // 菜单 id
	} `json:"menu"`
	ConditionalMenu struct {
		Button    []Button  `json:"button"`
		MenuId    int       `json:"menuid"`    // 菜单 id
		MatchRule MatchRule `json:"matchrule"` // 菜单匹配规则
	} `json:"conditionalmenu"`
}

// Button 菜单上的按钮
type Button struct {
	Name      string   `json:"name"`
	Type      MenuType `json:"type,omitempty"`
	Key       string   `json:"key,omitempty"`
	URL       string   `json:"url,omitempty"`
	SubButton []Button `json:"sub_button,omitempty"`
}

// CreateMenu 创建菜单
func CreateMenu(buttons []Button) (err error) {
	if len(buttons) > 3 {
		return errors.New("too many first level menu, must less than 3")
	}
	for _, sub := range buttons {
		if len(sub.SubButton) > 5 {
			return errors.New("too many second level menu, must less than 5")
		}
	}

	menu := struct {
		Button []Button `json:"button"`
	}{buttons}

	url := fmt.Sprintf(MenuCreateURL, AccessToken())
	return Post(url, menu, nil)
}

// GetMenu 查询菜单
func GetMenu() (all *AllMenu, err error) {
	url := fmt.Sprintf(MenuGetURL, AccessToken())
	all = &AllMenu{}
	err = Get(url, all)
	return all, err
}

// DeleteMenu 删除菜单
func DeleteMenu() (err error) {
	url := fmt.Sprintf(MenuDeleteURL, AccessToken())
	return Get(url, nil)
}

// GetMenuInfo 获取自定义菜单配置
func GetMenuInfo() (mi *MenuInfo, err error) {
	url := fmt.Sprintf(MenuInfoURL, AccessToken())
	mi = &MenuInfo{}
	err = Get(url, mi)
	return mi, err
}

// MenuInfo 自定义菜单配置
type MenuInfo struct {
	WXError
	IsMenuOpen   int `json:"is_menu_open"` // 菜单是否开启，0代表未开启，1代表开启
	SelfmenuInfo struct {
		Button []struct {
			// 菜单的类型，公众平台官网上能够设置的菜单类型有view（跳转网页）、text（返回文本，下同）、
			// img、photo、video、voice。使用API设置的则有8种，详见《自定义菜单创建接口》
			Type MenuType `json:"type"`
			Name string   `json:"name"` // 菜单名称
			// 对于不同的菜单类型，value的值意义不同。官网上设置的自定义菜单：
			// Text:保存文字到value； Img、voice：保存mediaID到value； Video：保存视频下载链接到value；
			//  News：保存图文消息到news_info，同时保存mediaID到value； View：保存链接到url。
			// 使用API设置的自定义菜单： click、scancode_push、scancode_waitmsg、pic_sysphoto、
			// pic_photo_or_album、	pic_weixin、location_select：保存值到key；view：保存链接到url
			Key       string `json:"key"`
			URL       string `json:"url"`
			Value     string `json:"value"`
			SubButton struct {
				List []struct {
					Name     string   `json:"name"`
					Type     MenuType `json:"type"`
					Key      string   `json:"key"`
					URL      string   `json:"url"`
					Value    string   `json:"value"`
					NewsInfo struct {
						List []struct {
							Title      string `json:"title"`       // 图文消息的标题
							Author     string `json:"author"`      // 作者
							Digest     string `json:"digest"`      // 摘要
							ShowCover  string `json:"show_cover"`  // 是否显示封面，0为不显示，1为显示
							CoverURL   string `json:"cover_url"`   // 封面图片的URL
							ContentURL string `json:"content_url"` // 正文的URL
							SourceURL  string `json:"source_url"`  // 原文的URL，若置空则无查看原文入口
						}
					} `json:"news_info"` // news_info	图文消息的信息
				} `json:"list"`
			} `json:"sub_button"` // 菜单按钮
		} `json:"button"` // 菜单按钮
	} `json:"selfmenu_info"` // 菜单信息
}
