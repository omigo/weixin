package weixin

import (
	"time"

	"github.com/gotips/log"
)

// 微信公众平台测试号
// http://mp.weixin.qq.com/debug/cgi-bin/sandboxinfo?action=showinfo&t=sandbox/index
// 测试号 毛毛雨
const (
	originId       = "gh_398ea213a0d3"
	appId          = "wxad7f6407fcbb9cae"
	appSecret      = "0884d88aaa8ef70dfcd484fcc52bd484"
	token          = "0t37dWsIYg6NsVLgEY1fNuB1rSLyyeQE"
	encodingAESKey = "dkVSmV3CtG3IBCY96A8RSNoCOzaPcx36uGJo8fl9wWn" // 为了通过验证
)

// 单元测试 testURL
const testURL = "http://127.0.0.1:3080/weixin?signature=292ea5c2b515b3615eecca128f5f90b05d1786dc&timestamp=1449649715&nonce=2109185587"

func init() {
	Initialize(originId, appId, appSecret, token, encodingAESKey)
	log.Info("Initializing, please wait 5 seconds...")
	time.Sleep(5 * time.Second) // waiting to refresh token
}
