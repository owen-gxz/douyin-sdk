package service

import (
	"fmt"
	"github.com/owen-gxz/douyin-sdk/life/partner"
	"testing"
	"time"
)

func TestService_PartnerInfo(t *testing.T) {

}

// 8000008079389676546 外卖订单
func TestService_OrderInfo(t *testing.T) {
	feige := getCli()
	oir, err := feige.OrderInfo("7249644892274575392", "1003295345414420877", "10")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(oir)
}

func TestService_OrderInfo2(t *testing.T) {
	feige := getCli()
	//148
	oir, err := feige.LedgerDetailed("7257842953208152119")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(oir)
}

func TestService_OrderInfo3(t *testing.T) {
	feige := getCli()
	//148
	oir, err := feige.LedgerDetailedBySubFulfilID("7230286522502922252")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(oir)
}

func TestService_Category(t *testing.T) {
	feige := getCli()
	//148
	oir, err := feige.CategoryList("7207323963586775080")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(oir)
}

func TestService_Template(t *testing.T) {
	feige := getCli()
	//148
	oir, err := feige.Template("1", "1001001")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(oir)
}

func TestService_CreatePartner(t *testing.T) {
	feige := getCli()
	//148

	start := time.Now()
	oir, err := feige.CreatePartner(partner.CreateReq{
		AccountId:          "6948730580385187843",
		CooperationContent: 5,
		StartTime:          start.Add(10 * time.Second).Unix(),
		EndTime:            start.AddDate(0, 11, 0).Unix(),
		ChargeType:         1,
		CommissionRatio:    "1000",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(oir)
}

func TestService_Commission(t *testing.T) {
	cli := getCli()
	cr, err := cli.Commission("7229258005718583333", "1")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(cr)
}
