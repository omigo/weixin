package weixin

import "testing"

func TestCreateTemporaryQRCodeTicket(t *testing.T) {
	ticket, err := CreateTemporaryQRCodeTicket(1)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v", ticket)
}

func TestCreatePermanentQRCodeTicket(t *testing.T) {
	ticket, err := CreatePermanentQRCodeTicket(1)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v", ticket)
}

func TestCreatePermanentQRCodeTicketString(t *testing.T) {
	ticket, err := CreatePermanentQRCodeTicketString("1")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v", ticket)
}

func TestGetQRCodeImg(t *testing.T) {
	img := GetQRCodeImg("gQG_8DoAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL20zV18zNW5sTjF0T3M3cFo5MTBLAAIEQ51mVgMEAAAAAA==")

	t.Logf("%#v", img)
}
