package goods

import (
	"bytes"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

type ListSkuResp struct {
	BaseResp struct {
		StatusCode    int    `json:"StatusCode"`
		StatusMessage string `json:"StatusMessage"`
	} `json:"BaseResp"`
	Data struct {
		Products []struct {
			Product struct {
				CategoryId       int    `json:"category_id"`
				ProductType      int    `json:"product_type"`
				SoldEndTime      int    `json:"sold_end_time"`
				AccountName      string `json:"account_name"`
				CategoryFullName string `json:"category_full_name"`
				SoldStartTime    int    `json:"sold_start_time"`
				CreateTime       int    `json:"create_time"`
				ProductName      string `json:"product_name"`
				Pois             []struct {
					PoiId         string `json:"poi_id"`
					SupplierExtId string `json:"supplier_ext_id"`
				} `json:"pois"`
				ProductId        string `json:"product_id"`
				OutID            string `json:"out_id"`
				UpdateTime       int    `json:"update_time"`
				CreatorAccountId int64  `json:"creator_account_id"`
				OwnerAccountId   int64  `json:"owner_account_id"`
				AttrKeyValueMap  struct {
					CanNoUseDate          string `json:"can_no_use_date"`
					ImageList             string `json:"image_list"`
					RecPersonNum          string `json:"rec_person_num"`
					ShowChannel           string `json:"show_channel"`
					SuperimposedDiscounts string `json:"superimposed_discounts"`
					UseTime               string `json:"use_time"`
					Appointment           string `json:"appointment"`
					ComsumptionThreshold  string `json:"comsumption_threshold"`
					CustomerReservedInfo  string `json:"customer_reserved_info"`
					RealNameInfo          string `json:"real_name_info"`
					UseDate               string `json:"use_date"`
					AutoRenew             string `json:"auto_renew"`
				} `json:"attr_key_value_map"`
				BizLine int `json:"biz_line"`
			} `json:"product"`
			Sku struct {
				OriginAmount int    `json:"origin_amount"`
				SkuId        string `json:"sku_id"`
				SkuName      string `json:"sku_name"`
				Status       int    `json:"status"`
				Stock        struct {
					StockQty  int `json:"stock_qty"`
					LimitType int `json:"limit_type"`
				} `json:"stock"`
				ActualAmount    int `json:"actual_amount"`
				AttrKeyValueMap struct {
					RefundNeedMerchantConfirm string `json:"refund_need_merchant_confirm"`
					RefundType                string `json:"refund_type"`
					SettleType                string `json:"settle_type"`
					UseType                   string `json:"use_type"`
					CodeSourceType            string `json:"code_source_type"`
					Commodity                 string `json:"commodity"`
					LimitRule                 string `json:"limit_rule"`
					IsOriginAmountEdited      string `json:"is_origin_amount_edited"`
					TakeRate                  string `json:"take_rate"`
				} `json:"attr_key_value_map"`
				CreateTime int    `json:"create_time"`
				Extra      string `json:"extra"`
				UpdateTime int    `json:"update_time"`
			} `json:"sku"`
			Skus           interface{} `json:"skus"`
			CommissionInfo struct {
				PlatformTakeRate int `json:"platform_take_rate"`
			} `json:"commission_info"`
			OnlineStatus int `json:"online_status"`
		} `json:"products"`
		Description string `json:"description"`
		ErrorCode   int    `json:"error_code"`
		HasMore     bool   `json:"has_more"`
		NextCursor  string `json:"next_cursor"`
	} `json:"data"`
	Extra struct {
		ErrorCode      int    `json:"error_code"`
		Description    string `json:"description"`
		SubErrorCode   int    `json:"sub_error_code"`
		SubDescription string `json:"sub_description"`
		Now            int    `json:"now"`
		Logid          string `json:"logid"`
	} `json:"extra"`
}

// list 创建代运营订单
func SkuList(accountToken string, accountID, cursor, status, goodsCreatorType string) (*ListSkuResp, error) {
	var buf bytes.Buffer
	buf.WriteString(listSkuUrl)
	v := url.Values{
		"cursor":             {cursor},
		"count":              {"4"},
		"goods_creator_type": {goodsCreatorType},
	}
	if accountID != "" {
		v.Add("account_id", accountID)
	}
	if status != "" {
		v.Add("status", status)
	}
	if strings.Contains(listSkuUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	fmt.Println(buf.String())
	resp := &ListSkuResp{}
	err := util.Get2Response2(buf.String(), accountToken, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
