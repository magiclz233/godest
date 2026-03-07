很好 👍 我给你整理了一套 **Go 后端工程师学习路线图（从入门 → 能做项目 → 接近工作水平）**。
这套路线是很多 Go 开发者实际走过的路径，**大约 3–6 个月可以达到比较扎实的水平**（取决于每天投入时间）。

我把它分成 **5 个阶段 + 项目练习**。

---

# 一、Go 基础阶段（必须打牢）

目标：**熟练写 Go 代码，不查文档也能写大部分基础代码**

重点知识：

### 1 基础语法

必须熟练：

* 变量 / 常量
* if / for / switch
* 函数
* defer
* panic / recover

示例：

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divide by zero")
    }
    return a / b, nil
}
```

---

### 2 数据结构

必须掌握：

* array
* slice ⭐
* map ⭐
* struct ⭐

示例：

```go
type User struct {
    ID   int
    Name string
}
```

---

### 3 指针

Go 里经常用：

```go
func updateName(u *User) {
    u.Name = "Tom"
}
```

---

### 4 interface（非常重要）

```go
type Reader interface {
    Read(p []byte) (int, error)
}
```

必须理解：

* interface 本质
* duck typing
* 接口解耦

---

### 5 package 和模块

```bash
go mod init myapp
go mod tidy
```

理解：

* 包结构
* 模块管理
* import

---

# 二、Go 标准库（核心能力）

Go 强大的地方就是 **标准库**。

必须熟练：

### 文件操作

```go
os.ReadFile
os.WriteFile
```

---

### JSON

```go
json.Marshal
json.Unmarshal
```

---

### HTTP

```go
http.Get()
http.ListenAndServe()
```

---

### 时间

```go
time.Now()
time.Sleep()
```

---

### context

Go 后端非常常用：

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
```

---

# 三、Go 并发（Go 的灵魂）

这是 Go 最核心的能力。

### goroutine

```go
go func() {
    fmt.Println("hello")
}()
```

---

### channel

```go
ch := make(chan int)

go func() {
    ch <- 1
}()

fmt.Println(<-ch)
```

---

### select

```go
select {
case msg := <-ch:
    fmt.Println(msg)
case <-time.After(time.Second):
    fmt.Println("timeout")
}
```

---

### sync 包

必须掌握：

```
sync.Mutex
sync.RWMutex
sync.WaitGroup
sync.Once
```

---

# 四、Go Web 开发

开始进入 **后端开发能力**。

推荐框架：

```
Gin
Echo
Fiber
```

建议先学 **Gin**。

示例：

```go
r := gin.Default()

r.GET("/hello", func(c *gin.Context) {
    c.JSON(200, gin.H{"msg": "hello"})
})

r.Run()
```

---

需要掌握：

* 路由
* handler
* middleware
* JSON API
* 参数校验

---

# 五、数据库

Go 后端必须会数据库。

### ORM

推荐：

```
gorm
```

示例：

```go
type User struct {
    ID   uint
    Name string
}

db.Create(&User{Name: "Tom"})
```

---

必须理解：

* CRUD
* 事务
* 索引
* 连接池

---

# 六、缓存

必须掌握：

```
Redis
```

Go 客户端：

```
go-redis
```

示例：

```go
rdb.Set(ctx, "key", "value", time.Hour)
```

---

# 七、认证系统

必须会：

```
JWT
```

登录流程：

```
用户登录
   ↓
生成 JWT
   ↓
客户端保存 token
   ↓
请求带 token
   ↓
middleware 验证
```

---

# 八、Go 项目工程化

必须掌握：

### 项目结构

推荐：

```text
cmd/
internal/
pkg/
configs/
```

---

### 日志

推荐：

```
zap
logrus
```

---

### 配置

推荐：

```
viper
```

---

### 测试

```bash
go test
```

---

# 九、Go 必学并发模式

至少掌握 **10 个模式**：

1️⃣ Worker Pool

2️⃣ Fan Out / Fan In

3️⃣ Pipeline

4️⃣ Context Cancel

5️⃣ Rate Limit

6️⃣ Pub/Sub

7️⃣ Future Pattern

8️⃣ Semaphore

9️⃣ Task Queue

🔟 Timeout Pattern

---

# 十、完整项目练习

按顺序做：

### 项目 1

CLI Todo 工具

---

### 项目 2

并发爬虫

---

### 项目 3

用户管理 API

技术：

```
Gin + MySQL + Gorm
```

---

### 项目 4

登录认证系统

技术：

```
JWT
Redis
Middleware
```

---

### 项目 5

任务队列系统

练：

```
goroutine
worker pool
```

---

### 项目 6（最终项目）

博客系统 API

技术：

