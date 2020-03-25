package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/owen-gxz/douyin-sdk/oauth"
	"time"
)

const (
	DouYinOpenKey = "DOUYIN::OPEN::KEY::%s"
)

type RedisAccountTokenService struct {
	redis  *redis.Client
	config *oauth.Config
}

func NewRedisService(redisClient *redis.Client, cfg *oauth.Config) *RedisAccountTokenService {
	return &RedisAccountTokenService{redisClient, cfg}
}

func (s *RedisAccountTokenService) SaveToken(response oauth.TokenResponse) error {
	data, err := json.Marshal(&response)
	if err != nil {
		return err
	}
	return s.redis.Set(getRedisOpenKey(response.Data.OpenID), string(data), 0).Err()
}

func (s *RedisAccountTokenService) GetToken(openid string) (token string, err error) {
	tokenData, err := s.redis.Get(getRedisOpenKey(openid)).Result()
	if err != nil {
		return "", err
	}
	t := oauth.TokenResponse{}
	err = json.Unmarshal([]byte(tokenData), &t)
	if err != nil {
		return "", err
	}
	if t.Data.ExpiresIn < time.Now().Unix() {
		rToken, err := s.config.RefreshToken(t.Data.RefreshToken)
		if err != nil {
			return "", err
		}
		if rToken.Data.ErrorCode != 0 {
			return "", errors.New(rToken.Data.Description)
		}
		err = s.SaveToken(rToken.TokenResponse)
		if err != nil {
			return "", err
		}
		return rToken.Data.AccessToken, nil
	}
	return t.Data.AccessToken, nil
}

func getRedisOpenKey(openid string) string {
	return fmt.Sprintf(DouYinOpenKey, openid)
}
