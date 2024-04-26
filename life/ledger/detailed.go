package ledger

import (
	"bytes"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

const uri = "https://open.douyin.com/goodlife/v1/settle/ledger/query_record_by_cert/"

type LedgerDetailedResp struct {
	Data struct {
		Records []struct {
			CertificateId string `json:"certificate_id"`
			FundAmount    struct {
				// 商家收款金额
				Goods       int `json:"goods"`
				PayHandling int `json:"pay_handling"`
				//达人佣金
				TalentCommission int `json:"talent_commission"`
				//   服务商获得的总佣金
				TotalAgentMerchant           int `json:"total_agent_merchant"`
				TotalMerchantPlatformService int `json:"total_merchant_platform_service"`
			} `json:"fund_amount"`
			LedgerId   string `json:"ledger_id"`
			OrderId    string `json:"order_id"`
			VerifyId   string `json:"verify_id"`
			VerifyTime int    `json:"verify_time"`
			Amount     struct {
				Pay                 int `json:"pay"`
				PayDiscount         int `json:"pay_discount"`
				TotalOperationAgent int `json:"total_operation_agent"`
				MerchantTicket      int `json:"merchant_ticket"`
				Original            int `json:"original"`
			} `json:"amount"`
			Code   string `json:"code"`
			Cursor string `json:"cursor"`
			Sku    struct {
				Title         string `json:"title"`
				MarketPrice   int    `json:"market_price"`
				SkuId         string `json:"sku_id"`
				SoldStartTime int    `json:"sold_start_time"`
				ThirdSkuId    string `json:"third_sku_id"`
			} `json:"sku"`
			//分账单状态，0表示初始化、1表示分账中、2表示分账成功、3表示分账失败、4表示分账单已废弃
			Status int `json:"status"`
		} `json:"records"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Extra struct {
		Now            int    `json:"now"`
		SubDescription string `json:"sub_description"`
		SubErrorCode   int    `json:"sub_error_code"`
		Description    string `json:"description"`
		ErrorCode      int    `json:"error_code"`
		Logid          string `json:"logid"`
	} `json:"extra"`
}

// list 创建代运营订单
func LedgerDetailed(accountToken string, certificateID ...string) (*LedgerDetailedResp, error) {
	var buf bytes.Buffer
	buf.WriteString(uri)
	v := url.Values{}
	v.Add("certificate_ids", strings.Join(certificateID[:], ","))
	if strings.Contains(uri, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &LedgerDetailedResp{}
	err := util.Get2Response2(buf.String(), accountToken, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
