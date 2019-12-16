package im

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/util"
	"net/url"
)

var (
	sendMessageUrl = fmt.Sprintf("%s/im/message/send/", util.ServiceUrl)
)

const (
	MessageTypeText  = "text"
	MessageTypeImage = "image"
)

type MessageReq struct {
	ToUserID    string `json:"to_user_id"`
	MessageType string `json:"message_type"`
	Content     string `json:"content"`
}

type SendMessageResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
}

// 发送私信消息
func SendMessage(accountToken, openID string, reply MessageReq) (*SendMessageResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(sendMessageUrl)
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
	resp := &SendMessageResponse{}
	err = util.Post2Response(buf.String(), bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
