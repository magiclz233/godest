package user

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"godest/pkg/cache"
	"godest/pkg/utils"
)

// Service 用户业务服务
type Service struct {
	repo  Repository
	redis *cache.RedisClient
	jwt   *utils.JWTUtil
	pwd   *utils.PasswordUtil
}

// NewService 创建用户服务
func NewService(
	repo Repository,
	redis *cache.RedisClient,
	jwt *utils.JWTUtil,
	pwd *utils.PasswordUtil,
) *Service {
	return &Service{
		repo:  repo,
		redis: redis,
		jwt:   jwt,
		pwd:   pwd,
	}
}

// Register 用户注册
func (s *Service) Register(username, email, password string) error {
	if _, err := s.repo.GetByUsername(username); err == nil {
		return errors.New("用户名已存在")
	}

	hashedPassword, err := s.pwd.HashPassword(password)
	if err != nil {
		return err
	}

	u := &User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}
	return s.repo.Create(u)
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// Login 用户登录
func (s *Service) Login(username, password string) (*LoginResponse, error) {
	u, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	if !s.pwd.CheckPassword(password, u.Password) {
		return nil, errors.New("密码错误")
	}

	token, err := s.jwt.GenerateToken(u.ID, u.Username)
	if err != nil {
		return nil, errors.New("生成 Token 失败")
	}

	return &LoginResponse{
		Token: token,
		User:  *u,
	}, nil
}

// ListUsers 获取用户列表（带缓存）
func (s *Service) ListUsers() ([]User, error) {
	ctx := context.Background()
	cacheKey := "users:all"

	if s.redis != nil && s.redis.Client != nil {
		if val, err := s.redis.Client.Get(ctx, cacheKey).Result(); err == nil {
			var users []User
			if err := json.Unmarshal([]byte(val), &users); err == nil {
				return users, nil
			}
		}
	}

	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	if s.redis != nil && s.redis.Client != nil {
		if data, err := json.Marshal(users); err == nil {
			s.redis.Client.Set(ctx, cacheKey, data, time.Minute)
		}
	}

	return users, nil
}
