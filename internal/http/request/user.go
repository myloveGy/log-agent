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
