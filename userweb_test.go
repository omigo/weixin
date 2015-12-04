package weixin

import "testing"

func TestGenRedirectURL(t *testing.T) {
	t.Log(GenRedirectURL("https://omigo.tunnel.phpor.me/index.html", ScopeSNSAPIBase, "test"))
}

func TestGetJsToken(t *testing.T) {
	code := "0318985eaa8cf3f0d6f7deba44407daA"
	token, err := GetJsToken(code)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%#v", token)
}

func TestRefreshToken(t *testing.T) {
	refreshToken := "OezXcEiiBSKSxW0eoylIeGrNYdd9Xv-i9tXZN7AnXUhD37adoE4xCoIa3iAYvWg6jpA8yXNAiWTRCuocJ-XKZgLsIkouhlCpndd43Lhr08k3w9SxLtA8MWvaM08fpGR8UDd53a68it5F88lrtMRqCw"
	token, err := RefreshToken(refreshToken)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%#v", token)
}
