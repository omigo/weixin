package weixin

import (
	"errors"
	"fmt"
)

// 个性化菜单
const (
	MenuCreateConditionalURL       = "https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=%s"
	MenuTryMatchConditionalMenuURL = "https://api.weixin.qq.com/cgi-bin/menu/trymatch?access_token=%s"
	MenuDeleteConditionalURL       = "https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token=%s"
)

// ConditionalMenu  个性化菜单
type ConditionalMenu struct {
	Button    []Button  `json:"button"`    // 一级菜单数组，个数应为1~3个
	MatchRule MatchRule `json:"matchrule"` // 菜单匹配规则
}

// MatchRule 菜单匹配规则，六个字段，均可为空，但不能全部为空，至少要有一个匹配信息是不为空的，
type MatchRule struct {
	GroupId int `json:"group_id,omitempty"` // 用户分组id，可通过用户分组管理接口获取
	Sex     int `json:"sex,omitempty"`      // 性别：男（1）女（2），不填则不做匹配
	// 客户端版本，当前只具体到系统型号：IOS(1), Android(2),Others(3)，不填则不做匹配
	ClientPlatformType int `json:"client_platform_type,omitempty"`
	// country、province、city组成地区信息，将按照country、province、city的顺序进行验证，
	// 要符合地区信息表的内容。地区信息从大到小验证，小的可以不填，即若填写了省份信息，则国家信
	// 息也必填并且匹配，城市信息可以不填。 例如 “中国 广东省 广州市”、“中国 广东省”都是合法
	// 的地域信息，而“中国 广州市”则不合法，因为填写了城市信息但没有填写省份信息。
	// 地区信息表：http://mp.weixin.qq.com/wiki/static/assets/870a3c2a14e97b3e74fde5e88fa47717.zip
	Country  string `json:"country,omitempty"`  // 国家信息，是用户在微信中设置的地区，具体请参考地区信息表
	Province string `json:"province,omitempty"` // 省份信息，是用户在微信中设置的地区，具体请参考地区信息表
	City     string `json:"city,omitempty"`     // 城市信息，是用户在微信中设置的地区，具体请参考地区信息表
}

// CreateConditionalMenu 创建个性化菜单
func CreateConditionalMenu(cm *ConditionalMenu) (menuId string, err error) {
	if len(cm.Button) > 3 {
		return "", errors.New("too many first level menu, must less than 3")
	}
	for _, sub := range cm.Button {
		if len(sub.SubButton) > 5 {
			return "", errors.New("too many second level menu, must less than 5")
		}
	}

	url := fmt.Sprintf(MenuCreateConditionalURL, AccessToken())

	wapper := &struct {
		MenuId string `json:"menuid"`
	}{}
	err = PostMarshalUnmarshal(url, cm, wapper)
	return wapper.MenuId, err
}

// DeleteConditionalMenu 删除个性化菜单，menuId 为菜单 id，可以通过自定义菜单查询接口获取
func DeleteConditionalMenu(menuId int) (err error) {
	url := fmt.Sprintf(MenuDeleteConditionalURL, AccessToken())
	js := fmt.Sprintf(`{"menuid":"%d"}`, menuId)
	err = Post(url, []byte(js))
	return err
}

// TryMatchConditionalMenu 测试个性化菜单匹配结果，userId 可以是粉丝的 OpenID，也可以是粉丝的微信号
func TryMatchConditionalMenu(userId string) (buttons []Button, err error) {
	url := fmt.Sprintf(MenuTryMatchConditionalMenuURL, AccessToken())
	js := fmt.Sprintf(`{"user_id":"%s"}`, userId)

	wapper := &struct {
		Button []Button `json:"button"`
	}{}
	err = PostUnmarshal(url, []byte(js), wapper)
	return wapper.Button, err
}
