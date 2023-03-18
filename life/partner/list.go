package partner

import (
	"bytes"
	"encoding/json"
	"github.com/owen-gxz/douyin-sdk/util"
)

type ListReq struct {
	CooperationContents string `json:"cooperation_contents,omitempty"`
	StartTime           int64  `json:"start_time,omitempty"`
	EndTime             int64  `json:"end_time,omitempty"`
	Status              string `json:"status,omitempty"`
	AccountID           string `json:"account_id,omitempty"`
	Page                int64  `json:"page,omitempty"`
	Size                int64  `json:"size,omitempty"`
	//AccountID           string `json:"page,omitempty"`
}

type ListResp struct {
	Extra struct {
		SubErrorCode   int    `json:"sub_error_code"`
		Description    string `json:"description"`
		ErrorCode      int    `json:"error_code"`
		Logid          string `json:"logid"`
		Now            int    `json:"now"`
		SubDescription string `json:"sub_description"`
	} `json:"extra"`
	Data struct {
		Orders []struct {
			Status             int    `json:"status"`
			AccountId          string `json:"account_id"`
			ChargeType         int    `json:"charge_type"`
			CooperationContent int    `json:"cooperation_content"`
			CreateTime         int    `json:"create_time"`
			EndTime            int    `json:"end_time"`
			Id                 string `json:"id"`
			MerchantName       string `json:"merchant_name"`
			ProductName        string `json:"product_name"`
			StartTime          int    `json:"start_time"`
		} `json:"orders"`
		Total       int    `json:"total"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
}

// list 创建代运营订单
func List(accountToken string, accountID string, page, size int64) (*ListResp, error) {
	req := ListReq{
		AccountID: accountID,
		Page:      page,
		Size:      size,
	}
	data, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	resp := &ListResp{}
	err = util.Post2Response2(listPartnerOrderUrl, accountToken, bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
