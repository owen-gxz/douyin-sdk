package service

import (
	"github.com/owen-gxz/douyin-sdk/life/ledger"
)

func (s *Service) LedgerDetailed(certificateID ...string) (*ledger.LedgerDetailedResp, error) {
	return ledger.LedgerDetailed(s.ClientToken(), certificateID...)
}

func (s *Service) LedgerDetailedBySubFulfilID(subFulfilID string) (*ledger.LedgerDetailedBySubFulfilIDResp, error) {
	return ledger.LedgerDetailedBySubFulfilID(s.ClientToken(), subFulfilID)
}
