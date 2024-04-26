package service

import (
	"errors"
	"github.com/owen-gxz/douyin-sdk/life/shop"
)

// shop
func (s *Service) ShopInfo(accountID string) ([]shop.PoiModel, error) {
	qr, err := shop.Query(s.ClientToken(), accountID, "1", "100")
	if err != nil {
		return nil, err
	}
	if qr.Data.ErrorCode != 0 {
		return nil, errors.New(qr.Data.Description)
	}
	if len(qr.Data.Pois) < 1 {
		return nil, errors.New("地址错误")
	}
	return qr.Data.Pois, nil
}

func (s *Service) Match(req shop.MatchReq) (*shop.MatchResp, error) {
	qr, err := shop.Match(s.ClientToken(), req)
	if err != nil {
		return nil, err
	}
	if qr.Data.ErrorCode != 0 {
		return nil, errors.New(qr.Data.Description)
	}
	return qr, nil
}

func (s *Service) Supplier(supplierExtID string) (*shop.QueryResp, error) {
	qr, err := shop.Supplier(s.ClientToken(), supplierExtID)
	if err != nil {
		return nil, err
	}
	if qr.Data.ErrorCode != 0 {
		return nil, errors.New(qr.Data.Description)
	}
	return qr, nil
}
