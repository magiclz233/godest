# AGENT 说明（godest）

## 语言要求

- 回复语言：简体中文
- 代码注释：简体中文
- 对外接口返回的文案：优先简体中文（与现有接口保持一致）

## 项目概览

- Go 模块名：`godest`
- Web 框架：Gin
- 依赖注入：手写构造函数注入（`cmd/server/app.go`）
- 配置管理：Viper（`config/config.go`，配置文件 `config/config.yaml`）
- 数据库：GORM + SQLite（默认），初始化在 `pkg/database`
- 缓存：Redis（`pkg/cache`）
- 鉴权：JWT（`pkg/utils/jwt.go`），中间件在 `internal/platform/http/middleware/auth.go`
- 日志：Zap（`pkg/logger`）
- 测试：Testify（用户服务测试在 `internal/user/service_test.go`）

## 目录结构（关键路径）

- `cmd/server`：应用入口与手写 DI 组装
- `config`：配置结构与加载逻辑
- `internal/user`：用户业务包（实体、仓储、服务、HTTP Handler）
- `internal/platform/http/router`：路由注册
- `internal/platform/http/middleware`：HTTP 中间件
- `pkg`：可复用基础组件（cache/database/logger/utils）

## 代码规范

- Import 路径统一使用 `godest/...`，不要再出现旧模块名
- 代码格式化使用 `gofmt`
- 不要把密钥、密码等敏感信息写入代码或提交到仓库；本地开发应通过配置文件/环境变量管理

## 常用命令

```bash
go fmt ./...
go mod tidy
go test ./...
go vet ./...
go run ./cmd/server
```

