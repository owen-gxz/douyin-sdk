package partner

import (
	"bytes"
	"encoding/json"
	"github.com/owen-gxz/douyin-sdk/util"
)

type CreateCommissionReq struct {
	OrderId      string `json:"order_id"`
	ProductItems []struct {
		CommissionRatio string `json:"commission_ratio"`
		ProductId       string `json:"product_id"`
	} `json:"product_items"`
}

type CreateCommissionResp struct {
	Data struct {
		InvalidCommissionProductList []struct {
			AllianceCommissionRatio string `json:"alliance_commission_ratio"`
			CommissionRatio         string `json:"commission_ratio"`
			ProductId               string `json:"product_id"`
		} `json:"invalid_commission_product_list"`
		RecordId    string `json:"record_id"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Extra struct {
		Logid          string `json:"logid"`
		Now            int    `json:"now"`
		SubDescription string `json:"sub_description"`
		SubErrorCode   int    `json:"sub_error_code"`
		Description    string `json:"description"`
		ErrorCode      int    `json:"error_code"`
	} `json:"extra"`
}

// Create 创建代运营订单
func CreateCommission(accountToken string, req CreateCommissionReq) (*CreateCommissionResp, error) {
	data, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	resp := &CreateCommissionResp{}
	err = util.Post2Response2(createCommissionProductUrl, accountToken, bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
