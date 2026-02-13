package utils

import (
	"errors"
	"time"

	"godest/config"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims 自定义 JWT Claims
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWTUtil JWT 工具结构体
type JWTUtil struct{}

// NewJWTUtil 创建 JWT 工具实例
func NewJWTUtil() *JWTUtil {
	return &JWTUtil{}
}

// GenerateToken 生成 JWT Token
func (j *JWTUtil) GenerateToken(userID uint, username string) (string, error) {
	cfg := config.GlobalConfig.JWT
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.Expire) * time.Hour)),
			Issuer:    "godest",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.Secret))
}

// ParseToken 解析 JWT Token
func (j *JWTUtil) ParseToken(tokenString string) (*JWTClaims, error) {
	cfg := config.GlobalConfig.JWT
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
