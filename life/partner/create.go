package partner

import (
	"bytes"
	"encoding/json"
	"github.com/owen-gxz/douyin-sdk/util"
)

const (
	createPartnerOrderUrl = "https://open.douyin.com/goodlife/v1/partner/order/create/"
	listPartnerOrderUrl   = "https://open.douyin.com/goodlife/v1/partner/order/query/"
	infoPartnerOrderUrl   = "https://open.douyin.com/goodlife/v1/partner/order/get/"
)

type CreateReq struct {
	AccountId          string `json:"account_id"`
	CooperationContent int    `json:"cooperation_content"`
	StartTime          int64  `json:"start_time"`
	EndTime            int64  `json:"end_time"`
	ChargeType         int    `json:"charge_type"`
	CommissionRatio    string `json:"commission_ratio"`
}

type CreateResp struct {
	Data struct {
		OrderId     string `json:"order_id"`
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

// Create 创建代运营订单
func Create(accountToken string, req CreateReq) (*CreateResp, error) {
	data, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	resp := &CreateResp{}
	err = util.Post2Response2(createPartnerOrderUrl, accountToken, bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
