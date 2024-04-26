package webhook

const (
	DouYinSignHeader = "X-Douyin-Signature"

	// 事件列表
	WebHookVerifyEvent      = "verify_webhook"
	CreateVideoEvent        = "create_video"
	AuthorizeEvent          = "authorize"
	UnAuthorizeEvent        = "unauthorize"
	ReceiveMsgEvent         = "receive_msg"
	EnterImEvent            = "enter_im"
	DialPhoneEvent          = "dial_phone"
	WebsiteContactEvent     = "website_contact"
	PersonalTabContactEvent = "personal_tab_contact"

	// 生活服务订单状态变更通知
	WebHookOrderEvent = "life_trade_order_notify"
	// 生活服务券状态变更通知
	WebHookCertificateEvent = "life_trade_certificate_notify"
	// 生活服务商品审核结果通知
	WebHookProductEvent = "life_product_audit"
	//代运营服务商订单通知
	WebHookPartnerOrderEvent = "life_partner_order_notify"
	//代运营服务商佣金变更通知
	WebHookPartnerCommissionEvent = "life_partner_commission_notify"
	//商品状态变更时通知服务商
	WebHookProductStatusEvent = "life_product_status_change"
	// 生活服务商品审核结果通知
	WebHookPartnerStatusEvent = "life_product_common_audit"

	//  订单信息
	CertificateActionRefundSuccess = "refund_success"
	CertificateActionVerifyCancel  = "verify_cancel"
	CertificateActionVerify        = "verify"
)

type WebhooksReq struct {
	Event      string      `json:"event"`
	ClientKey  string      `json:"client_key"`
	FromUserId string      `json:"from_user_id"`
	Content    interface{} `json:"content"`
	LogId      string      `json:"log_id"`
	EventId    string      `json:"event_id"`
}

type PartnerOrder struct {
	OrderId    string `json:"order_id"`
	MerchantId string `json:"merchant_id"`
	PartnerId  string `json:"partner_id"`
	Status     int    `json:"status"`
	EventTime  int64  `json:"event_time"`
}

type Order struct {
	Action  string `json:"action"`
	MsgTime int    `json:"msg_time"`
	Order   struct {
		OrderId        string `json:"order_id"`
		PayAmount      int    `json:"pay_amount"`
		OriginalAmount int    `json:"original_amount"`
		AccountId      string `json:"account_id"`
		CreateTime     int    `json:"create_time"`
		PayTime        int    `json:"pay_time"`
	} `json:"order"`
}

type Certificate struct {
	Action      string `json:"action"`
	MsgTime     int    `json:"msg_time"`
	Certificate struct {
		CertificateId string `json:"certificate_id"`
		OrderId       string `json:"order_id"`
		AccountId     string `json:"account_id"`
	} `json:"certificate"`
}
