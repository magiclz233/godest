package app

import (
	"godest/internal/handler"
	"godest/internal/repository"
	"godest/internal/service"
	"godest/internal/transport/http/router"
	"godest/pkg/cache"
	"godest/pkg/utils"

	"github.com/gin-gonic/gin"
)

// App 包含应用所有的依赖项
type App struct {
	Engine *gin.Engine
}

// NewApp 初始化所有依赖并返回一个新的 App 实例
func NewApp() *App {
	// 1. 初始化公共组件 (Pkg/Utils/Cache)
	jwtUtil := utils.NewJWTUtil()
	passwordUtil := utils.NewPasswordUtil()
	redisClient := cache.NewRedisClient()

	// 2. 初始化 Repositories
	repos := initRepositories()

	// 3. 初始化 Services
	services := initServices(repos, redisClient, jwtUtil, passwordUtil)

	// 4. 初始化 Handlers (这里会组装成 handler.Handlers 结构体)
	handlers := initHandlers(services)

	// 5. 初始化 Router
	engine := router.NewRouter(handlers, jwtUtil)

	return &App{
		Engine: engine,
	}
}

// ---------------------------------------------------------
// 内部辅助结构体和函数，用于模块化扩展
// ---------------------------------------------------------

// repositories 结构体用于管理所有的 Repository
type repositories struct {
	user repository.UserRepository
	// Order repository.OrderRepository // 示例：后续扩展订单模块
}

func initRepositories() *repositories {
	return &repositories{
		user: repository.NewUserRepository(),
		// Order: repository.NewOrderRepository(),
	}
}

// services 结构体用于管理所有的 Service
type services struct {
	user *service.UserService
	// order *service.OrderService // 示例：后续扩展订单模块
}

func initServices(repos *repositories, redisClient *cache.RedisClient, jwtUtil *utils.JWTUtil, passwordUtil *utils.PasswordUtil) *services {
	return &services{
		user: service.NewUserService(repos.user, redisClient, jwtUtil, passwordUtil),
		// order: service.NewOrderService(repos.order, ...),
	}
}

// initHandlers 将所有的 Service 注入到对应的 Handler 中
func initHandlers(s *services) *handler.Handlers {
	return &handler.Handlers{
		User: handler.NewUserHandler(s.user),
		// Order: handler.NewOrderHandler(s.order),
	}
}

// Run 启动应用
func (a *App) Run(port string) error {
	return a.Engine.Run(port)
}
