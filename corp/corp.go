package corp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
)

var (
	leadsUserListUrl       = fmt.Sprintf("%s/enterprise/leads/user/list/", util.ServiceUrl)
	leadsUserDetailUrl     = fmt.Sprintf("%s/enterprise/leads/user/detail/", util.ServiceUrl)
	leadsUserActionListUrl = fmt.Sprintf("%s/enterprise/leads/user/action/list/", util.ServiceUrl)
	leadsTagListUrl        = fmt.Sprintf("%s/enterprise/leads/tag/list/", util.ServiceUrl)
	leadsTagUserListUrl    = fmt.Sprintf("%s/enterprise/leads/tag/user/list/", util.ServiceUrl)
	leadsTagCreateUrl      = fmt.Sprintf("%s/enterprise/leads/tag/create/", util.ServiceUrl)
	leadsTagUpdateUrl      = fmt.Sprintf("%s/enterprise/leads/tag/update/", util.ServiceUrl)
	leadsTagDeleteUrl      = fmt.Sprintf("%s/enterprise/leads/tag/delete/", util.ServiceUrl)
	leadsTagUserUpdateUrl  = fmt.Sprintf("%s/enterprise/leads/tag/user/update/", util.ServiceUrl)
)

const (
	//leads_level 相关：
	// 没兴趣
	NoInterestLeadsLevel = iota - 1
	// 了解
	UnderstandLeadsLevel
	// 有兴趣
	InterestLeadsLevel
	// 有意愿
	DesireLeadsLevel

	// 已转换
	TransformLeadsLevel = 10

	//action_type 相关：
	AllActionType = iota
	//1 - 私信互动
	PrivateActionType
	//2 - 组件互动
	ClusterActionType
	//3 - 主页互动
	MainActionType
)

type LeadsUserListResponse struct {
	Data struct {
		ErrorCode   int        `json:"error_code"`
		Description string     `json:"description"`
		Cursor      int        `json:"cursor"`
		HasMore     bool       `json:"has_more"`
		Users       []userinfo `json:"users"`
	} `json:"data"`
}

