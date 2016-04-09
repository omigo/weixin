package weixin

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/gotips/log"
)

// AESCBCEncrypt 采用 CBC 模式的 AES 加密
func AESCBCEncrypt(src, key, iv []byte) (enc []byte, err error) {
	log.Tracef("src: %s", src)
	src = PKCS7Padding(src, len(key))

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(src, src)
	enc = src

	log.Tracef("enc: % x", enc)
	return enc, nil
}

// AESCBCDecrypt 采用 CBC 模式的 AES 解密
func AESCBCDecrypt(enc, key, iv []byte) (src []byte, err error) {
	log.Tracef("enc: % x", enc)
	if len(enc) < len(key) {
		return nil, fmt.Errorf("the length of encrypted message too short: %d", len(enc))
	}
	if len(enc)&(len(key)-1) != 0 { // or len(enc)%len(key) != 0
		return nil, fmt.Errorf("encrypted message is not a multiple of the key size(%d), the length is %d", len(key), len(enc))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(enc, enc)
	src = PKCS7UnPadding(enc)

	log.Tracef("src: %s", src)
	return src, nil
}

// PKCS7Padding PKCS#7填充，Buf需要被填充为K的整数倍，
// 在buf的尾部填充(K-N%K)个字节，每个字节的内容是(K- N%K)
func PKCS7Padding(src []byte, k int) (padded []byte) {
	padLen := k - len(src)%k
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(src, padding...)
}

// PKCS7UnPadding 去掉PKCS#7填充，Buf需要被填充为K的整数倍，
// 在buf的尾部填充(K-N%K)个字节，每个字节的内容是(K- N%K)
func PKCS7UnPadding(src []byte) (padded []byte) {
	padLen := int(src[len(src)-1])
	return src[:len(src)-padLen]
}
