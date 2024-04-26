package goods

import (
	"bytes"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

type SkuInfoResp struct {
	Extra struct {
		Description    string `json:"description"`
		SubErrorCode   int    `json:"sub_error_code"`
		SubDescription string `json:"sub_description"`
		Now            int    `json:"now"`
		Logid          string `json:"logid"`
		ErrorCode      int    `json:"error_code"`
	} `json:"extra"`
	Data struct {
		Description    string `json:"description"`
		ErrorCode      int    `json:"error_code"`
		ProductOnlines []struct {
			CommissionInfo struct {
				PlatformTakeRate int `json:"platform_take_rate"`
			} `json:"commission_info"`
			OnlineStatus int `json:"online_status"`
			Product      struct {
				BizLine          int               `json:"biz_line"`
				OutId            string            `json:"out_id"`
				ProductId        string            `json:"product_id"`
				Telephone        []string          `json:"telephone"`
				AccountName      string            `json:"account_name"`
				AttrKeyValueMap  map[string]string `json:"attr_key_value_map"`
				SoldStartTime    int               `json:"sold_start_time"`
				CategoryFullName string            `json:"category_full_name"`
				CategoryId       int               `json:"category_id"`
				CreatorAccountId int64             `json:"creator_account_id"`
				ProductName      string            `json:"product_name"`
				SoldEndTime      int               `json:"sold_end_time"`
				OwnerAccountId   int64             `json:"owner_account_id"`
				Pois             []struct {
					SupplierExtId interface{} `json:"supplier_ext_id"`
					SupplierId    interface{} `json:"supplier_id"`
					PoiId         string      `json:"poi_id"`
				} `json:"pois"`
				ProductType int `json:"product_type"`
			} `json:"product"`
			Sku struct {
				CreateTime   int    `json:"create_time"`
				SkuId        string `json:"sku_id"`
				SkuName      string `json:"sku_name"`
				Status       int    `json:"status"`
				UpdateTime   int    `json:"update_time"`
				ActualAmount int    `json:"actual_amount"`
				OriginAmount int    `json:"origin_amount"`
				Stock        struct {
					LimitType int `json:"limit_type"`
					StockQty  int `json:"stock_qty"`
				} `json:"stock"`
				AttrKeyValueMap map[string]string `json:"attr_key_value_map"`
			} `json:"sku"`
		} `json:"product_onlines"`
	} `json:"data"`
}

// list 创建代运营订单
func SkuInfo(accountToken string, accountID, productID string) (*SkuInfoResp, error) {
	var buf bytes.Buffer
	buf.WriteString(SkuInfoUrl)
	v := url.Values{
		"product_ids": {productID},
		"account_id":  {accountID},
	}
	if strings.Contains(SkuInfoUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &SkuInfoResp{}
	err := util.Get2Response2(buf.String(), accountToken, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
