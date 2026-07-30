package constant

// APIName 定义 api 名称的 key
const APIName = "api_name"

// IsClusterAPI 标识当前请求是否属于 cluster api
const IsClusterAPI = "is_cluster_api"

// APIRequestEntity 上下文中保存请求实体的 key
const APIRequestEntity = "api_request_entity"

// Prometheus 指标标签名称常量
const (
	LabelMethod     = "method"
	LabelStatus     = "status"
	LabelBkUserName = "bk_username"
	LabelBkAppCode  = "bk_app_code"
	LabelCode       = "code"
	LabelResult     = "result"
)
