package service

import (
	"fmt"
	"testing"
)

// account_id=7205159428390946855&count=10&cursor=&goods_creator_type=0&state=1
// 20   22
func TestService_Goods(t *testing.T) {
	cli := getCli()
	rc, err := cli.GoodsInfo("7027817458694096910", "7238536627658835983")
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Println(len(rc.Data.Products))
	fmt.Println(rc.Extra.Logid)
}
