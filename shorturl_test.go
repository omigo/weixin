package weixin

import "testing"

func TestLongURL2Short(t *testing.T) {
	shortURL, err := LongURL2Short("http://wap.koudaitong.com/v2/showcase/goods?alias=128wi9shh&spm=h56083&redirect_count=1")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v", shortURL)
}
