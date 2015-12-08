package weixin

import "testing"

func TestGenRedirectURL(t *testing.T) {
	// t.Log(GenRedirectURL("https://omigo.tunnel.phpor.me/index.html", ScopeSNSAPIBase, "test"))
	t.Log(GenRedirectURL("https://api.shou.money/404", ScopeSNSAPIBase, "test"))
}

func TestGetWebToken(t *testing.T) {
	code := "0018405ea79f591fd493fb394bdf0beF"
	token, err := GetWebToken(code)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%#v", token)
}

func TestRefreshWebToken(t *testing.T) {
	refreshToken := "OezXcEiiBSKSxW0eoylIeGrNYdd9Xv-i9tXZN7AnXUhD37adoE4xCoIa3iAYvWg6jpA8yXNAiWTRCuocJ-XKZgLsIkouhlCpndd43Lhr08k3w9SxLtA8MWvaM08fpGR8UDd53a68it5F88lrtMRqCw"
	token, err := RefreshWebToken(refreshToken)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%#v", token)
}

func TestGetWebUserInfo(t *testing.T) {
	code := "02193e8fd89b2e38db807aa4de22b3bK"
	webToken, err := GetWebToken(code)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%#v", webToken)

	info, err := GetWebUserInfo(webToken.AccessToken, webToken.OpenId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%#v", info)
}

func TestCheckWebToken(t *testing.T) {
	token := "OezXcEiiBSKSxW0eoylIeGrNYdd9Xv-i9tXZN7AnXUhD37adoE4xCoIa3iAYvWg6y57zk_rcZQ87v7JyYCaJl424nN6t_dgessfE_zaU1NS-GT2U5WNHy5uezUhUaLbSt5eQpm1wSnZC9VDkuWfTWQ"
	openId := "ozmLcjnM7vnrXmb3DimFLi0EOiY8"
	err := CheckWebToken(token, openId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
