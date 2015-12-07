package weixin

import (
	"encoding/json"
	"fmt"
)

// 用户分组管理
const (
	UserGroupCreateURL                 = "https://api.weixin.qq.com/cgi-bin/groups/create?access_token=%s"
	UserGroupUpdateURL                 = "https://api.weixin.qq.com/cgi-bin/groups/update?access_token=%s"
	UserGroupDeleteURL                 = "https://api.weixin.qq.com/cgi-bin/groups/delete?access_token=%s"
	UserGroupGetAllURL                 = "https://api.weixin.qq.com/cgi-bin/groups/get?access_token=%s"
	UserGroupGetGroupIdURL             = "https://api.weixin.qq.com/cgi-bin/groups/getid?access_token=%s"
	UserGroupUpdateMemberGroupURL      = "https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token=%s"
	UserGroupBatchUpdateMemberGroupURL = "https://api.weixin.qq.com/cgi-bin/groups/members/batchupdate?access_token=%s"
)

// GroupsWapper 所有用户分组包装器
type GroupsWapper struct {
	WeixinError
	Groups []Group `json:"groups"`
}

// GroupWapper 用户分组包装器
type GroupWapper struct {
	WeixinError
	Group Group `json:"group"`
}

// Group 用户分组
type Group struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Count int    `json:"count,omitempty"`
}

// CreateUserGroup 创建用户分组
func CreateUserGroup(name string) (g *Group, err error) {
	url := fmt.Sprintf(UserGroupCreateURL, AccessToken())

	w := &GroupWapper{
		Group: Group{
			Name: name,
		},
	}

	err = PostMarshalUnmarshal(url, w, w)
	if err != nil {
		return nil, err
	}
	return &w.Group, nil
}

// UpdateUserGroup 修改用户分组名
func UpdateUserGroup(id int, name string) (g *Group, err error) {
	url := fmt.Sprintf(UserGroupCreateURL, AccessToken())

	w := &GroupWapper{
		Group: Group{
			Id:   id,
			Name: name,
		},
	}

	err = PostMarshalUnmarshal(url, w, w)
	if err != nil {
		return nil, err
	}
	return &w.Group, nil
}

// DeleteUserGroup 修改用户分组名
func DeleteUserGroup(groupId int) (err error) {
	url := fmt.Sprintf(UserGroupDeleteURL, AccessToken())

	body := fmt.Sprintf(`{"group":{"id":%d}}`, groupId)
	return Post(url, []byte(body))
}

// GetAllUserGroups 查询所有分组
func GetAllUserGroups() (gs []Group, err error) {
	url := fmt.Sprintf(UserGroupGetAllURL, AccessToken())

	w := &GroupsWapper{}
	err = GetUnmarshal(url, w)
	if err != nil {
		return nil, err
	}
	return w.Groups, nil
}

// GroupIdWapper 用户所在分组包装器
type GroupIdWapper struct {
	GroupId int `json:"groupid"`
}

// GetGroupIdByOpenId 查询用户所在分组
func GetGroupIdByOpenId(openId string) (groupId int, err error) {
	url := fmt.Sprintf(UserGroupGetGroupIdURL, AccessToken())

	body := fmt.Sprintf(`{"openid":"%s"}`, openId)
	wapper := &GroupIdWapper{}
	err = PostUnmarshal(url, []byte(body), wapper)
	if err != nil {
		return 0, err
	}
	return wapper.GroupId, nil
}

// UpdateMemberGroup 移动用户分组
func UpdateMemberGroup(openId string, toGroupId int) (err error) {
	url := fmt.Sprintf(UserGroupUpdateMemberGroupURL, AccessToken())
	body := fmt.Sprintf(`{"openid":"%s","to_groupid":%d}`, openId, toGroupId)
	return Post(url, []byte(body))
}

// BatchUpdateMemberGroup 批量移动用户分组
func BatchUpdateMemberGroup(openIds []string, toGroupId int) (err error) {
	url := fmt.Sprintf(UserGroupBatchUpdateMemberGroupURL, AccessToken())

	js, _ := json.Marshal(openIds)
	body := fmt.Sprintf(`{"openid_list":%s,"to_groupid":%d}`, js, toGroupId)
	return Post(url, []byte(body))
}
