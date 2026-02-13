package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"go_pro/internal/model"
	"go_pro/internal/repository"
	"go_pro/pkg/cache"
	"go_pro/pkg/utils"
)

// UserService 用户业务逻辑层
// UserService handles business logic for users
type UserService struct {
	repo  repository.IUserRepository
	redis *cache.RedisClient
	jwt   *utils.JWTUtil
	pwd   *utils.PasswordUtil
}

// NewUserService 创建 UserService 实例
// NewUserService creates a new instance of UserService
func NewUserService(
	repo repository.IUserRepository,
	redis *cache.RedisClient,
	jwt *utils.JWTUtil,
	pwd *utils.PasswordUtil,
) *UserService {
	return &UserService{
		repo:  repo,
		redis: redis,
		jwt:   jwt,
		pwd:   pwd,
	}
}

// Register 用户注册
// Register creates a new user
func (s *UserService) Register(username, email, password string) error {
	// 检查用户是否已存在
	if _, err := s.repo.GetByUsername(username); err == nil {
		return errors.New("用户名已存在")
	}

	// 密码加密
	hashedPassword, err := s.pwd.HashPassword(password)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}
	return s.repo.Create(user)
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}

// Login 用户登录
// Login authenticates a user
func (s *UserService) Login(username, password string) (*LoginResponse, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 校验密码
	if !s.pwd.CheckPassword(password, user.Password) {
		return nil, errors.New("密码错误")
	}

	// 生成 Token
	token, err := s.jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, errors.New("生成 Token 失败")
	}

	return &LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}

// ListUsers 获取用户列表 (带缓存)
// ListUsers returns all users with caching
func (s *UserService) ListUsers() ([]model.User, error) {
	ctx := context.Background()
	cacheKey := "users:all"

	// 1. 尝试从 Redis 获取
	val, err := s.redis.Client.Get(ctx, cacheKey).Result()
	if err == nil {
		var users []model.User
		if err := json.Unmarshal([]byte(val), &users); err == nil {
			return users, nil
		}
	}

	// 2. 从数据库获取
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// 3. 写入 Redis (过期时间 1 分钟)
	if data, err := json.Marshal(users); err == nil {
		s.redis.Client.Set(ctx, cacheKey, data, time.Minute)
	}

	return users, nil
}
