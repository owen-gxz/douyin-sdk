package goods

import (
	"bytes"
	"encoding/json"
	"github.com/owen-gxz/douyin-sdk/util"
)

//https://open.douyin.com/goodlife/v1/goods/product/operate/

type OperateReq struct {
	OpType    int    `json:"op_type"`
	ProductID string `json:"product_id"`
	AccountID string `json:"account_id"`
}

type OperateResp struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Extra struct {
		ErrorCode      int    `json:"error_code"`
		Description    string `json:"description"`
		SubErrorCode   int    `json:"sub_error_code"`
		SubDescription string `json:"sub_description"`
		Logid          string `json:"logid"`
		Now            int    `json:"now"`
	} `json:"extra"`
}

// Create 创建代运营订单
func Operate(accountToken string, req OperateReq) (*OperateResp, error) {
	data, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	resp := &OperateResp{}
	err = util.Post2Response2(operateGoodsUrl, accountToken, bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
