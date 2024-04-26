package goods

import (
	"bytes"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

// list 创建代运营订单
func Draft(accountToken string, accountID, cursor, status string) (*ListSkuResp, error) {
	var buf bytes.Buffer
	buf.WriteString(draftGoodsUrl)
	v := url.Values{
		"cursor": {cursor},
		"count":  {"15"},
	}
	if accountID != "" {
		v.Add("account_id", accountID)
	}
	if status != "" {
		v.Add("status", status)
	}
	if strings.Contains(draftGoodsUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &ListSkuResp{}
	err := util.Get2Response2(buf.String(), accountToken, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
