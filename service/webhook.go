package service

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

//create_video	aweme.share（默认都有）	用户使用开发者应用分享视频到抖音（携带分享id)
//authorize	无	用户授权给开发者应用
//receive_msg	im	接收到用户向授权给应用的抖音用户发送的私信具体类型
//enter_im	im	用户进入授权给应用的抖音用户的私信对话框
//dial_phone	im	用户拨打授权给应用的抖音用户的个人主页智能电话
//website_contact	im	用户在授权给应用的抖音用户的个人建站留资
//personal_tab_contact	im	用户在授权给应用的抖音用户的主页留资

const (
	DouYinSignHeader = "X-Douyin-Signature"

	// 事件列表
	WebHookVerifyEvent      = "verify_webhook"
	CreateVideoEvent        = "create_video"
	AuthorizeEvent          = "authorize"
	ReceiveMsgEvent         = "receive_msg"
	EnterImEvent            = "enter_im"
	DialPhoneEvent          = "dial_phone"
	WebsiteContactEvent     = "website_contact"
	PersonalTabContactEvent = "personal_tab_contact"
)

type HookReq []byte

type WebHookFunc func(req HookReq) []byte

func (s *Service) AddHandle(event string, fc WebHookFunc) {
	s.Lock()
	defer s.Unlock()
	s.handlers[event] = fc
}

func (s Service) CallBack(req *http.Request, resp http.ResponseWriter) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(data))
	if !s.SignValid(req.Header.Get(DouYinSignHeader), data) {
		fmt.Println("sign error")
		return
	}
	value := gjson.Get(string(data), "event")
	if value.String() == WebHookVerifyEvent {
		vwr := VerifyWebhookRequest{}
		err = json.Unmarshal(data, &vwr)
		if err != nil {
			fmt.Println(46, err.Error())
			resp.Write([]byte(err.Error()))
			return
		}
		respData, err := json.Marshal(&vwr.Content)
		if err != nil {
			fmt.Println(52, err.Error())
			resp.Write([]byte(err.Error()))
			return
		}
		resp.Write(respData)
		return
	}
	if s.handlers[value.String()] != nil {
		respData := s.handlers[value.String()](data)
		resp.Write(respData)
		return
	}
}

func (s Service) SignValid(sign string, data []byte) bool {
	sh1 := sha1.New()
	sh1.Write([]byte(fmt.Sprintf("%s%s", s.ClientSecret, string(data))))
	vsign := sh1.Sum(nil)
	signStr := fmt.Sprintf("%x", vsign)
	fmt.Println(sign, signStr)
	if signStr != sign {
		return false
	}
	return true
}

// url验证事件
type VerifyWebhookRequest struct {
	Event     string `json:"event"`
	ClientKey string `json:"client_key"`
	Content   struct {
		Challenge int `json:"challenge"`
	} `json:"content"`
}

type VerifyWebhookResponse struct {
	Challenge int `json:"challenge"`
}

// 接收到用户向授权给应用的抖音用户发送的私信具体类型
type ReceiveMsgRequest struct {
	Event      string `json:"event"`
	FromUserID string `json:"from_user_id"`
	ToUserID   string `json:"to_user_id"`
	ClientKey  string `json:"client_key"`
	Content    struct {
		MessageType string `json:"message_type,omitempty"`
		Text        string `json:"text,omitempty"`
		// 表情消息
		// MessageType    string `json:"message_type"`
		ResourceType   string `json:"resource_type,omitempty"`
		ResourceHeight int    `json:"resource_height,omitempty"`
		ResourceWidth  int    `json:"resource_width,omitempty"`
		ResourceURL    string `json:"resource_url,omitempty"`
		//卡片消息
		Title       string `json:"title"`
		IconURL     string `json:"icon_url"`
		Description string `json:"description"`
		LinkURL     string `json:"link_url"`
	} `json:"content"`
}

//用户使用开发者应用分享视频到抖音（携带分享id)
type CreateVideoRequest struct {
	Event      string `json:"event"`
	FromUserID string `json:"from_user_id"`
	ToUserID   string `json:"to_user_id"`
	ClientKey  string `json:"client_key"`
	Content    struct {
		MessageType string `json:"message_type,omitempty"`
		Text        string `json:"text,omitempty"`
		// 表情消息
		// MessageType    string `json:"message_type"`
		ResourceType   string `json:"resource_type,omitempty"`
		ResourceHeight int    `json:"resource_height,omitempty"`
		ResourceWidth  int    `json:"resource_width,omitempty"`
		ResourceURL    string `json:"resource_url,omitempty"`
		//卡片消息
		Title       string `json:"title"`
		IconURL     string `json:"icon_url"`
		Description string `json:"description"`
		LinkURL     string `json:"link_url"`
	} `json:"content"`
}

// 推送消息
func (data HookReq) WebHookVerifyEvent() *VerifyWebhookRequest {
	req := VerifyWebhookRequest{}
	err := json.Unmarshal(data, &req)
	if err != nil {
		return nil
	}
	return &req
}

// 表情  文字   卡片消息
func (data HookReq) ReceiveMsgEvent() *ReceiveMsgRequest {
	req := ReceiveMsgRequest{}
	err := json.Unmarshal(data, &req)
	if err != nil {
		return nil
	}
	return &req
}

//用户使用开发者应用分享视频到抖音（携带分享id)
func (data HookReq) CreateVideoRequest() *CreateVideoRequest {
	req := CreateVideoRequest{}
	err := json.Unmarshal(data, &req)
	if err != nil {
		return nil
	}
	return &req
}


