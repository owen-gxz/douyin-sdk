package douyin_model

//const (
//	DouyinEventVerity                 = "verify_webhook"
//	DouyinEventLifePartnerOrderNotify = "life_partner_order_notify"
//)
//
//type WebhooksReq struct {
//	Event      string      `json:"event"`
//	ClientKey  string      `json:"client_key"`
//	FromUserId string      `json:"from_user_id"`
//	Content    interface{} `json:"content"`
//	LogId      string      `json:"log_id"`
//	EventId    string      `json:"event_id"`
//}
//
//type PartnerOrder struct {
//	OrderId    string `json:"order_id"`
//	MerchantId string `json:"merchant_id"`
//	PartnerId  string `json:"partner_id"`
//	Status     int    `json:"status"`
//	EventTime  int64  `json:"event_time"`
//}

//1。 服务商发起代运营请求，webhook请求
//{"event":"life_partner_order_notify","client_key":"awfu87gyyi5ktef3","from_user_id":"","content":"{\"order_id\":\"7212142857733638179\",\"merchant_id\":\"7126518681177163816\",\"partner_id\":\"7207450161063135284\",\"status\":102,\"event_time\":1679210010000}","log_id":"021679210010614fdbddc01000b04950000000000000039e59b04"}

//2。 代运营同意webhook
//{"event":"life_partner_order_notify","client_key":"awfu87gyyi5ktef3","from_user_id":"","content":"{\"order_id\":\"7212142857733638179\",\"merchant_id\":\"7126518681177163816\",\"partner_id\":\"7207450161063135284\",\"status\":301,\"event_time\":1679210231000}","log_id":"021679210231184fdbddc01000b0495000000000000003931bf5c"}

//3. 团购商品支付成功
// {"event":"life_trade_order_notify","client_key":"awfu87gyyi5ktef3","from_user_id":"_000vHbSsXCxQE6wOF2T0qv3byKqefdHisNw","content":"{\"action\":\"pay_success\",\"msg_time\":1679304435,\"order\":{\"order_id\":\"8000007678483824642\",\"account_id\":\"7126518681177163816\",\"pay_amount\":5300,\"original_amount\":5400,\"create_time\":1679304410,\"pay_time\":1679304433}}","log_id":"021679304434213fdbddc01000b050300000000000000437669d6"}

//4。团购商品核销
//{"event":"life_trade_certificate_notify","client_key":"awfu87gyyi5ktef3","from_user_id":"_000vHbSsXCxQE6wOF2T0qv3byKqefdHisNw","content":"{\"action\":\"verify\",\"msg_time\":1679914989,\"certificate\":{\"certificate_id\":\"7212557509230723106\",\"order_id\":\"8000007678483824642\",\"account_id\":\"7126518681177163816\"}}","log_id":"20230327190308EA6D9C39D1F3650CAC89"}
