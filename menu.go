package weixin

import (
	"encoding/json"
	"errors"
	"fmt"
)

// 自定义菜单
const (
	MenuCreateURL = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
)

// Menu 菜单
type Menu struct {
	Button []Button `json:"button"`
}

// Button 菜单上的按钮
type Button struct {
	Name      string   `json:"name"`
	Type      string   `json:"type,omitempty"`
	Key       string   `json:"key,omitempty"`
	SubButton []Button `json:"sub_button,omitempty"`
}

// CreateMenu 创建菜单
func CreateMenu(menu *Menu) (err error) {
	if len(menu.Button) > 3 {
		return errors.New("too many first level menu, must less than 3")
	}
	for _, sub := range menu.Button {
		if len(sub.SubButton) > 5 {
			return errors.New("too many second level menu, must less than 5")
		}
	}

	jsonBytes, err := json.Marshal(menu)
	if err != nil {
		return err
	}
	url := fmt.Sprintf(MenuCreateURL, AccessToken())
	err = Post(url, jsonBytes)
	return err
}
