package app_test

import (
	"errors"
	"testing"

	"godest/config"
	"godest/internal/user/app"
	"godest/internal/user/domain"
	"godest/pkg/cache"
	"godest/pkg/utils"

	"github.com/stretchr/testify/assert"
)

type fakeUserRepo struct {
	createFn        func(user *domain.User) error
	getByUsernameFn func(username string) (*domain.User, error)
	getAllFn        func() ([]domain.User, error)
}

func (f *fakeUserRepo) Create(user *domain.User) error {
	return f.createFn(user)
}

func (f *fakeUserRepo) GetByUsername(username string) (*domain.User, error) {
	return f.getByUsernameFn(username)
}

func (f *fakeUserRepo) GetAll() ([]domain.User, error) {
	if f.getAllFn == nil {
		return []domain.User{}, nil
	}
	return f.getAllFn()
}

func TestUserService_Register(t *testing.T) {
	config.GlobalConfig = &config.Config{
		JWT: config.JWTConfig{
			Secret: "test_secret",
			Expire: 24,
		},
	}

	repo := &fakeUserRepo{
		getByUsernameFn: func(username string) (*domain.User, error) {
			if username == "existing" {
				return &domain.User{Username: "existing"}, nil
			}
			return nil, errors.New("not found")
		},
		createFn: func(user *domain.User) error { return nil },
	}

	svc := app.NewUserService(repo, &cache.RedisClient{}, utils.NewJWTUtil(), utils.NewPasswordUtil())

	err := svc.Register("testuser", "test@example.com", "password123")
	assert.NoError(t, err)

	err = svc.Register("existing", "existing@example.com", "password123")
	assert.Error(t, err)
	assert.Equal(t, "用户名已存在", err.Error())
}

func TestUserService_Login(t *testing.T) {
	config.GlobalConfig = &config.Config{
		JWT: config.JWTConfig{
			Secret: "test_secret",
			Expire: 24,
		},
	}

	pwdUtil := utils.NewPasswordUtil()
	jwtUtil := utils.NewJWTUtil()
	hashedPwd, _ := pwdUtil.HashPassword("secret")

	repo := &fakeUserRepo{
		getByUsernameFn: func(username string) (*domain.User, error) {
			if username == "testuser" {
				user := &domain.User{
					Username: "testuser",
					Password: hashedPwd,
				}
				user.ID = 1
				return user, nil
			}
			return nil, errors.New("not found")
		},
		createFn: func(user *domain.User) error { return nil },
	}

	svc := app.NewUserService(repo, &cache.RedisClient{}, jwtUtil, pwdUtil)

	resp, err := svc.Login("testuser", "secret")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Token)

	resp, err = svc.Login("testuser", "wrongpassword")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "密码错误", err.Error())
}
