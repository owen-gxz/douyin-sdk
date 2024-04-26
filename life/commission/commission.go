package commission

import (
	"bytes"
	"encoding/json"
	"github.com/owen-gxz/douyin-sdk/util"
)

var (
	cpsPlanUri  = "https://open.douyin.com/poi/common/plan/save/"
	listPlanUri = "https://open.douyin.com/poi/plan/list/"
)

type CpsCommissionReq struct {
	CommissionRate int    `json:"commission_rate"`
	ContentType    int    `json:"content_type"`
	PlanId1        string `json:"plan_id1,omitempty"`
	PlanId         int64  `json:"plan_id"`
	SpuId1         string `json:"spu_id1,omitempty"`
	SpuId          int64  `json:"spu_id"`
}

type CpsCommissionResp struct {
	Data struct {
		PlanId      int64  `json:"plan_id"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Extra struct {
		ErrorCode      int    `json:"error_code"`
		Logid          string `json:"logid"`
		Now            int    `json:"now"`
		SubDescription string `json:"sub_description"`
		SubErrorCode   int    `json:"sub_error_code"`
		Description    string `json:"description"`
	} `json:"extra"`
}

func CpsCommission(token string, req CpsCommissionReq) (*CpsCommissionResp, error) {
	data, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	resp := &CpsCommissionResp{}
	err = util.Post2Response2(cpsPlanUri, token, bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
