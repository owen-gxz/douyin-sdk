package shop

import (
	"bytes"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

const (
	shopQueryUrl = "https://open.douyin.com/goodlife/v1/shop/poi/query/"
)

type QueryResp struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		Total       int    `json:"total"`
		Pois        []struct {
			Poi struct {
				PoiId     string  `json:"poi_id"`
				Address   string  `json:"address"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				PoiName   string  `json:"poi_name"`
			} `json:"poi"`
		} `json:"pois"`
	} `json:"data"`
	Extra struct {
		ErrorCode      int    `json:"error_code"`
		Description    string `json:"description"`
		SubErrorCode   int    `json:"sub_error_code"`
		SubDescription string `json:"sub_description"`
		Logid          string `json:"logid"`
		Now            int    `json:"now"`
	} `json:"extra"`
}

// Create 创建代运营订单
func Query(accountToken string, accountID string, page, size string) (*QueryResp, error) {
	var buf bytes.Buffer
	buf.WriteString(shopQueryUrl)
	v := url.Values{
		"page": {page},
		"size": {size},
	}
	if accountID != "" {
		v.Add("account_id", accountID)
	}
	if strings.Contains(shopQueryUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &QueryResp{}
	fmt.Println(buf.String())
	err := util.Get2Response2(buf.String(), accountToken, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
