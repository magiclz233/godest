package service_test

import (
	"errors"
	"testing"

	"go_pro/config"
	"go_pro/internal/model"
	"go_pro/internal/repository/mock"
	"go_pro/internal/service"
	"go_pro/pkg/cache"
	"go_pro/pkg/utils"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserService_Register(t *testing.T) {
	// 初始化配置
	config.GlobalConfig = &config.Config{
		JWT: config.JWTConfig{
			Secret: "test_secret",
			Expire: 24,
		},
	}

	// 初始化 Mock 控制器
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建 Mock Repository
	mockRepo := mock.NewMockIUserRepository(ctrl)

	// 初始化依赖
	// 注意：这里为了简化测试，Redis 传 nil (Register 方法不使用 Redis)
	// JWT 和 PasswordUtil 使用真实实例
	svc := service.NewUserService(mockRepo, &cache.RedisClient{}, utils.NewJWTUtil(), utils.NewPasswordUtil())

	t.Run("Success", func(t *testing.T) {
		// 模拟 GetByUsername 返回错误 (用户不存在)
		mockRepo.EXPECT().GetByUsername("testuser").Return(nil, errors.New("user not found"))

		// 模拟 Create 成功
		mockRepo.EXPECT().Create(gomock.Any()).Return(nil)

		err := svc.Register("testuser", "test@example.com", "password123")
		assert.NoError(t, err)
	})

	t.Run("UserAlreadyExists", func(t *testing.T) {
		// 模拟 GetByUsername 返回用户 (用户已存在)
		mockRepo.EXPECT().GetByUsername("existing").Return(&model.User{}, nil)

		err := svc.Register("existing", "test@example.com", "password123")
		assert.Error(t, err)
		assert.Equal(t, "用户名已存在", err.Error())
	})
}

func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockIUserRepository(ctrl)
	pwdUtil := utils.NewPasswordUtil()
	jwtUtil := utils.NewJWTUtil()
	svc := service.NewUserService(mockRepo, &cache.RedisClient{}, jwtUtil, pwdUtil)

	// 预先生成加密密码
	hashedPwd, _ := pwdUtil.HashPassword("secret")

	t.Run("Success", func(t *testing.T) {
		user := &model.User{
			Username: "testuser",
			Password: hashedPwd,
		}
		user.ID = 1

		mockRepo.EXPECT().GetByUsername("testuser").Return(user, nil)

		resp, err := svc.Login("testuser", "secret")
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Token)
		assert.Equal(t, "testuser", resp.User.Username)
	})

	t.Run("WrongPassword", func(t *testing.T) {
		user := &model.User{
			Username: "testuser",
			Password: hashedPwd,
		}

		mockRepo.EXPECT().GetByUsername("testuser").Return(user, nil)

		resp, err := svc.Login("testuser", "wrongpassword")
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "密码错误", err.Error())
	})
}
