package douyin_model

const (
	DouyinEventVerity                 = "verify_webhook"
	DouyinEventLifePartnerOrderNotify = "life_partner_order_notify"
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
