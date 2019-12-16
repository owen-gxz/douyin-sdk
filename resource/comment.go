package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
	"strings"
)

var (
	videoCommentListUrl      = fmt.Sprintf("%s/video/comment/list/", util.ServiceUrl)
	videoCommentReplyListUrl = fmt.Sprintf("%s/video/comment/reply/list/", util.ServiceUrl)
	videoCommentReplyUrl     = fmt.Sprintf("%s/video/comment/reply/", util.ServiceUrl)
	videoCommentTopUrl       = fmt.Sprintf("%s/video/comment/top/", util.ServiceUrl)
)

type VideoCommentListResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		Cursor      int    `json:"cursor"`
		HasMore     bool   `json:"has_more"`
		List        []struct {
			CommentID         string `json:"comment_id"`
			CommentUserID     string `json:"comment_user_id"`
			Content           string `json:"content"`
			CreateTime        int    `json:"create_time"`
			Top               bool   `json:"top"`
			DiggCount         int    `json:"digg_count"`
			ReplyCommentTotal int    `json:"reply_comment_total"`
		} `json:"list"`
	} `json:"data"`
}

func GetVideoCommentList(accountToken, openID string, cursor, count int, itemID string) (*VideoCommentListResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(videoCommentListUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
		"cursor":       {fmt.Sprintf("%d", cursor)},
		"count":        {fmt.Sprintf("%d", count)},
		"item_id":      {itemID},
	}
	if strings.Contains(videoCommentListUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &VideoCommentListResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type VideoCommentReplyListResponse struct {
	VideoCommentListResponse
}

func GetVideoCommentReplyList(accountToken, openID string, cursor, count int, itemID, commentID string) (*VideoCommentReplyListResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(videoCommentReplyListUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
		"cursor":       {fmt.Sprintf("%d", cursor)},
		"count":        {fmt.Sprintf("%d", count)},
		"item_id":      {itemID},
		"comment_id":   {commentID},
	}
	if strings.Contains(videoCommentReplyListUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &VideoCommentReplyListResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ComentReq struct {
	CommentID string `json:"comment_id"`
	ItemID    string `json:"item_id"`
	Content   string `json:"content"`
}
type VideoCommentReplyResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
}

// 回复评论
func VideoCommentReply(accountToken, openID string, reply ComentReq) (*VideoCommentReplyResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(videoCommentReplyUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	data, err := json.Marshal(&reply)
	if err != nil {
		return nil, err
	}
	resp := &VideoCommentReplyResponse{}
	err = util.Post2Response(buf.String(), bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type VideoCommentTopResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
}

// 置顶视频评论(企业号)
func VideoCommentTop(accountToken, openID string, reply ComentReq) (*VideoCommentTopResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(videoCommentTopUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	data, err := json.Marshal(&reply)
	if err != nil {
		return nil, err
	}
	resp := &VideoCommentTopResponse{}
	err = util.Post2Response(buf.String(), bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
