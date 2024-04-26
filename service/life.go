package service

import (
	"github.com/owen-gxz/douyin-sdk/life/order"
	"github.com/owen-gxz/douyin-sdk/life/partner"
)

// Partner
func (s *Service) CreatePartner(req partner.CreateReq) (*partner.CreateResp, error) {
	return partner.Create(s.ClientToken(), req)
}

func (s *Service) PartnerInfo(orderID string) (*partner.InfoResp, error) {
	return partner.Info(s.ClientToken(), orderID)
}

func (s *Service) OrderInfo(accountID, orderID, pageNum string) (*order.InfoResp, error) {
	return order.Info(s.ClientToken(), accountID, orderID, pageNum)
}

func (s *Service) Commission(orderID, page string) (*partner.CommissionsResp, error) {
	return partner.Commissions(s.ClientToken(), orderID, page)
}

func (s *Service) CreateCommission(req partner.CreateCommissionReq) (*partner.CreateCommissionResp, error) {
	return partner.CreateCommission(s.ClientToken(), req)
}
