package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
)

var (
	eventStateList   = fmt.Sprintf("%s/event/status/list/", util.ServiceUrl)
	eventStateUpdate = fmt.Sprintf("%s/event/status/update/", util.ServiceUrl)
)

type EventStateListResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		List        []struct {
			Event  string `json:"event"`
			Status int    `json:"status"`
		} `json:"list"`
	} `json:"data"`
}

// 事件状态
func (s Service) EventStateList() (*EventStateListResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(eventStateList)
	v := url.Values{
		"access_token": {s.ClientToken()},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &EventStateListResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type EventStateUpdateRequest struct {
	List []struct {
		Event  string `json:"event"`
		Status int    `json:"status"`
	} `json:"list"`
}
type EventStateUpdateResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
}

// 更改事件推送状态
func (s Service) EventStateUpdate(list EventStateUpdateRequest) (*EventStateUpdateResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(eventStateUpdate)
	v := url.Values{
		"access_token": {s.ClientToken()},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &EventStateUpdateResponse{}
	md, err := json.Marshal(&list)
	if err != nil {
		return nil, err
	}
	err = util.Post2Response(buf.String(), bytes.NewReader(md), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
