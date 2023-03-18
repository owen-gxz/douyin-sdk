package partner

import (
	"fmt"
	"github.com/owen-gxz/douyin-sdk/oauth"
	"github.com/owen-gxz/douyin-sdk/service"
	"testing"
	"time"
)

func getCli() *service.Service {
	return service.NewService(&oauth.Config{
		ClientKey:    "awfu87gyyi5ktef3",
		ClientSecret: "953e37e3f8172bf5d9adbfb8730ec123",
	}, nil)
}

func TestSendMessage(t *testing.T) {
	sr := getCli()
	start := time.Now()
	token := sr.ClientToken()
	fmt.Println(token)
	cr, err := Create(token, CreateReq{
		AccountId:          "7126518681177163816",
		CooperationContent: 5,
		StartTime:          start.AddDate(0, 0, 1).Unix(),
		EndTime:            start.AddDate(1, 0, 0).Unix(),
		ChargeType:         1,
		CommissionRatio:    "1000",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(cr)

}
