package partner

import (
	"bytes"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

type CommissionsResp struct {
	Data struct {
		ProductCommissions []struct {
			LatestCommissionRecordItem struct {
				AuditCommissionRatio  string `json:"audit_commission_ratio"`
				CommissionAuditStatus int    `json:"commission_audit_status"`
				CommissionRatioBefore string `json:"commission_ratio_before"`
				ItemId                string `json:"item_id"`
				MerchantAckTime       int    `json:"merchant_ack_time"`
				ProductId             string `json:"product_id"`
			} `json:"latest_commission_record_item"`
			Product struct {
				ProductId                string `json:"product_id"`
				ProductName              string `json:"product_name"`
				ProductStatus            int    `json:"product_status"`
				ActualPrice              string `json:"actual_price"`
				EffectiveCommissionRatio string `json:"effective_commission_ratio"`
				ProductCommissionStatus  int    `json:"product_commission_status"`
			} `json:"product"`
		} `json:"product_commissions"`
		Total       int    `json:"total"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Extra struct {
		SubErrorCode   int    `json:"sub_error_code"`
		Description    string `json:"description"`
		ErrorCode      int    `json:"error_code"`
		Logid          string `json:"logid"`
		Now            int    `json:"now"`
		SubDescription string `json:"sub_description"`
	} `json:"extra"`
}

func Commissions(accountToken string, orderID, page string) (*CommissionsResp, error) {
	var buf bytes.Buffer
	buf.WriteString(commissionProductUrl)
	v := url.Values{}
	v.Add("order_id", orderID)
	v.Add("page", page)
	v.Add("size", "20")
	if strings.Contains(commissionProductUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &CommissionsResp{}
	err := util.Get2Response2(buf.String(), accountToken, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
