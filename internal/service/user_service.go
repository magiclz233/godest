package service

import (
	"context"
	"errors"

	internalcache "godest/internal/cache"
	"godest/internal/model"
	"godest/internal/repository"
	pkgcache "godest/pkg/cache"
	"godest/pkg/utils"

	"gorm.io/gorm"
)

type UserService struct {
	repo  repository.UserRepository
	cache *internalcache.UserCache
	jwt   *utils.JWTUtil
	pwd   *utils.PasswordUtil
}

var (
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrTokenGenerateFailed = errors.New("failed to generate token")
)

func NewUserService(
	repo repository.UserRepository,
	redis *pkgcache.RedisClient,
	jwt *utils.JWTUtil,
	pwd *utils.PasswordUtil,
) *UserService {
	return &UserService{
		repo:  repo,
		cache: internalcache.NewUserCache(redis),
		jwt:   jwt,
		pwd:   pwd,
	}
}

func (s *UserService) Register(username, email, password string) error {
	if _, err := s.repo.GetByUsername(username); err == nil {
		return ErrUserAlreadyExists
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
	if err := s.repo.Create(u); err != nil {
		return err
	}

	ctx := context.Background()
	s.cache.Set(ctx, u)
	s.cache.InvalidateList(ctx)
	return nil
}

func (s *UserService) Login(username, password string) (*model.LoginResponse, error) {
	u, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if !s.pwd.CheckPassword(password, u.Password) {
		return nil, ErrInvalidPassword
	}

	token, err := s.jwt.GenerateToken(u.ID, u.Username)
	if err != nil {
		return nil, ErrTokenGenerateFailed
	}

	return &model.LoginResponse{
		Token: token,
		User:  *u,
	}, nil
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	ctx := context.Background()
	if user, ok := s.cache.GetByID(ctx, id); ok {
		return user, nil
	}

	u, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	s.cache.Set(ctx, u)
	return u, nil
}

func (s *UserService) ListUsers() ([]model.User, error) {
	ctx := context.Background()

	if users, ok := s.cache.GetAll(ctx); ok {
		return users, nil
	}

	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	s.cache.SetAll(ctx, users)

	return users, nil
}
