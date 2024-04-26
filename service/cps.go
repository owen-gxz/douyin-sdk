package service

import (
	"github.com/owen-gxz/douyin-sdk/life/commission"
)

func (s *Service) CpsCommission(req commission.CpsCommissionReq) (*commission.CpsCommissionResp, error) {
	return commission.CpsCommission(s.ClientToken(), req)
}

func (s *Service) CpsList(productID, page int64) (*commission.ListResp, error) {
	return commission.List(s.ClientToken(), productID, page, 20)
}
