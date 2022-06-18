package request

type DatabaseQueryRequest struct {
	// 查询的集合名称
	Collection string `json:"collection" validate:"required,min=1"`
	// 查询条件
	Query map[string]interface{} `json:"query"`
	// 分页第几页
	Page int64 `json:"page" validate:"required,min=1"`
	// 分页每页数量
	PageSize int64 `json:"page_size" validate:"required,min=10"`
	// 排序条件
	Order string `json:"order" validate:"omitempty,oneof=asc desc"`
	// 开始时间
	StartTime string `json:"start_time"`
	// 结束时间
	EndTime string `json:"end_time"`
}
