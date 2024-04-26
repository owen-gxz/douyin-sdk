package service

import (
	"github.com/owen-gxz/douyin-sdk/life/goods"
)

func (s *Service) CategoryList(accountID string) (*goods.CategoryListResponse, error) {
	return goods.CategoryList(s.ClientToken(), accountID)
}

func (s *Service) Template(productType, categoryID string) (*goods.TemplateResp, error) {
	return goods.Template(s.ClientToken(), productType, categoryID)
}

func (s *Service) UpCreateGoods(req goods.CreateReq) (*goods.CreateResp, error) {
	return goods.UpData(s.ClientToken(), req)
}

func (s *Service) Goods(accountID, cursor, state, goodsCreatorType string) (*goods.ListSkuResp, error) {
	return goods.SkuList(s.ClientToken(), accountID, cursor, state, goodsCreatorType)
}

func (s *Service) Draft(accountID, cursor, state string) (*goods.ListSkuResp, error) {
	return goods.Draft(s.ClientToken(), accountID, cursor, state)
}

func (s *Service) GoodsInfo(accountID, productID string) (*goods.SkuInfoResp, error) {
	return goods.SkuInfo(s.ClientToken(), accountID, productID)
}

func (s *Service) Operate(req goods.OperateReq) (*goods.OperateResp, error) {
	return goods.Operate(s.ClientToken(), req)
}
