package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"godest/internal/model"
	"godest/internal/repository"
	"godest/pkg/cache"
	"godest/pkg/utils"
)

type UserService struct {
	repo  repository.UserRepository
	redis *cache.RedisClient
	jwt   *utils.JWTUtil
	pwd   *utils.PasswordUtil
}

func NewUserService(
	repo repository.UserRepository,
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

func (s *UserService) Register(username, email, password string) error {
	if _, err := s.repo.GetByUsername(username); err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := s.pwd.HashPassword(password)
	if err != nil {
		return err
	}

	u := &model.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}
	return s.repo.Create(u)
}

func (s *UserService) Login(username, password string) (*model.LoginResponse, error) {
	u, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !s.pwd.CheckPassword(password, u.Password) {
		return nil, errors.New("invalid password")
	}

	token, err := s.jwt.GenerateToken(u.ID, u.Username)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &model.LoginResponse{
		Token: token,
		User:  *u,
	}, nil
}

func (s *UserService) ListUsers() ([]model.User, error) {
	ctx := context.Background()
	cacheKey := "users:all"

	if s.redis != nil && s.redis.Client != nil {
		if val, err := s.redis.Client.Get(ctx, cacheKey).Result(); err == nil {
			var users []model.User
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
