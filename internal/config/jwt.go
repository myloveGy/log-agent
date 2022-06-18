package config

// Jwt 相关配置信息
type Jwt struct {
	Secret      string `ini:"secret"`       // Jwt 秘钥
	ExpiresTime int    `ini:"expires_time"` // 过期时间(单位秒)
}
