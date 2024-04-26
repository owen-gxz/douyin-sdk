package service

import (
	"fmt"
	"github.com/owen-gxz/douyin-sdk/life/commission"
	"testing"
)

func TestService_CpsCommission(t *testing.T) {
	cli := getCli()
	ccr, err := cli.CpsCommission(commission.CpsCommissionReq{
		SpuId:          1765123501986832,
		PlanId:         7230640063537154085,
		ContentType:    3,
		CommissionRate: 200,

		//SpuId:          7238536627658836000,
		//PlanId:         7240279287369370000,
		//ContentType:    3,
		//CommissionRate: 200,
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ccr)
}

func TestService_CpsList(t *testing.T) {
	cli := getCli()
	ccr, err := cli.CpsList(1765046529079375, 1)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ccr)
}
