package util

import (
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword 生成密码
func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword 验证面膜
func ValidatePassword(password, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
