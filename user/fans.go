package user

import (
	"bytes"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

type FansResponse struct {
	Data struct {
		ErrorCode   int        `json:"error_code"`
		Description string     `json:"description"`
		Cursor      int        `json:"cursor"`
		HasMore     bool       `json:"has_more"`
		List        []userinfo `json:"list"`
	}
}

// cursor 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
// count 每页数量
func GetFans(accountToken, openID string, cursor, count int) (*FansResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(fansUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
		"cursor":       {fmt.Sprintf("%d", cursor)},
		"count":        {fmt.Sprintf("%d", count)},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &FansResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type FansDataResponse struct {
	ErrorCode   int     `json:"error_code"`
	Description string  `json:"description"`
	FansData    fansata `json:"fans_data"`
}
type fansata struct {
	AllFansNum                int64           `json:"all_fans_num"`
	GenderDistributions       []distributions `json:"gender_distributions"`
	AgeDistributions          []distributions `json:"age_distributions"`
	GeographicalDistributions struct {
	} `json:"geographical_distributions"`
	ActiveDaysDistributions []distributions `json:"active_days_distributions"`
	DeviceDistributions     []distributions `json:"device_distributions"`
	InterestDistributions   []distributions `json:"interest_distributions"`
	FlowContributions       []struct {
		Flow    string `json:"flow"`
		FansSum string `json:"fans_sum"`
		AllSum  string `json:"all_sum"`
	} `json:"flow_contributions"`
}

type distributions struct {
	Item  string `json:"item"`  //分布的种类
	Value int64  `json:"value"` //数量
}

//该接口用于查询用户的粉丝数据，如性别分布，年龄分布，地域分布等。
//注：用户首次授权应用后，需要间隔2天才会产生全部的数据。
func GetFansData(accountToken, openID string) (*FansDataResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(fansDataUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	if strings.Contains(userInfoUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &FansDataResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
