package service_test

import (
	"errors"
	"testing"

	"godest/internal/config"
	"godest/internal/model"
	"godest/internal/repository"
	"godest/internal/service"
	"godest/pkg/cache"
	"godest/pkg/utils"

	"github.com/stretchr/testify/assert"
)

type fakeUserRepo struct {
	createFn        func(u *model.User) error
	getByUsernameFn func(username string) (*model.User, error)
	getAllFn        func() ([]model.User, error)
}

func (f *fakeUserRepo) Create(u *model.User) error {
	return f.createFn(u)
}

func (f *fakeUserRepo) GetByUsername(username string) (*model.User, error) {
	return f.getByUsernameFn(username)
}

func (f *fakeUserRepo) GetAll() ([]model.User, error) {
	if f.getAllFn == nil {
		return []model.User{}, nil
	}
	return f.getAllFn()
}

var _ repository.UserRepository = (*fakeUserRepo)(nil)

func TestUserServiceRegister(t *testing.T) {
	config.GlobalConfig = &config.Config{
		JWT: config.JWTConfig{
			Secret: "test_secret",
			Expire: 24,
		},
	}

	repo := &fakeUserRepo{
		getByUsernameFn: func(username string) (*model.User, error) {
			if username == "existing" {
				return &model.User{Username: "existing"}, nil
			}
			return nil, errors.New("not found")
		},
		createFn: func(u *model.User) error { return nil },
	}

	svc := service.NewUserService(repo, &cache.RedisClient{}, utils.NewJWTUtil(), utils.NewPasswordUtil())

	err := svc.Register("testuser", "test@example.com", "password123")
	assert.NoError(t, err)

	err = svc.Register("existing", "existing@example.com", "password123")
	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
}

func TestUserServiceLogin(t *testing.T) {
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
		getByUsernameFn: func(username string) (*model.User, error) {
			if username == "testuser" {
				u := &model.User{
					Username: "testuser",
					Password: hashedPwd,
				}
				u.ID = 1
				return u, nil
			}
			return nil, errors.New("not found")
		},
		createFn: func(u *model.User) error { return nil },
	}

	svc := service.NewUserService(repo, &cache.RedisClient{}, jwtUtil, pwdUtil)

	resp, err := svc.Login("testuser", "secret")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Token)

	resp, err = svc.Login("testuser", "wrongpassword")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "invalid password", err.Error())
}
