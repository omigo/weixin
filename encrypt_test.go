package weixin

import (
	"encoding/base64"
	"testing"

	"github.com/arstd/log"
)

func TestEncryptMsg(t *testing.T) {
	log.SetLevel(log.Ltrace)

	appId := "wxfabf18ec7ccd2d1a"
	aesKey, _ := base64.StdEncoding.DecodeString("0t37dWsIYg6NsVLgEY1fNuB1rSLyyeQEHOAlIfMhQUV=")

	msg := `<xml><ToUserName><![CDATA[gh_274da2028f77]]></ToUserName>
<FromUserName><![CDATA[ozmLcjnM7vnrXmb3DimFLi0EOiY8]]></FromUserName>
<CreateTime>1448604897</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[sts]]></Content>
<MsgId>6221710657841833060</MsgId>
</xml>`

	// AES CBC 加密报文
	b64Enc, err := EncryptMsg([]byte(msg), aesKey, appId)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%s", b64Enc)
}

func TestDecryptMsg(t *testing.T) {
	log.SetLevel(log.Ltrace)

	appId := "wxfabf18ec7ccd2d1a"
	aesKey, _ := base64.StdEncoding.DecodeString("0t37dWsIYg6NsVLgEY1fNuB1rSLyyeQEHOAlIfMhQUV=")

	b64Enc := "Z8JufHXESFt4chL0Q6vusyowhizt4mpo9Zn3DkyomP7vVhFKi3ICTa1yCOs2XjSl1BaDkKUWl0lQf7psDRwJtP+YD/I6l+TCw0DrRQQyOY9Lf/4FKQ9cpBN+TyhZErDtDJN2E6Euw8VjtV0FmSqH3dGj4sPmWmEiRLldM0luY1WjW1tKGGB2x5vWwFC4piADCw5v9uPYvRk3gZCeknPHmCkCg8ERhi89J7yUuALHwheCo38+4WdQ+YCVVoj7vzZypRiytdwWxvga8OmOk3H99WJdcKQxO7UsgKtpdV/m4rhl3S+iA0HvSOXgQd3v+lAvS8eXsejFUQj92hUP+tV1wKxdg0jK1vxT1Mww0O77N5hIA38atfMMSo8IjVV+HleLbFZ3ByCiyNxrrGDh8ljqFNyVwcJJcz9ZZAnu3XOf+BQ="

	// AES CBC 解密报文
	src, err := DecryptMsg(b64Enc, aesKey, appId)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%s", src)
}
