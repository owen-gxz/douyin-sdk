package goods

import (
	"bytes"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

type TemplateResp struct {
	Data struct {
		ErrorCode    int `json:"error_code"`
		ProductAttrs []struct {
			Name       string `json:"name"`
			ValueDemo  string `json:"value_demo"`
			ValueType  string `json:"value_type"`
			Desc       string `json:"desc"`
			IsMulti    bool   `json:"is_multi"`
			IsRequired bool   `json:"is_required"`
			Key        string `json:"key"`
		} `json:"product_attrs"`
		SkuAttrs []struct {
			IsRequired bool   `json:"is_required"`
			Key        string `json:"key"`
			Name       string `json:"name"`
			ValueDemo  string `json:"value_demo"`
			ValueType  string `json:"value_type"`
			Desc       string `json:"desc"`
			IsMulti    bool   `json:"is_multi"`
		} `json:"sku_attrs"`
		SpuAttrs      interface{} `json:"spu_attrs"`
		CalendarAttrs interface{} `json:"calendar_attrs"`
		Description   string      `json:"description"`
	} `json:"data"`
	Extra struct {
		ErrorCode      int    `json:"error_code"`
		Description    string `json:"description"`
		SubErrorCode   int    `json:"sub_error_code"`
		SubDescription string `json:"sub_description"`
		Logid          string `json:"logid"`
		Now            int    `json:"now"`
	} `json:"extra"`
}

// list 创建代运营订单
func Template(accountToken string, productType, categoryID string) (*TemplateResp, error) {
	var buf bytes.Buffer
	buf.WriteString(templateUrl)
	v := url.Values{
		"product_type": {productType},
		"category_id":  {categoryID},
	}
	if strings.Contains(templateUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &TemplateResp{}
	err := util.Get2Response2(buf.String(), accountToken, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
