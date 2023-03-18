package service

import (
	"github.com/owen-gxz/douyin-sdk/oauth"
)

// 获取openid的token
func (s *Service) GetAccessToken(openid string) (string, error) {
	return s.tokenService.GetToken(openid)
}

// 设置openid的token
func (s *Service) SetAccessToken(response oauth.TokenResponse) error {
	return s.tokenService.SaveToken(response)
}
