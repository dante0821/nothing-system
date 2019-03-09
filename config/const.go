package config

const (
	DefaultPassWd = "123456abc"
)

const (
	AmountFixed   = 10
	CurrencyFixed = 2
)
const (
	MerchantPayType          = 1101 // 普通用户充值
	MerchantRechangeType     = 1102 // 商户充值
	MerchantRefundType       = 1201 // 商户退款支出
	MerchantWithdrawType     = 1202 // 商户提币
	MerchantWithdrawCashType = 1203 // 商户提现
	OtcAderBuyType           = 2101 // otc广告主买币
	OtcAderRechangeType      = 2103 // otc商户充值
	OtcAderSellType          = 2201 // otc广告主卖币
	OtcAderWithdrawType      = 2202 // otc广告提币
	AgentPayType             = 3101 // 普通用户充值分佣
	AgentRefundType          = 3201 // 商户退款支出分佣
	AgentWithdrawType        = 3202 // 代理商提币
	AgentWithdrawCashType    = 3203 // 代理商提现
)

const (
	OtcAdOnStatus       = 1 //交易中
	OtcAdOffStatus      = 2 //已下架
	OtcAdCancelStatus   = 3 //已撤销
	OtcAdCompleteStatus = 4 //已完成
)

const (
	AuthSession = "Session"
	OrderNumber = "order_number"
)
const (
	Merchant = "merchant"
	OtcAder  = "otc_ader"
	User     = "user"
	Agent    = "agent"
)

const (
	OrderPayType    = 1
	OrderRefundType = 2
)
const (
	NormalUserType   = 0
	MerchantUserType = 1
	AgentUserType    = 2
)
const (
	OtcAdSellType = 1
	OtcAdBuyType  = 2
)

const (
	// 1.创建、2.支付、3.取消、4.冻结、5.完成  6.客服完成、7.客服取消 8.超时关闭
	OtcTransacCreateType        = 1
	OtcTransacPayType           = 2
	OtcTransacCancelType        = 3
	OtcTransacFreezeType        = 4
	OtcTransacFinishType        = 5
	OtcTransacServiceFinishType = 6
	OtcTransacServiceCancelType = 7
	OtcTransacCloseType         = 8
)

const (
	OrderNoActiveType = 0
	OrderCreateType   = 1 //待付款
	OrderFinishType   = 2 //完成
	OrderCloseType    = 3 //取消
)

const (
	SafeGoogleType   = 1 //谷歌验证
	SafeTelPhoneType = 2 //手机验证
	SafeEmailType    = 3 //邮箱验证
)
