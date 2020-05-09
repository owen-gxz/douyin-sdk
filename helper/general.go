package helper

import (
	"errors"
	"github.com/owen-gxz/douyin-sdk/oauth"
	"sync"
	"time"
)

var (
	OpenKeyEmpty = errors.New("get account token, key is empty")
	OpenIDEmpty = errors.New("save account token, open id is empty")
)

// 内存存储
type General struct {
	sync.RWMutex
	Accesses map[string]oauth.TokenResponse
	config   *oauth.Config
}

func NewGeneral(cfg *oauth.Config) *General {
	return &General{Accesses: make(map[string]oauth.TokenResponse), config: cfg}
}

func (g *General) GetToken(openid string) (token string, err error) {
	g.Lock()
	defer g.Unlock()
	t, ok := g.Accesses[getOpenKey(openid)]
	if !ok {
		return "", OpenKeyEmpty
	}
	if t.Data.ExpiresIn < time.Now().Unix() {
		rToken, err := g.config.RefreshToken(t.Data.RefreshToken)
		if err != nil {
			return "", err
		}
		if rToken.Data.ErrorCode != 0 {
			return "", errors.New(rToken.Data.Description)
		}
		err = g.SaveToken(rToken.TokenResponse)
		if err != nil {
			return "", err
		}
		return rToken.Data.AccessToken, nil
	}
	return t.Data.AccessToken, nil
}

func (g *General) SaveToken(response oauth.TokenResponse) error {
	if response.Data.OpenID=="" {
		return OpenIDEmpty
	}
	g.Lock()
	defer g.Unlock()
	g.Accesses[response.Data.OpenID] = response
	return nil
}

func getOpenKey(openid string) string {
	return getRedisOpenKey(openid)
}
