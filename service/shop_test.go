package service

import (
	"fmt"
	"github.com/owen-gxz/douyin-sdk/life/shop"
	"github.com/owen-gxz/douyin-sdk/oauth"
	"testing"
)

func getCli() *Service {
	return NewService(&oauth.Config{
		ClientKey:    "awyxo5cmzmfu0z1y",
		ClientSecret: "e035ee8ed08f2f3e8bc52f8671f3cd19",
	}, nil)
}

// 7217321267504875553
func TestService_ShopInfo(t *testing.T) {
	cli := getCli()
	si, err := cli.ShopInfo("7217321267504875553")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(si[0].Poi)
}

func TestService_Match(t *testing.T) {
	cli := getCli()
	mr, err := cli.Match(shop.MatchReq{
		Datas: []shop.MatchData{
			shop.MatchData{
				ExtId:     "cggkmnogb9eje5kfpi20",
				PoiId:     "7080822239024646185",
				PoiName:   "港饮之港(河大新区店)",
				Address:   "志学路河北大学坤舆德翰园小区25号底商",
				Longitude: "115.571230",
				Latitude:  "38.879386",
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(mr)
}

func TestService_Super(t *testing.T) {
	cli := getCli()
	mr, err := cli.Supplier("6706742652081145868")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(mr)
}
