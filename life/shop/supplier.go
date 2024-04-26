package shop

import (
	"bytes"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

var (
	supplierQueryUrl = "https://open.douyin.com/poi/supplier/query/"
)

func Supplier(accountToken string, supplierExtID string) (*QueryResp, error) {
	var buf bytes.Buffer
	buf.WriteString(supplierQueryUrl)
	v := url.Values{
		"supplier_ext_id": {supplierExtID},
	}
	if strings.Contains(supplierQueryUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &QueryResp{}
	err := util.Get2Response2(buf.String(), accountToken, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
