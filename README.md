# godest

一个用于练手的简化 Go Web 项目，采用清晰、低复杂度的分层结构。

## 设计目标

- 简单直接：先跑通核心链路，不引入复杂框架
- 职责清晰：HTTP、业务、数据访问分层明确
- 易于迭代：新增业务时按固定位置扩展

## 目录结构

```text
godest/
├── cmd/
│   └── main.go
├── internal/
│   ├── transport/
│   │   └── http/
│   │       ├── router/
│   │       └── middleware/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   └── config/
├── pkg/
├── deploy/
├── go.mod
└── ...
```

## 各目录职责

- `cmd/main.go`
  - 程序入口
  - 负责启动流程：加载配置、初始化日志/数据库、组装依赖、启动 HTTP 服务

- `internal/config`
  - 配置结构定义和加载逻辑
  - 管理配置文件与环境变量映射

- `internal/transport/http/router`
  - 负责路由注册与路由分组
  - 只做路由组织，不写业务逻辑

- `internal/transport/http/middleware`
  - 负责 HTTP 中间件（鉴权、日志、限流等）
  - 只处理横切关注点

- `internal/handler`
  - HTTP 请求/响应适配层
  - 负责参数绑定、输入校验、调用 service、返回 HTTP 响应

- `internal/service`
  - 业务逻辑层
  - 编排业务流程、处理领域规则
  - 不关心 HTTP 细节

- `internal/repository`
  - 数据访问层
  - 封装数据库读写细节，对 service 暴露接口

- `internal/model`
  - 数据结构定义（实体、DTO 等）
  - 避免把业务逻辑塞进 model

- `pkg`
  - 可复用通用能力（如 logger、加密、工具函数）
  - 尽量保持与具体业务解耦

## 分层依赖规则

建议遵循以下依赖方向（单向）：

- `handler -> service -> repository`
- `router -> handler`
- `middleware` 由 `router` 挂载
- `service/repository` 可使用 `model`
- `cmd` 负责依赖注入与组装

避免：

- `repository` 反向调用 `service`
- `service` 依赖 Gin 的 `Context`
- 在 `handler` 中直接写数据库操作

## 新增业务时怎么放

以 `order` 模块为例：

- 新增结构体：`internal/model/order.go`
- 新增仓储：`internal/repository/order_repository.go`
- 新增服务：`internal/service/order_service.go`
- 新增处理器：`internal/handler/order_handler.go`
- 在 `internal/transport/http/router` 注册路由

## 运行

```bash
go run ./cmd
```

## 常用命令

```bash
go fmt ./...
go test ./...
go build -o bin/godest ./cmd
```

## Docker

```bash
docker build -f deploy/docker/Dockerfile -t godest:latest .
docker run --rm -p 8080:8080 godest:latest
```