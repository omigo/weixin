package weixin

import "time"

// 微信公众平台测试号
// http://mp.weixin.qq.com/debug/cgi-bin/sandboxinfo?action=showinfo&t=sandbox/index
// 测试号 毛毛雨
// const (
// 	originId       = "gh_398ea213a0d3"
// 	appId          = "wxad7f6407fcbb9cae"
// 	appSecret      = "0884d88aaa8ef70dfcd484fcc52bd484"
// 	token          = "0t37dWsIYg6NsVLgEY1fNuB1rSLyyeQE"
// 	encodingAESKey = "dkVSmV3CtG3IBCY96A8RSNoCOzaPcx36uGJo8fl9wWn" // 为了通过验证
// )

// ******
const (
	originId       = "gh_××××××××××"
	appId          = "wx****************"
	appSecret      = "*******************************"
	token          = "*******************************"
	encodingAESKey = "*******************************************"
)

// 单元测试 testURL
const testURL = "http://127.0.0.1:3080/weixin?signature=46ff81d2fad29f279b2c83eb6c1f1ea352eb9a16&timestamp=1449049868&nonce=496941103"

func init() {
	Initialize(originId, appId, appSecret, token, encodingAESKey)
	time.Sleep(3 * time.Second) // waiting to refresh token
}
