package opensearch

type Args struct {
	AppName string `json:"app_name"` // 要查询的应用名
	Query   string `json:"query"`    // 用于设置搜索条件
	Filter  string `json:"filter"`   // 用于设置过滤条件
	Sort    string `json:"sort"`     // 用于设置文档排序条件（仅支持单字段int类型，仅限v3版API及SDK）
}

type ScrollArgs struct {
	Args
	Scroll     string `json:"scroll"`      // 表示下一次 scroll请求的有效期，每次请求都必须设置该参数，可以用1m表示1min；支持的时间单位包括：w=Week, d=Day, h=Hour, m=minute, s=second
	SearchType string `json:"search_type"` // 第一次查询的时必须填写，后续无需填写，后续通过指定 scroll_id 实现下一次查询
	ScrollId   string `json:"scroll_id"`   // 第一次调用scroll方法会返回scroll_id 但并不包含数据，后续每次搜索都必须指定上一次返回scroll_id，并且后续搜索结果中都会返回scroll_id及对应匹配的数据，后续查询该参数必填
}

type SearchArgs struct {
	Args
	Qp             string `json:"qp"`               // 指定要使用的查询分析规则
	Disable        string `json:"disable"`          // 关闭已生效的查询分析功能
	FirstRankName  string `json:"first_rank_name"`  // 设置粗排表达式名字。
	SecondRankName string `json:"second_rank_name"` // 设置精排表达式名字。
	UserId         string `json:"user_id"`          // 用来标识发起当前搜索请求的终端用户。该值可以设置为下列值，优先级从高到低：1. 终端用户的长登录会员ID2. 终端用户的移动设备imei标识
	Abtest         string `json:"abtest"`           // 使用A/B Test功能时需要设置该参数。
	RawQuery       string `json:"raw_query"`        // 用于类目预测等算法训练使用，因此，建议所有查询都设置该参数；
	ReSearch       string `json:"re_search"`        // 用来设置重查策略，当前只支持按照total hits的阈值来设置。
	Biz            string `json:"biz"`              // 用来描述本次请求的相关业务信息 。比如本次请求来源的业务类型。
	Summary        string `json:"summary"`          // 动态摘要的配置
	FromRequestId  string `json:"from_request_id"`  // 本搜索请求从哪里引导而来，如果当前的query来自下拉提示、热词、底纹等功能的推荐列表，那么请求这个推荐列表的request_id可以赋给这个参数，通过关联这个引导事件，可以计算上游功能的各项指标，衡量使用效果，为优化功能提供依据。详见下拉提示产品文档。
}

// SearchResponse - The response body of OpenSearch API.
type SearchResponse struct {
	Status    string       `json:"status"`
	RequestID string       `json:"request_id"`
	Result    SearchResult `json:"result"`
	Errors    []ErrorInfo  `json:"errors"`
	Tracer    string       `json:"tracer"`
}

type SearchResult struct {
	SearchTime float64                  `json:"searchtime"`
	Total      int                      `json:"total"`
	Num        int                      `json:"num"`
	ViewTotal  int                      `json:"viewtotal"`
	ScrollId   string                   `json:"scroll_id"`
	Items      []map[string]interface{} `json:"items"`
}

type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