```
Gin
MySQL
Redis
JWT
Docker
```

功能：

```
用户
文章
评论
点赞
```

---

# 十一、Go 后端必备技术栈

建议学习：

### Web

```
Gin
```

### ORM

```
Gorm
sqlx
```

### DB

```
MySQL
PostgreSQL
```

### Cache

```
Redis
```

### MQ

```
Kafka
RabbitMQ
```

### RPC

```
gRPC
```

### DevOps

```
Docker
K8s（进阶）
```

---

# 十二、Go 必看开源项目（非常重要）

学习高手怎么写代码。

推荐：

### Docker

```
github.com/moby/moby
```

---

### Kubernetes

```
github.com/kubernetes/kubernetes
```

---

### Gin

```
github.com/gin-gonic/gin
```

---

### etcd

```
github.com/etcd-io/etcd
```

---

# 最后给你一个非常关键的建议

Go 学习一定要：

**边学边写项目。**

只看语法：

```
永远学不会 Go
```

但如果你：

```
写 6 个项目
```

基本就能达到 **Go 后端入门 → 熟练**。

---


这份路线本身没问题，但你在这个项目里不要按“看完一章再做下一章”的方式学。更有效的是把 [practice.md](e:\code\golang\godest\practice.md#L1) 当成知识索引，把当前这个仓库当成唯一练习场，每学一个点，立刻在项目里落一处代码、一条测试、一个总结。

这个仓库很适合这样学，因为分层已经有了：入口在 [cmd/main.go](e:\code\golang\godest\cmd\main.go#L1)，组装在 [internal/app/app.go](e:\code\golang\godest\internal\app\app.go#L1)，业务主线在 [internal/service/user_service.go](e:\code\golang\godest\internal\service\user_service.go#L1)，HTTP 适配在 [internal/handler/user_handler.go](e:\code\golang\godest\internal\handler\user_handler.go#L1)，数据访问在 [internal/repository/user_repository.go](e:\code\golang\godest\internal\repository\user_repository.go#L1)，配置在 [internal/config/config.go](e:\code\golang\godest\internal\config\config.go#L1)。你现在最该做的不是新开很多项目，而是先把“用户模块”完整吃透，再扩一个新模块。

建议你按这个顺序在本项目里推进：

1. 基础语法 + package + interface  
直接吃透 `handler -> service -> repository` 这条链路，重点看接口抽象为什么放在 repository 层，先能自己从头讲清楚这套依赖方向。

2. 标准库 + JSON + HTTP + context  
从注册、登录、查用户这几个接口入手，自己手写调用流程图。然后把 service 方法逐步改成带 `context.Context`，理解请求上下文是怎么往下传的。

3. 测试  
先把 [internal/service/user_service_test.go](e:\code\golang\godest\internal\service\user_service_test.go#L1) 补完整。当前只有部分 service 测试，你应该继续补：
   - `ListUsers`
   - 注册失败分支
   - 登录用户不存在分支
   - 缓存命中/缓存失效分支  
测试是你把 Go 学扎实的最快方式。

4. 数据库 + ORM  
把 `User` 模块彻底做完，不要停在 CRUD 表面。继续加：
   - 唯一索引
   - 分页查询
   - 按邮箱查询
   - 更新密码
   - 软删除/删除后的查询行为  
这样你会真正熟悉 GORM 和 repository 设计。

5. 缓存  
现在 `ListUsers` 已经有 Redis 缓存雏形了，在 [internal/service/user_service.go](e:\code\golang\godest\internal\service\user_service.go#L1)。下一步你要补的是：
   - 注册成功后删除用户列表缓存
   - 缓存 key 设计
   - 序列化失败怎么处理
   - Redis 不可用时是否降级  
这一步能把“业务逻辑”和“基础设施容错”串起来。

6. 认证系统  
现有登录/JWT 已经在做了。你接下来应该加：
   - token 过期处理
   - 从中间件取当前用户
   - 受保护接口里的用户身份校验
   - 刷新 token 或登出黑名单机制  
这会让你真正理解中间件、JWT、鉴权边界。

7. 工程化  
这个项目已经有 Docker、K8s、配置、日志基础。你下一步可以做：
   - 区分 dev/prod 配置
   - 统一错误码
   - request id 日志
   - `Makefile` 补常用命令
   - README 增加接口说明和启动步骤

最重要的一点：你不要同时学 12 个主题，而是给自己固定一个循环。

- 先选一个知识点
- 在本项目里改一个小功能
- 给它补测试
- 写 5 行总结：学到了什么、踩了什么坑、下一步改什么

如果你愿意，我下一步可以直接帮你把这个仓库拆成一个“4 周学习任务表”，精确到每天该看哪个文件、改哪个功能、补哪条测试。

