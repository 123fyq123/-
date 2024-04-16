package public

const (
	ValidatorKey        = "ValidatorKey"
	TranslatorKey       = "TranslatorKey"
	AdminSessionInfoKey = "AdminSessionInfoKey"

	// 规则类型
	LoadTypeHTTP = 0
	LoadTypeTCP  = 1
	LoadTypeGRPC = 2

	// http接入 域名:1or前缀:0
	HTTPRuleTypePrefixURL = 0
	HTTPRuleTypeDomain    = 1

	// daykey和hourkey
	RedisFlowDayKey  = "flow_day_count"
	RedisFlowHourKey = "flow_hour_count"

	// 统计类型：全站、服务、租户
	FlowTotal         = "flow_total"
	FlowServicePrefix = "flow_service"
	FlowAppPrefix     = "flow_app"
)

var (
	LoadTypeMap = map[int]string{
		LoadTypeHTTP: "HTTP",
		LoadTypeTCP:  "TCP",
		LoadTypeGRPC: "GRPC",
	}
)
