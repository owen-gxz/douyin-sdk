package goods

import (
	"bytes"
	"encoding/json"
	"github.com/owen-gxz/douyin-sdk/util"
)

const (
	createGoodsUrl  = "https://open.douyin.com/goodlife/v1/goods/product/save/"
	draftGoodsUrl   = "https://open.douyin.com/goodlife/v1/goods/product/draft/query/"
	operateGoodsUrl = "https://open.douyin.com/goodlife/v1/goods/product/operate/"
	listSkuUrl      = "https://open.douyin.com/goodlife/v1/goods/product/online/query/"
	SkuInfoUrl      = "https://open.douyin.com//goodlife/v1/goods/product/online/get/"
	CreateImageIL   = "image_list"
	CreateImageEIL  = "environment_image_list"
	CreateImageDIL  = "dishes_image_list"
	CreateImageDEIL = "detail_image_list"
)

type CreateReq struct {
	AccountID string  `json:"account_id,omitempty"`
	Product   Product `json:"product,omitempty"`
	Sku       Sku     `json:"sku,omitempty"`
}

type AttrKeyValueMap struct {
	LimitUseRule              *LimitUseRuleStruct         `json:"limit_use_rule"`                         //是否限制、每人单次消费最多使用代金劵张数
	Appointment               *AppointmentStruct          `json:"appointment,omitempty"`                  // 预约信息
	OrderSettleRule           *OrderSettleRule            `json:"order_settle_rule,omitempty"`            //结算规则
	LimitGender               string                      `json:"limit_gender,omitempty"`                 // 是否限制性别   0-不限制 1-限制男性 2-限制女性
	LimitHairLength           string                      `json:"limit_hair_length,omitempty"`            // 是否限制长短发   0-不限制 1-限制长发 2-限制短发
	SuitableGroup             string                      `json:"SuitableGroup,omitempty"`                // 适用人群
	AutoRenew                 string                      `json:"auto_renew,omitempty"`                   // 是否开启自动延期	"true"/"false"
	ContainsInsurance         string                      `json:"contains_insurance,omitempty"`           // 游玩景点票卷，园内套票	"true"/"false"
	ConsumptionTimes          string                      `json:"consumption_times,omitempty"`            // 可用次数
	OriginalVipCanExperience  string                      `json:"original_vip_can_experience,omitempty"`  // 商家原会员是否可以体验1=可以   0=不可以"
	AttachItemDetail          *AttachItemDetailStruct     `json:"attach_item_detail,omitempty"`           // 附赠项目明细"
	BringOutMeal              string                      `json:"bring_out_meal,omitempty"`               // 是否可以外带餐食
	Description               []string                    `json:"Description"`                            // 商品描述
	CanNoUseDate              *CanNoUseDateStruct         `json:"can_no_use_date,omitempty"`              // 不可使用日期    消费提示里注明的不可使用日期，可以天、星期和节日
	CustomerReservedInfo      *CustomerReservedInfoStruct `json:"customer_reserved_info,omitempty"`       // 是否留资
	DescriptionRichText       []*NoteStruct               `json:"description_rich_text,omitempty"`        // 其他说明信息
	DetailImageList           []*ImageStruct              `json:"detail_image_list,omitempty"`            // 长图
	DishesImageList           []*ImageStruct              `json:"dishes_image_list,omitempty"`            // 菜品图	图片比例：375:280
	EnvironmentImageList      []*ImageStruct              `json:"environment_image_list,omitempty"`       // 环境图	图片比例：375:280
	ImageList                 []*ImageStruct              `json:"image_list,omitempty"`                   // 封面图	图片比例：375:280
	FreePack                  string                      `json:"free_pack,omitempty"`                    // 是否可以打包	消费提示：做展示使用
	Notification              []*NotificationStruct       `json:"Notification,omitempty"`                 //使用规则
	PrivateRoom               string                      `json:"private_room,omitempty"`                 // 是否可以使用包间	消费提示：做展示使用 "true"/"false"
	RealNameInfo              *RealNameInfoStruct         `json:"real_name_info,omitempty"`               // 实名信息
	RecommendWord             string                      `json:"RecommendWord,omitempty"`                // 推荐语
	RecPersonNum              string                      `json:"rec_person_num,omitempty"`               // 建议使用人数
	RecPersonNumMax           string                      `json:"rec_person_num_max,omitempty"`           // 最多使用人数
	RefundPolicy              string                      `json:"RefundPolicy,omitempty"`                 // 退款政策	 "1-允许退款 2-不可退款 3-有条件退",
	RefundNeedMerchantConfirm string                      `json:"refund_need_merchant_confirm,omitempty"` // 退款是否需商家审核
	ReleaseSource             string                      `json:"release_source,omitempty"`               // 商品发布渠道	 "MERCHANT = 1 // 商家; BD = 2 // BD; FACILITATOR = 3 // 服务商;",
	ShowChannel               string                      `json:"show_channel,omitempty"`                 // show_channel	 投放渠道	 "1-不限制 2-仅直播间可见",
	SortWeight                string                      `json:"SortWeight,omitempty"`                   // SortWeight	 排序权重
	SuperimposedDiscounts     string                      `json:"superimposed_discounts,omitempty"`       // 可以享受店内其他优惠	消费提示：做展示使用   true/false
	UseDate                   *UseDateStruct              `json:"use_date,omitempty"`                     // 使用日期	券码的可以核销日期，履约核销强依赖
	UseTime                   *UseTimeStruct              `json:"use_time,omitempty"`                     // 使用时间	 用户可以消费的时间
	CodeSourceType            string                      `json:"code_source_type,omitempty"`             // 券码生成方式 使用1
	Commodity                 []*ItemGroupStruct          `json:"commodity,omitempty"`                    // 菜品搭配
	LimitRule                 *LimitRuleStruct            `json:"limit_rule"`
	MarketPrice               string                      `json:"market_price,omitempty"` // 市场价	即菜品搭配里的总价
	SettleType                string                      `json:"settle_type"`            // 收款方式	" ""1-总店结算 2-分店结算"", 总店结算：即商品的结算资金统一结算到商家（不是开发者）的收款账户。 分店结算：按核销POI将资金结算到对应的POI的收款账户，如果POI没有设置收款账户，会将对应的POI的结算资金打款到总店账户；"
	UseType                   string                      `json:"use_type,omitempty"`     // 团购使用方式	 "1-到店核销",默认值
	SubTitle                  string                      `json:"SubTitle,omitempty"`     // 副标题	过期退；随时退；x日内可退；免预约；提前x日预约；多个副标题以|（英文半角）分隔，不要有空格
	AccountName               string                      `json:"account_name,omitempty"` //
	ProductName               string                      `json:"product_name,omitempty"`
	SoldStartTime             string                      `json:"sold_start_time,omitempty"`
	SoldEndTime               string                      `json:"sold_end_time,omitempty"`
}

