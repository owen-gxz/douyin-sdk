package shop

import (
	"bytes"
	"encoding/json"
	"github.com/owen-gxz/douyin-sdk/util"
)

const (
	matchUrl = "https://open.douyin.com/goodlife/v1/poi/match/task/submit/"
)

type MatchReq struct {
	Datas []MatchData `json:"datas"`
}

type MatchData struct {
	ExtId     string `json:"ext_id"`
	PoiId     string `json:"poi_id"`
	PoiName   string `json:"poi_name"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Address   string `json:"address"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type MatchResp struct {
	Data struct {
		TaskId      int64  `json:"task_id"`
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

// Match 提交门店匹配任务，匹配结果通过“查询门店匹配结果”获取 (无用)
func Match(accountToken string, req MatchReq) (*MatchResp, error) {
	data, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	resp := &MatchResp{}
	err = util.Post2Response2(matchUrl, accountToken, bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
