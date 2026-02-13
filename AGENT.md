# AGENT 说明（godest）

## 语言要求

- 回复语言：简体中文
- 代码注释：简体中文
- 对外接口返回的文案：优先简体中文（与现有接口保持一致）

## 项目概览

- Go 模块名：`godest`
- Web 框架：Gin
- 依赖注入：Google Wire（`cmd/server/wire.go` + 生成文件 `wire_gen.go`）
- 配置管理：Viper（`config/config.go`，配置文件 `config/config.yaml`）
- 数据库：GORM + SQLite（默认），初始化在 `pkg/database`
- 缓存：Redis（`pkg/cache`）
- 鉴权：JWT（`pkg/utils/jwt.go`），中间件在 `internal/middleware/auth.go`
- 日志：Zap（`pkg/logger`）
- 测试：Testify + GoMock（mock 代码在 `internal/repository/mock`）

## 目录结构（关键路径）

- `cmd/server`：应用入口与 Wire 组装
- `config`：配置结构与加载逻辑
- `internal/handler`：HTTP Handler
- `internal/service`：业务逻辑（包含缓存策略）
- `internal/repository`：数据访问层（接口 + 实现 + mock）
- `pkg`：可复用基础组件（cache/database/logger/utils）
- `router`：路由注册

## 代码规范

- Import 路径统一使用 `godest/...`，不要再出现旧模块名
- 代码格式化使用 `gofmt`
- 变更依赖注入关系后，需要重新生成 Wire：
  - `wire cmd/server/wire.go`
- 不要把密钥、密码等敏感信息写入代码或提交到仓库；本地开发应通过配置文件/环境变量管理

## 常用命令

```bash
go fmt ./...
go mod tidy
go test ./...
go vet ./...
go run cmd/server/main.go
wire cmd/server/wire.go
```

