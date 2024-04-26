package ledger

import (
	"bytes"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

type LedgerDetailedBySubFulfilIDResp struct {
	Extra struct {
		ErrorCode      int    `json:"error_code"`
		Description    string `json:"description"`
		SubErrorCode   int    `json:"sub_error_code"`
		SubDescription string `json:"sub_description"`
		Logid          string `json:"logid"`
		Now            int    `json:"now"`
	} `json:"extra"`
	Data struct {
		Records []struct {
			AccountName          string                       `json:"account_name"`
			LedgerId             string                       `json:"ledger_id"`
			AccountId            string                       `json:"account_id"`
			LedgerPlatformTicket int                          `json:"ledger_platform_ticket"`
			PayAmount            int                          `json:"pay_amount"`
			SubFulfilId          string                       `json:"sub_fulfil_id"`
			WithdrawId           string                       `json:"withdraw_id"`
			BizTime              int                          `json:"biz_time"`
			CommissionBaseAmount int                          `json:"commission_base_amount"`
			ItemOrderId          string                       `json:"item_order_id"`
			ShopOrderId          string                       `json:"shop_order_id"`
			LedgerDatailMap      map[string]ledgerDetailedMap `json:"ledger_datail_map"`
			SkuId                string                       `json:"sku_id"`
			SkuType              int                          `json:"sku_type"`
		} `json:"records"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
}

type ledgerDetailedMap struct {
	Amount          int64 `json:"amount"`
	CommissionRatio int   `json:"commission_ratio"`
}

var (
	subFulfilIDUri                            = "https://open.douyin.com/goodlife/v1/settle/bill/query_record_by_subfulfil/"
	LedgerDetailedMapPay                      = "pay_handling"                 // 支付手续费
	LedgerDetailedMapGood                     = "goods"                        // 商家收款金额
	LedgerDetailedMapMerchantPlatform         = "merchant_platform_commission" // 抖音平台抽佣
	LedgerDetailedMapOperationAgentCommission = "operation_agent_commission"   // 服务商佣金

)

func LedgerDetailedBySubFulfilID(accountToken string, subFulfilID string) (*LedgerDetailedBySubFulfilIDResp, error) {
	var buf bytes.Buffer
	buf.WriteString(subFulfilIDUri)
	v := url.Values{}
	v.Add("sub_fulfil_id", subFulfilID)
	if strings.Contains(subFulfilIDUri, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &LedgerDetailedBySubFulfilIDResp{}
	err := util.Get2Response2(buf.String(), accountToken, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
