package shop

import (
	"fmt"
	"github.com/owen-gxz/douyin-sdk/oauth"
	"github.com/owen-gxz/douyin-sdk/service"
	"testing"
)

func getCli() *service.Service {
	return service.NewService(&oauth.Config{
		ClientKey:    "awfu87gyyi5ktef3",
		ClientSecret: "953e37e3f8172bf5d9adbfb8730ec123",
	}, nil)
}

func TestQuery(t *testing.T) {
	cli := getCli()
	qr, err := Query(cli.ClientToken(), "7126518681177163816", "1", "1")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(qr)
}