type OrderSettleRule struct {
	//BY_ORDER = 1, // 整单结算
	//BY_TIMES = 2, // 按使用次数结算
	OrderSettleType int `json:"order_settle_type"`
}

// 附赠项目
type AttachItemDetailStruct struct {
	AttachItemList []AttachItem `json:"attach_item_list"` // 附赠项目明细列表
}
type AttachItem struct {
	ItemName       string // 项目名称
	ItemCount      int    // 项目份数
	ItemTotalPrice int    // 总价值
}

type LimitUseRuleStruct struct {
	IsLimitUse       bool `json:"is_limit_use"`        // 是否限用
	UseNumPerConsume int  `json:"use_num_per_consume"` // 每人单次消费最多使用代金劵张数
}

type LimitRuleStruct struct {
	IsLimit     bool `json:"is_limit"`                // 是否限购
	TotalBuyNum int  `json:"total_buy_num,omitempty"` // 每人最多购买X张
}

type ItemGroupStruct struct {
	GroupName   string     `json:"group_name,omitempty"`   // 商品组名
	TotalCount  int        `json:"total_count,omitempty"`  // 总数
	OptionCount int        `json:"option_count,omitempty"` // 选几
	ItemList    []ItemList `json:"item_list,omitempty"`    // 菜品
}
type ItemList struct {
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"` //分
	Count int    `json:"count,omitempty"`
	Unit  string `json:"unit,omitempty"`
}

type UseTimeStruct struct {
	UseTimeType    int              `json:"use_time_type,omitempty"`    // 1全天可用，2仅指定时间可用
	TimePeriodList []TimePeriodList `json:"time_period_list,omitempty"` // 时间段
}
type TimePeriodList struct {
	UseStartTime     string `json:"use_start_time,omitempty"`
	UseEndTime       string `json:"use_end_time,omitempty"`
	EndTimeIsNextDay bool   `json:"end_time_is_next_day"`
}

type UseDateStruct struct {
	UseDateType  int    `json:"use_date_type,omitempty"`  // 1:指定日期 2:指定天数；(暂时只做2)
	DayDuration  int    `json:"day_duration,omitempty"`   // 购买后X天有效，use_date_type=2时有效
	UseStartDate string `json:"use_start_date,omitempty"` // yyyy-MM-dd  开始日期，use_date_type=1时有效
	UseEndDate   string `json:"use_end_date,omitempty"`   // yyyy-MM-dd  结束日期，use_date_type=1时有效
}

type RealNameInfoStruct struct {
	Enable bool `json:"enable"`
	Scene  int  `json:"scene,omitempty"` // 3   // 出行人姓名与手机号码 2   // 仅填写一位游客信息 1   // 每张门票都要填写用户信息
}

type ImageStruct struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"` // 传base64
}

