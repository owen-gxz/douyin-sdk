package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/helper"
	"github.com/owen-gxz/douyin-sdk/oauth"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"sync"
	"time"
)

type Service struct {
	*oauth.Config

	clientToken *AccessToken
	sync.Mutex  // accessToken读取锁

	handlers map[string]WebHookFunc

	// Access Token Server
	tokenService AccessTokenServer
}

const (
	grantTypeClientCredential = "client_credential"
)

type AccessToken struct {
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func NewService(conf *oauth.Config, tokenService AccessTokenServer) *Service {
	s := &Service{
		Config: conf,
	}
	if s.handlers == nil {
		s.handlers = make(map[string]WebHookFunc)
	}
	if tokenService == nil {
		tokenService = helper.NewGeneral(conf)
	}
	s.tokenService = tokenService
	err := s.getClientToken()
	if err != nil {
		fmt.Errorf("getClientToken err: %s", err)
	}
	return s
}

// 抖音service token
func (s Service) ClientToken() string {
	s.Lock()
	defer s.Unlock()
	var err error
	if s.clientToken == nil || s.clientToken.ExpiresIn < time.Now().Unix() {
		for i := 0; i < 3; i++ {
			err = s.getClientToken()
			if err == nil {
				break
			}
			fmt.Errorf("getClientToken[%v] %v", s.ClientKey, err)
			time.Sleep(time.Second)
		}
		if err != nil {
			return ""
		}
	}
	return s.clientToken.AccessToken
}

func (s Service) getClientToken() error {
	var buf bytes.Buffer
	buf.WriteString(s.Endpoint.ClientTokenURL)
	v := url.Values{
		"grant_type":    {grantTypeClientCredential},
		"client_key":    {s.ClientKey},
		"client_secret": {s.ClientSecret},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &AccessToken{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return err
	}
	if resp.ErrorCode != 0 {
		return errors.New(fmt.Sprintf("error_code:%d ,msg: %s", resp.ErrorCode, resp.Description))
	}
	resp.ExpiresIn = time.Now().Unix() + resp.ExpiresIn - 3
	s.clientToken = resp
	return nil
}

type AccessTokenServer interface {
	GetToken(openid string) (token string, err error)
	SaveToken(response oauth.TokenResponse) error
}
