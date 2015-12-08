package weixin

import "testing"

func TestUpdateUserRemark(t *testing.T) {
	err := UpdateUserRemark("ozmLcjnM7vnrXmb3DimFLi0EOiY8", "毛毛雨")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestGetUserInfo(t *testing.T) {
	info, err := GetUserInfo("ozmLcjnM7vnrXmb3DimFLi0EOiY8", LangZHTW)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v", info)
}

func TestBatchGetUserInfo(t *testing.T) {
	infos, err := BatchGetUserInfo([]string{"ozmLcjnM7vnrXmb3DimFLi0EOiY8",
		"ozmLcjnM7vnrXmb3DimFLi0EOiY8"}, LangZHTW)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v", infos)
}

func TestGetUserList(t *testing.T) {
	openIds, total, count, nextOpenId, err := GetUserList()

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v %#v %#v %#v", openIds, total, count, nextOpenId)
}
