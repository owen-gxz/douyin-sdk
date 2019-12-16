package user

import (
	"bytes"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
)

var (
	userInfoUrl  = fmt.Sprintf("%s/oauth/userinfo/", util.ServiceUrl)
	fansUrl      = fmt.Sprintf("%s/fans/list/", util.ServiceUrl)
	followingUrl = fmt.Sprintf("%s/following/list/", util.ServiceUrl)
	fansDataUrl  = fmt.Sprintf("%s/fans/data/", util.ServiceUrl)
)

type UserInfoResponse struct {
	Data    userinfo `json:"data"`
	Message string   `json:"message"`
}

type userinfo struct {
	ErrorCode   int    `json:"error_code,omitempy"`
	Description string `json:"description,omitempy"`
	OpenID      string `json:"open_id"`
	UnionID     string `json:"union_id"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Country     string `json:"country"`
	Gender      int    `json:"gender"`
}

func GetUserInfo(accountToken, openID string) (*UserInfoResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(userInfoUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &UserInfoResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
