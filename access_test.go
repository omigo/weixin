package weixin

import (
	"net/http"
	"testing"
)

const (
	addr = ":3080"
)

func TestValidateURL(t *testing.T) {
	http.HandleFunc("/weixin", HandleAccess)

	http.ListenAndServe(addr, nil)
}
