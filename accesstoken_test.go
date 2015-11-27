package weixin

import (
	"testing"
	"time"
)

func TestAccessToken(t *testing.T) {
	appId := "**********"
	appSecret := "**********"
	RefreshAccessToken(appId, appSecret)

	time.Sleep(3 * time.Second)
	t.Log(AccessToken())
}
