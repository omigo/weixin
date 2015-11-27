package weixin

import (
	"crypto/aes"
	"encoding/base64"
	"testing"
)

func TestAESCBCDecrypt(t *testing.T) {
	aesKey := []byte("0123456789abcdef0123456789abcdef")
	src := "我们都是好孩子"

	b64Enc := "MTp5u8m7i4zMqJFlSo1QBIUn+iASUNmd+Co9u0y4Y5w="
	enc, _ := base64.StdEncoding.DecodeString(b64Enc)

	actual, err := AESCBCDecrypt(enc, aesKey, aesKey[:aes.BlockSize])
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if string(actual) != src {
		t.Logf("expect %s, but get %s", src, actual)
		t.FailNow()
	}
}

func TestAESCBCEncryptAndDecrypt(t *testing.T) {
	aesKey := []byte("0123456789abcdef0123456789abcdef")
	src := "我们都是好孩子"

	enc, err := AESCBCEncrypt([]byte(src), aesKey, aesKey[:aes.BlockSize])
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%s", base64.StdEncoding.EncodeToString(enc))

	actual, err := AESCBCDecrypt(enc, aesKey, aesKey[:aes.BlockSize])
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if string(actual) != src {
		t.Logf("expect %s, but get %s", src, actual)
		t.FailNow()
	}
}
