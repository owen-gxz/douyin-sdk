package commission

import (
	"bytes"
	"encoding/json"
	"github.com/owen-gxz/douyin-sdk/util"
)

type ListReq struct {
	SpuId    int64 `json:"spu_id"`
	PageNo   int64 `json:"page_no"`
	PageSize int64 `json:"page_size"`
}

type ListResp struct {
	Data struct {
		Data []struct {
			CommissionRate int    `json:"commission_rate"`
			ContentType    int    `json:"content_type"`
			CreateTime     string `json:"create_time"`
			PlanId         int64  `json:"plan_id,number"`
			Status         int    `json:"status"`
		} `json:"data"`
		PageCount   int    `json:"page_count"`
		Total       int    `json:"total"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
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
func List(accountToken string, productID, page, size int64) (*ListResp, error) {
	req := ListReq{
		SpuId:    productID,
		PageNo:   page,
		PageSize: size,
	}
	data, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	resp := &ListResp{}
	err = util.Post2Response2(listPlanUri, accountToken, bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
