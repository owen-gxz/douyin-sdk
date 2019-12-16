package oauth

import (
	"bytes"
	"errors"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
	"time"
)

type RefreshTokenResponse struct {
	TokenResponse
}
type TokenResponse struct {
	Message string `json:"message"`
	Data    struct {
		ErrorCode    int    `json:"error_code"`
		Description  string `json:"description"`
		AccessToken  string `json:"access_token"`
		ExpiresIn    int64  `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		OpenID       string `json:"open_id"`
		Scope        string `json:"scope"`
	} `json:"data"`
}

var (
	CodeIsNullError         = errors.New("code is null")
	RefreshTokenIsNullError = errors.New("refresh token is null")
)

const (
	grantTypeAuthorizationCode = "authorization_code"
	grantTypeRefreshToken      = "refresh_token"
	responseTypeCode           = "code"
)

func (c *Config) Token(code string) (*TokenResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(c.Endpoint.TokenURL)
	//oauth2.Config{}
	if code == "" {
		return nil, CodeIsNullError
	}
	v := url.Values{
		"grant_type":    {grantTypeAuthorizationCode},
		"client_key":    {c.ClientKey},
		"client_secret": {c.ClientSecret},
		"code":          {code},
	}
	if strings.Contains(c.Endpoint.TokenURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &TokenResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	resp.Data.ExpiresIn = time.Now().Unix() + resp.Data.ExpiresIn
	return resp, nil
}

func (c *Config) RefreshToken(refreshToken string) (*RefreshTokenResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(c.Endpoint.RefreshTokenURL)
	//oauth2.Config{}
	if refreshToken == "" {
		return nil, RefreshTokenIsNullError
	}
	v := url.Values{
		"grant_type":    {grantTypeRefreshToken},
		"client_key":    {c.ClientKey},
		"refresh_token": {refreshToken},
	}
	if strings.Contains(c.Endpoint.RefreshTokenURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &RefreshTokenResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	resp.Data.ExpiresIn = time.Now().Unix() + resp.Data.ExpiresIn
	return resp, nil
}
