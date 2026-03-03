# AGENT 说明（godest）

## 语言约定

- 默认使用简体中文沟通
- 代码标识符、日志、错误信息保持原语言
- 代码注释可中英文，但应简洁、准确

## 项目定位

`godest` 是一个练手项目，目标是用尽量简单的分层结构完成完整 Web 后端链路。

## 目录职责（必须遵守）

- `cmd/main.go`
  - 应用入口与依赖组装（手写 DI）
  - 仅做启动编排，不写具体业务

- `internal/config`
  - 配置结构定义
  - 配置加载（文件 + 环境变量）

- `internal/transport/http/router`
  - 路由注册与分组
  - 将 handler 与 middleware 组装到路由上

- `internal/transport/http/middleware`
  - HTTP 中间件（认证、鉴权、请求链路增强）

- `internal/handler`
  - 处理 HTTP 输入输出
  - 参数绑定、校验、状态码与响应格式
  - 不直接访问数据库

- `internal/service`
  - 业务逻辑与流程编排
  - 调用 repository 完成持久化
  - 不依赖 Gin/HTTP 细节

- `internal/repository`
  - 数据访问实现（GORM/SQL）
  - 对外暴露接口，屏蔽存储细节

- `internal/model`
  - 实体与数据结构定义（Entity/DTO）

- `pkg`
  - 通用组件（cache/database/logger/utils）
  - 尽量保持业务无关

## 依赖方向规范

- 允许：`handler -> service -> repository`
- 允许：`router -> handler`，`router -> middleware`
- 允许：`service/repository -> model`
- 允许：`cmd -> 所有需要组装的包`
- 禁止：跨层反向依赖（如 `repository -> service`）
- 禁止：在 `handler` 写业务规则或数据库 SQL

## 开发规则

- 保持 KISS：优先简单可维护，不做过度抽象
- 保持 YAGNI：只实现当前明确需求
- 保持 DRY：重复逻辑提取到公共函数或组件
- 保持 SRP：每个文件/类型职责单一

## 新增功能放置规范

新增一个功能（如 `order`）时，按以下顺序落地：

1. `internal/model/order.go`
2. `internal/repository/order_repository.go`
3. `internal/service/order_service.go`
4. `internal/handler/order_handler.go`
5. `internal/transport/http/router` 中注册对应路由

## 常用命令

```bash
go fmt ./...
go mod tidy
go test ./...
go vet ./...
go run ./cmd
```