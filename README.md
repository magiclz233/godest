# godest

Go 企业级后端服务示例，采用「按业务模块划分，模块内分层」的结构。

## 架构说明

- 业务模块：`internal/user`
- 模块内分层：`api -> app -> domain -> infra`
- 平台层：`internal/platform/http`（路由、中间件）
- DI：Google Wire（编译期注入）
- 部署：Docker + Kustomize（`base + overlays`）

## 目录结构

```text
godest/
├── cmd/server/                         # 应用入口 + Wire 注入
├── config/                             # 配置与加载
├── internal/
│   ├── platform/http/
│   │   ├── middleware/
│   │   └── router/
│   └── user/
│       ├── api/
│       ├── app/
│       ├── domain/
│       └── infra/repository/
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

## Wire

当你修改构造函数依赖关系后，重新生成：

```bash
wire ./cmd/server
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
