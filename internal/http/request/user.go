package request

type UserRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserUpdateRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"omitempty,min=6"`
	Status   string `json:"status" validate:"omitempty,oneof=Y N"`
}

type UserQueryRequest struct {
	// 查询条件
	Query map[string]interface{} `json:"query"`
	// 分页第几页
	Page int64 `json:"page" validate:"required,min=1"`
	// 分页每页数量
	PageSize int64 `json:"page_size" validate:"required,min=10"`
}
