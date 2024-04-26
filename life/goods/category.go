package goods

import (
	"bytes"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

const (
	categoryUrl = "https://open.douyin.com/goodlife/v1/goods/category/get/"
	templateUrl = "https://open.douyin.com/goodlife/v1/goods/template/get/"
)

type CategoryListResponse struct {
	BaseResp BaseResp `json:"BaseResp"`
	Data     Data     `json:"data"`
	Extra    Extra    `json:"extra"`
}
type BaseResp struct {
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
}

type CategoryTreeInfos struct {
	Enable       bool                `json:"enable"`
	IsLeaf       bool                `json:"is_leaf"`
	Level        int                 `json:"level"`
	Name         string              `json:"name"`
	ParentID     int                 `json:"parent_id"`
	SubTreeInfos []CategoryTreeInfos `json:"sub_tree_infos"`
	CategoryID   int                 `json:"category_id"`
}
type Data struct {
	ErrorCode         int                 `json:"error_code"`
	CategoryTreeInfos []CategoryTreeInfos `json:"category_tree_infos"`
	Description       string              `json:"description"`
}
type Extra struct {
	ErrorCode      int    `json:"error_code"`
	Description    string `json:"description"`
	SubErrorCode   int    `json:"sub_error_code"`
	SubDescription string `json:"sub_description"`
	Now            int    `json:"now"`
	Logid          string `json:"logid"`
}

// list 创建代运营订单
func CategoryList(accountToken string, accountID string) (*CategoryListResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(categoryUrl)
	v := url.Values{
		"account_id":          {accountID},
		"query_category_type": {"1"},
	}
	if strings.Contains(categoryUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &CategoryListResponse{}
	err := util.Get2Response2(buf.String(), accountToken, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
