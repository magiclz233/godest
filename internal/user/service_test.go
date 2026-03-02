package user_test

import (
	"errors"
	"testing"

	"godest/config"
	"godest/internal/user"
	"godest/pkg/cache"
	"godest/pkg/utils"

	"github.com/stretchr/testify/assert"
)

type fakeUserRepo struct {
	createFn        func(u *user.User) error
	getByUsernameFn func(username string) (*user.User, error)
	getAllFn        func() ([]user.User, error)
}

func (f *fakeUserRepo) Create(u *user.User) error {
	return f.createFn(u)
}

func (f *fakeUserRepo) GetByUsername(username string) (*user.User, error) {
	return f.getByUsernameFn(username)
}

func (f *fakeUserRepo) GetAll() ([]user.User, error) {
	if f.getAllFn == nil {
		return []user.User{}, nil
	}
	return f.getAllFn()
}

func TestUserServiceRegister(t *testing.T) {
	config.GlobalConfig = &config.Config{
		JWT: config.JWTConfig{
			Secret: "test_secret",
			Expire: 24,
		},
	}

	repo := &fakeUserRepo{
		getByUsernameFn: func(username string) (*user.User, error) {
			if username == "existing" {
				return &user.User{Username: "existing"}, nil
			}
			return nil, errors.New("not found")
		},
		createFn: func(u *user.User) error { return nil },
	}

	svc := user.NewService(repo, &cache.RedisClient{}, utils.NewJWTUtil(), utils.NewPasswordUtil())

	err := svc.Register("testuser", "test@example.com", "password123")
	assert.NoError(t, err)

	err = svc.Register("existing", "existing@example.com", "password123")
	assert.Error(t, err)
	assert.Equal(t, "用户名已存在", err.Error())
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
		getByUsernameFn: func(username string) (*user.User, error) {
			if username == "testuser" {
				u := &user.User{
					Username: "testuser",
					Password: hashedPwd,
				}
				u.ID = 1
				return u, nil
			}
			return nil, errors.New("not found")
		},
		createFn: func(u *user.User) error { return nil },
	}

	svc := user.NewService(repo, &cache.RedisClient{}, jwtUtil, pwdUtil)

	resp, err := svc.Login("testuser", "secret")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Token)

	resp, err = svc.Login("testuser", "wrongpassword")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "密码错误", err.Error())
}
