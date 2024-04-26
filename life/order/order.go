package order

import (
	"bytes"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

const (
	orderInfoUrl = "https://open.douyin.com/goodlife/v1/trade/order/query/"
)

type InfoResp struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		Orders      []struct {
			PayTime         int    `json:"pay_time"`
			SkuId           string `json:"sku_id"`
			ThirdSkuId      string `json:"third_sku_id"`
			UpdateOrderTime int    `json:"update_order_time"`
			Count           int    `json:"count"`
			DiscountAmount  int    `json:"discount_amount"`
			OrderId         string `json:"order_id"`
			OriginalAmount  int    `json:"original_amount"`
			Products        []struct {
				Commodities []struct {
					ItemList []struct {
						ItemTag  string      `json:"item_tag"`
						Name     string      `json:"name"`
						Price    int         `json:"price"`
						Unit     string      `json:"unit"`
						AttrList interface{} `json:"attr_list"`
						Count    int         `json:"count"`
						Desc     string      `json:"desc"`
					} `json:"item_list"`
					OptionCount int    `json:"option_count"`
					TotalCount  int    `json:"total_count"`
					GroupName   string `json:"group_name"`
				} `json:"commodities"`
				Num         int    `json:"num"`
				ProductId   string `json:"product_id"`
				ProductName string `json:"product_name"`
				SkuId       string `json:"sku_id"`
			} `json:"products"`
			ReceiverInfo struct {
				ReceiverPhone   string `json:"receiver_phone"`
				SecretNumber    string `json:"secret_number"`
				City            string `json:"city"`
				District        string `json:"district"`
				DoorPlateNum    string `json:"door_plate_num"`
				Lat             int    `json:"lat"`
				Lng             int    `json:"lng"`
				LocationAddress string `json:"location_address"`
				Town            string `json:"town"`
				LocationName    string `json:"location_name"`
				Province        string `json:"province"`
				ReceiverName    string `json:"receiver_name"`
			} `json:"receiver_info"`
			AmountInfo struct {
				MerchantDiscountAmount int `json:"merchant_discount_amount"`
				OriginAmount           int `json:"origin_amount"`
				PayAmount              int `json:"pay_amount"`
				PayDiscountAmount      int `json:"pay_discount_amount"`
				PlatformDiscountAmount int `json:"platform_discount_amount"`
				ProductOriginAmount    int `json:"product_origin_amount"`
			} `json:"amount_info"`
			Certificate []struct {
				RefundTime     int    `json:"refund_time"`
				CertificateId  string `json:"certificate_id"`
				ItemStatus     int    `json:"item_status"`
				ItemUpdateTime int    `json:"item_update_time"`
				OrderItemId    string `json:"order_item_id"`
				RefundAmount   int    `json:"refund_amount"`
			} `json:"certificate"`
			DeliveryInfo struct {
				SysExpectTime  string `json:"sys_expect_time"`
				TableWare      string `json:"table_ware"`
				UserExpectTime string `json:"user_expect_time"`
				DeliverModel   int    `json:"deliver_model"`
				IsBook         bool   `json:"is_book"`
				Remark         string `json:"remark"`
				ShopNumber     string `json:"shop_number"`
			} `json:"delivery_info"`
			PoiId string `json:"poi_id"`
			Poi   struct {
				PoiId   string `json:"poi_id"`
				PoiName string `json:"poi_name"`
			} `json:"poi"`
			PayAmount       int    `json:"pay_amount"`
			SkuName         string `json:"sku_name"`
			CreateOrderTime int    `json:"create_order_time"`
			OpenId          string `json:"open_id"`
			OrderStatus     int    `json:"order_status"`
			OrderType       int    `json:"order_type"`
			Contacts        []struct {
				Phone string `json:"phone"`
				Name  string `json:"name"`
			} `json:"contacts"`
			MerchantInfo struct {
				AccountId   string `json:"account_id"`
				AccountName string `json:"account_name"`
			} `json:"merchant_info"`
			SubOrderAmountInfos []struct {
				PayAmount      int    `json:"pay_amount"`
				SubOrderId     string `json:"sub_order_id"`
				SubOrderType   int    `json:"sub_order_type"`
				DiscountAmount int    `json:"discount_amount"`
				Discounts      []struct {
					MerchantDiscountAmount int `json:"merchant_discount_amount"`
					PlatformDiscountAmount int `json:"platform_discount_amount"`
					DiscountAmount         int `json:"discount_amount"`
					DiscountType           int `json:"discount_type"`
				} `json:"discounts"`
				OriginAmount int `json:"origin_amount"`
			} `json:"sub_order_amount_infos"`
		} `json:"orders"`
		Page struct {
			PageNum  int `json:"page_num"`
			PageSize int `json:"page_size"`
			Total    int `json:"total"`
		} `json:"page"`
	} `json:"data"`
	Extra struct {
		Logid          string `json:"logid"`
		ErrorCode      int    `json:"error_code"`
		Description    string `json:"description"`
		SubErrorCode   int    `json:"sub_error_code"`
		SubDescription string `json:"sub_description"`
		Now            int    `json:"now"`
	} `json:"extra"`
}

// list 创建代运营订单
func Info(accountToken string, accountID, orderID, pageNum string) (*InfoResp, error) {
	var buf bytes.Buffer
	buf.WriteString(orderInfoUrl)
	v := url.Values{}
	v.Add("order_id", orderID)
	v.Add("page_num", pageNum)
	v.Add("page_size", "10")
	v.Add("get_secret_number", "true")
	v.Add("account_id", accountID)
	if strings.Contains(orderInfoUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &InfoResp{}
	err := util.Get2Response2(buf.String(), accountToken, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
