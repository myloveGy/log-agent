package response

type UserResponse struct {
	Token         string `json:"token"`
	Username      string `json:"username"`
	LastLoginTime string `json:"last_login_time"`
	LastLoginIp   string `json:"last_login_ip"`
	CreatedAt     string `json:"created_at"`
}