type NotificationStruct struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type NoteStruct struct {
	NoteType int    `json:"note_type,omitempty"` // 暂时填写1.文本
	Content  string `json:"content,omitempty"`   //内容
}

type CustomerReservedInfoStruct struct {
	Allow         bool `json:"allow"`                     //是否留资
	AllowTel      bool `json:"allow_tel,omitempty"`       //可以留电话
	AllowName     bool `json:"allow_name,omitempty"`      // 可以留姓名
	AllowIdentity bool `json:"allow_identity,omitempty"`  // 可以留身份证
	RequireForTel bool `json:"require_for_tel,omitempty"` //手机号是否必传 【默认非必传】
	NeedForAll    bool `json:"need_for_all,omitempty"`    // 是否每张券都需要留资
}

type CanNoUseDateStruct struct {
	Enable     bool     `json:"enable"`                 // 开关，启用需要为true
	DaysOfWeek []int    `json:"days_of_week,omitempty"` //指定周几不可用 1-7
	Holidays   []int    `json:"holidays,omitempty"`     //指定节假日不可用 1。元旦 2：春节 3：清明  4：劳动节 5：端午节 6：中秋节 7：国庆节 8：情人节  9：圣诞节
	DateList   []string `json:"date_list,omitempty"`    //yy-MM-dd 指定日期，不可用
}

type AppointmentStruct struct {
	NeedAppointment bool `json:"need_appointment"`          //是否需要预约
	AheadTimeType   int  `json:"ahead_time_type,omitempty"` //提前预约时间类型，  need_appointment = true时该字段必填 DAY  = 1  HOUR = 2  MINUTE = 3
	AheadDayNum     int  `json:"ahead_day_num,omitempty"`   //需要提前X天电话预约

	ExternalLink            string `json:"external_link,omitempty"`              //第三方提供预约入口，需要过机审。
	OrderAppointmentTimeUrl string `json:"order_appointment_time_url,omitempty"` //第三方提供的查看已预约订单入口，需要过机审
	AheadHourNum            int    `json:"ahead_hour_num,omitempty"`             //需要提前X小时电话预约
	AheadMinuteNum          int    `json:"ahead_minute_num,omitempty"`           //需要提前X分钟电话预约
}

type Pois struct {
	PoiID         string `json:"poi_id,omitempty"`
	SupplierID    int64  `json:"supplier_id,omitempty"`
	SupplierExtID string `json:"supplier_ext_id,omitempty"`
}

type Product struct {
	AccountName      string            `json:"account_name,omitempty"`
	AttrKeyValueMap  map[string]string `json:"attr_key_value_map,omitempty"`
	VK               *AttrKeyValueMap  `json:"vk,omitempty"`
	BizLine          int               `json:"biz_line,omitempty"`
	CategoryID       int               `json:"category_id,omitempty"`
	CategoryFullName string            `json:"category_full_name,omitempty"`
	OutID            string            `json:"out_id,omitempty"`
	Pois             []Pois            `json:"pois,omitempty"`
	ProductName      string            `json:"product_name,omitempty"`
	ProductType      int               `json:"product_type,omitempty"`
	SoldEndTime      int               `json:"sold_end_time,omitempty"`
	SoldStartTime    int               `json:"sold_start_time,omitempty"`
	Telephone        []string          `json:"telephone,omitempty"`
}

type Stock struct {
	LimitType int `json:"limit_type,omitempty"`
	StockQty  int `json:"stock_qty,omitempty"`
}

type Sku struct {
	ActualAmount    int               `json:"actual_amount,omitempty"`
	AttrKeyValueMap map[string]string `json:"attr_key_value_map,omitempty"`
	VK              *AttrKeyValueMap  `json:"vk,omitempty"`
	OriginAmount    int               `json:"origin_amount,omitempty"`
	SkuName         string            `json:"sku_name,omitempty"`
	Status          int               `json:"status,omitempty"`
	Stock           Stock             `json:"stock,omitempty"`
}

type CreateResp struct {
	Data struct {
		ErrorCode   int    `json:"error_code,omitempty"`
		Description string `json:"description,omitempty"`
		ProductId   string `json:"product_id,omitempty"`
	} `json:"data,omitempty"`
	Extra struct {
		ErrorCode      int    `json:"error_code,omitempty"`
		Description    string `json:"description,omitempty"`
		SubErrorCode   int    `json:"sub_error_code,omitempty"`
		SubDescription string `json:"sub_description,omitempty"`
		Logid          string `json:"logid,omitempty"`
		Now            int    `json:"now,omitempty"`
	} `json:"extra,omitempty"`
}

// Create 创建代运营订单
func UpData(accountToken string, req CreateReq) (*CreateResp, error) {
	data, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	resp := &CreateResp{}
	err = util.Post2Response2(createGoodsUrl, accountToken, bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
