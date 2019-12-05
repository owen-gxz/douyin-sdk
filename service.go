package dy_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/owen-gxz/dy-sdk/oauth"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Service struct {
	*oauth.Config

	clientToken *AccessToken
	sync.Mutex    // accessToken读取锁
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

func NewService(conf *oauth.Config) *Service {
	s := &Service{
		Config: conf,
	}
	err := s.getClientToken()
	if err != nil {
		fmt.Errorf("getClientToken err: %s", err)
	}
	return s
}

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
	if strings.Contains(s.Endpoint.ClientTokenURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	result, err := http.DefaultClient.Get(buf.String())
	if err != nil {
		return err
	}
	response, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}
	resp := AccessToken{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return err
	}
	if resp.ErrorCode != 0 {
		return errors.New(fmt.Sprintf("error_code:%d ,msg: %s", resp.ErrorCode, resp.Description))
	}
	resp.ExpiresIn = time.Now().Unix() + resp.ExpiresIn - 3
	s.clientToken = &resp
	return nil
}