type userinfo struct {
	OpenID     string `json:"open_id"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Telephone  string `json:"telephone"`
	Wechat     string `json:"wechat"`
	City       string `json:"city"`
	Gender     int    `json:"gender"`
	Age        int    `json:"age"`
	LeadsLevel int    `json:"leads_level"`
	TagList    []struct {
		TagID   int    `json:"tag_id"`
		TagName string `json:"tag_name"`
	} `json:"tag_list"`
	IsFollow bool `json:"is_follow"`
}

// 获取意向用户列表
func LeadsUserList(accountToken, openID string, cursor, count int, startTime, endTime int64, leadsLevel, actionType int) (*LeadsUserListResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(leadsUserListUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
		"cursor":       {fmt.Sprintf("%d", cursor)},
		"count":        {fmt.Sprintf("%d", count)},
		"start_time":   {fmt.Sprintf("%d", startTime)},
		"end_time":     {fmt.Sprintf("%d", endTime)},
		"leads_level":  {fmt.Sprintf("%d", leadsLevel)},
		"action_type":  {fmt.Sprintf("%d", actionType)},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &LeadsUserListResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type LeadsUserDetailResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		userinfo
	} `json:"data"`
}

//获取意向用户详情
func LeadsUserDetail(accountToken, openID string, userID string) (*LeadsUserDetailResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(leadsUserDetailUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
		"user_id":      {userID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &LeadsUserDetailResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type LeadsUserActionListResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		HasMore     bool   `json:"has_more"`
		Cursor      string `json:"cursor"`
		List        []struct {
			ActionType   int    `json:"action_type"`
			UserID       string `json:"user_id"`
			ActionSource string `json:"action_source"`
			ActionFlag   int64  `json:"action_flag"`
			CreateTime   int64  `json:"create_time"`
		} `json:"list"`
	} `json:"data"`
}

//获取意向用户互动记录
func LeadsUserActionList(accountToken, openID string, cursor, count int, userID string, actionType int) (*LeadsUserActionListResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(leadsUserActionListUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
		"cursor":       {fmt.Sprintf("%d", cursor)},
		"count":        {fmt.Sprintf("%d", count)},
		"user_id":      {userID},
		"action_type":  {fmt.Sprintf("%d", actionType)},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &LeadsUserActionListResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type LeadsTagListResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		Cursor      int    `json:"cursor"`
		HasMore     bool   `json:"has_more"`
		List        []struct {
			TagID   int    `json:"tag_id"`
			TagName string `json:"tag_name"`
		} `json:"list"`
	} `json:"data"`
}

// 获取标签列表
func LeadsTagList(accountToken, openID string, cursor, count int64) (*LeadsTagListResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(leadsTagListUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
		"cursor":       {fmt.Sprintf("%d", cursor)},
		"count":        {fmt.Sprintf("%d", count)},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &LeadsTagListResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type LeadsTagUserListResponse struct {
	Data struct {
		ErrorCode   int      `json:"error_code"`
		Description string   `json:"description"`
		Cursor      int      `json:"cursor"`
		HasMore     bool     `json:"has_more"`
		List        []string `json:"list"`
	} `json:"data"`
}

// 获取标签用户列表
func LeadsTagUserList(accountToken, openID string, cursor, count, tagID int64) (*LeadsTagUserListResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(leadsTagUserListUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
		"cursor":       {fmt.Sprintf("%d", cursor)},
		"count":        {fmt.Sprintf("%d", count)},
		"tag_id":       {fmt.Sprintf("%d", tagID)},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &LeadsTagUserListResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type LeadsTagCreateResponse struct {
	Message string `json:"message"`
	Data    struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		TagID       int    `json:"tag_id"`
	} `json:"data"`
}

//创建标签
func LeadsTagCreate(accountToken, openID string, name string) (*LeadsTagCreateResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(leadsTagCreateUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	tag := map[string]string{"tag_name": name}
	data, err := json.Marshal(&tag)
	if err != nil {
		return nil, err
	}
	resp := &LeadsTagCreateResponse{}
	err = util.Post2Response(buf.String(), bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type LeadsTagUpdateResponse struct {
	Message string `json:"message"`
	Data    struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		TagID       int    `json:"tag_id"`
	} `json:"data"`
}

//更新标签
func LeadsTagUpdate(accountToken, openID string, name string, id int64) (*LeadsTagUpdateResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(leadsTagUpdateUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	tag := map[string]interface{}{"tag_name": name, "tag_id": id}
	data, err := json.Marshal(&tag)
	if err != nil {
		return nil, err
	}
	resp := &LeadsTagUpdateResponse{}
	err = util.Post2Response(buf.String(), bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type LeadsTagDeleteResponse struct {
	Message string `json:"message"`
	Data    struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
}

//删除标签
func LeadsTagDelete(accountToken, openID string, id int64) (*LeadsTagDeleteResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(leadsTagDeleteUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	tag := map[string]interface{}{"tag_id": id}
	data, err := json.Marshal(&tag)
	if err != nil {
		return nil, err
	}
	resp := &LeadsTagDeleteResponse{}
	err = util.Post2Response(buf.String(), bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type LeadsTagUserUpdateResponse struct {
	Message string `json:"message"`
	Data    struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
}

type LeadsTagUserUpdateRequest struct {
	TagID  int64  `json:"tag_id"`
	UserID string `json:"user_id"`
	Bind   *bool   `json:"bind"`
}

//删除标签
func LeadsTagUserUpdate(accountToken, openID string, tag LeadsTagUserUpdateRequest) (*LeadsTagUserUpdateResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(leadsTagUserUpdateUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	data, err := json.Marshal(&tag)
	if err != nil {
		return nil, err
	}
	resp := &LeadsTagUserUpdateResponse{}
	err = util.Post2Response(buf.String(), bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
