package user

import (
	"bytes"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
)

type FollowingResponse struct {
	Data struct {
		ErrorCode   int        `json:"error_code"`
		Description string     `json:"description"`
		Cursor      int        `json:"cursor"`
		HasMore     bool       `json:"has_more"`
		List        []userinfo `json:"list"`
	} `json:"data"`
}

// cursor 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
// count 每页数量
func GetFollowing(accountToken, openID string, cursor, count int) (*FollowingResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(followingUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
		"cursor":       {fmt.Sprintf("%d", cursor)},
		"count":        {fmt.Sprintf("%d", count)},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &FollowingResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
