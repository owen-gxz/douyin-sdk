package partner

import (
	"bytes"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

type InfoReq struct {
	OrderID string `json:"order_id"`
	//AccountID           string `json:"page,omitempty"`
}

type InfoResp struct {
	Data struct {
		ChargeType         int    `json:"charge_type"`
		EndTime            int    `json:"end_time"`
		MerchantName       string `json:"merchant_name"`
		ProductName        string `json:"product_name"`
		StartTime          int    `json:"start_time"`
		Description        string `json:"description"`
		AccountId          string `json:"account_id"`
		CooperationContent int    `json:"cooperation_content"`
		CreateTime         int    `json:"create_time"`
		Id                 string `json:"id"`
		Status             int    `json:"status"`
		ErrorCode          int    `json:"error_code"`
		CommissionRatio    string `json:"commission_ratio"`
	} `json:"data"`
	Extra struct {
		Description    string `json:"description"`
		ErrorCode      int    `json:"error_code"`
		Logid          string `json:"logid"`
		Now            int    `json:"now"`
		SubDescription string `json:"sub_description"`
		SubErrorCode   int    `json:"sub_error_code"`
	} `json:"extra"`
}

// list 创建代运营订单
func Info(accountToken string, orderID string) (*InfoResp, error) {
	var buf bytes.Buffer
	buf.WriteString(infoPartnerOrderUrl)
	v := url.Values{}
	v.Add("order_id", orderID)
	if strings.Contains(infoPartnerOrderUrl, "?") {
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
