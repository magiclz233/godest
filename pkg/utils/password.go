package utils

import "golang.org/x/crypto/bcrypt"

// PasswordUtil 密码工具结构体
type PasswordUtil struct{}

// NewPasswordUtil 创建密码工具实例
func NewPasswordUtil() *PasswordUtil {
	return &PasswordUtil{}
}

// HashPassword 密码加密
func (u *PasswordUtil) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword 密码校验
func (u *PasswordUtil) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
