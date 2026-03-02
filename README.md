# godest

Go 企业级后端服务示例，采用「按业务模块划分」的结构。

## 架构说明

- 业务模块：`internal/user`
- 模块代码：统一放在 `internal/user` 包下，按业务语义分文件
- 平台层：`internal/platform/http`（路由、中间件）
- DI：手写构造函数注入（简单直接）
- 部署：Docker + Kustomize（`base + overlays`）

## 目录结构

```text
godest/
├── cmd/server/                         # 应用入口 + 手写 DI 组装
├── config/                             # 配置与加载
├── internal/
│   ├── platform/http/
│   │   ├── middleware/
│   │   └── router/
│   └── user/                           # 用户业务包
├── pkg/                                # 通用组件
├── deploy/
│   ├── docker/
│   │   └── Dockerfile
│   └── k8s/
│       ├── base/
│       └── overlays/
│           ├── dev/
│           └── prod/
└── Makefile
```

## 配置

- 默认从 `config/config.yaml` 读取
- 支持环境变量覆盖（前缀：`GODEST_`）
- 例子：`GODEST_APP_PORT=:9090`、`GODEST_JWT_SECRET=xxx`

## 常用命令

```bash
go fmt ./...
go test ./...
go run ./cmd/server
```

## Docker

```bash
docker build -f deploy/docker/Dockerfile -t godest:latest .
docker run --rm -p 8080:8080 godest:latest
```

## K8s（Kustomize）

开发环境：

```bash
kubectl apply -k deploy/k8s/overlays/dev
```

生产环境：

```bash
kubectl apply -k deploy/k8s/overlays/prod
```

说明：

- `base` 放通用资源
- `overlays/dev|prod` 做副本数等环境差异化
