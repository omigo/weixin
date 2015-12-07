package weixin

import "testing"

func TestCreateUserGroup(t *testing.T) {
	g, err := CreateUserGroup("test")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v", g)
}

func TestUpdateUserGroup(t *testing.T) {
	g, err := UpdateUserGroup(100, "test2")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v", g)
}

func TestGetAllUserGroups(t *testing.T) {
	gs, err := GetAllUserGroups()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v", gs)
}

func TestGetGroupIdByOpenId(t *testing.T) {
	g, err := GetGroupIdByOpenId("ozmLcjnM7vnrXmb3DimFLi0EOiY8")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v", g)
}

func TestUpdateMemberGroup(t *testing.T) {
	err := UpdateMemberGroup("ozmLcjnM7vnrXmb3DimFLi0EOiY8", 100)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestBatchUpdateMemberGroup(t *testing.T) {
	err := BatchUpdateMemberGroup([]string{"ozmLcjnM7vnrXmb3DimFLi0EOiY8"}, 100)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestDeleteUserGroup(t *testing.T) {
	err := DeleteUserGroup(101)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
