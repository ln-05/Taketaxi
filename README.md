# Taketaxi

## 项目简介
Taketaxi 是一个基于 Go 语言和 Kratos 框架开发的后端服务示例，包含用户注册、登录、短信验证码等基础功能。项目采用分层架构，集成了 MySQL 和 Redis，支持 HTTP 和 gRPC 两种服务接口，适合学习和二次开发。

## 项目结构
```
Taketaxi/
├── helloworld/           # 后端服务主目录
│   ├── api/              # Protobuf API 定义
│   ├── cmd/              # 服务启动入口
│   ├── configs/          # 配置文件
│   ├── internal/         # 内部应用代码（分层架构）
│   │   ├── biz/          # 业务逻辑层
│   │   ├── conf/         # 配置结构体定义
│   │   ├── data/         # 数据访问层（数据库/缓存）
│   │   ├── server/       # 服务启动（HTTP/gRPC）
│   │   └── service/      # 服务实现层
│   ├── pkg/              # 公共包，如 jwt
│   ├── third_party/      # 第三方 proto 依赖
│   ├── Makefile          # 构建和生成脚本
│   ├── Dockerfile        # Docker 镜像构建文件
│   └── README.md         # 后端说明文档
├── qian/                 # 前端项目目录（uni-app）
│   ├── pages/            # 页面目录
│   ├── static/           # 静态资源
│   ├── cloudfunctions-aliyun/ # 云函数目录
│   ├── App.vue           # 入口组件
│   ├── main.js           # 入口 JS
│   └── ...               # 其他前端相关文件
└── README.md             # 项目总览说明文档
```

- 详细的后端结构说明可见 `helloworld/README.md`
- 前端为 uni-app 项目，适合小程序开发

## 安装方法
1. 克隆项目：
   ```bash
   git clone <your-repo-url>
   cd Taketaxi/helloworld
   ```
2. 安装依赖：
   ```bash
   go mod tidy
   make init
   ```
3. 配置数据库和 Redis（可选）：
   编辑 `configs/config.yaml`，根据实际情况修改数据库和 Redis 连接信息。

4. 生成代码：
   ```bash
   make all
   ```
5. 构建并运行服务：
   ```bash
   go build -o ./bin/ ./...
   ./bin/helloworld -conf ./configs/config.yaml
   ```
   或使用 Makefile：
   ```bash
   make build
   ./bin/helloworld -conf ./configs/config.yaml
   ```

## 使用方法
### HTTP API 示例
- 发送短信验证码：
  ```bash
  curl -X POST http://127.0.0.1:8080/user/sendSms -d '{"mobile":"13800000000","SendSmsCode":"123456"}' -H 'Content-Type: application/json'
  ```
- 用户登录：
  ```bash
  curl -X POST http://127.0.0.1:8080/user/login -d '{"mobile":"13800000000","SendSmsCode":"123456"}' -H 'Content-Type: application/json'
  ```

### gRPC API
可参考 `api/helloworld/v1/greeter.proto` 文件定义，使用 grpcurl 或自定义客户端调用。

### 其他
- 更多命令可通过 `make help` 查看。
- 支持 Docker 部署，详见 `helloworld/README.md`。
