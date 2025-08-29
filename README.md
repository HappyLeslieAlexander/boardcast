<div align="center">
    <svg xmlns="http://www.w3.org/2000/svg" width="120" height="120" viewBox="0 0 14 14">
        <g fill="none" fill-rule="evenodd" clip-rule="evenodd">
            <path fill="#8fbffa" d="M.58 1.961A1.92 1.92 0 0 1 1.937 1.4H12.06a1.92 1.92 0 0 1 1.92 1.92v7.362a1.92 1.92 0 0 1-1.92 1.92H6.995A6.283 6.283 0 0 0 .017 5.524V3.32c0-.51.202-.998.563-1.358Z"/>
            <path fill="#2859c5" d="M.768 6.73a.75.75 0 1 0 0 1.5a3.533 3.533 0 0 1 3.533 3.532a.75.75 0 0 0 1.5 0A5.033 5.033 0 0 0 .768 6.73m0 2.676a.75.75 0 0 0 0 1.5a.856.856 0 0 1 .856.856a.75.75 0 1 0 1.5 0A2.356 2.356 0 0 0 .768 9.406"/>
        </g>
    </svg>
</div>

<div align="center">
    <strong>一个简单、安全的实时协作白板应用</strong>
</div>

<div align="center">
    <a href="#特性">特性</a> •
    <a href="#上手">上手</a> •
    <a href="#安装">安装</a> •
    <a href="#配置">配置</a> •
    <a href="#部署">部署</a> •
    <a href="#开发">开发</a>
</div>

## 概述

BoardCast 是一个轻量级的实时协作白板应用，使用 Go 语言开发。它允许多个用户通过密码认证后在同一个白板上实时协作编辑文本内容。应用具有简洁的用户界面，支持 WebSocket 实时同步，是团队协作、会议记录、头脑风暴的理想工具。

## 特性

### 🔐 安全认证
- 基于密码的访问控制
- 会话管理和安全的 Cookie 存储
- bcrypt 密码加密

### 🔄 实时协作
- WebSocket 实时通信
- 多用户同步编辑
- 自动内容保存和恢复

### 📱 响应式设计
- 适配桌面和移动设备
- 简洁直观的用户界面
- 现代化的设计风格

### 🚀 轻量高效
- 单一二进制文件部署
- 低资源占用
- 无数据库依赖

### 🛠️ 运维友好
- Docker 容器化支持
- 多平台二进制发布
- 优雅关闭和错误处理
- 结构化日志

## 上手

### 使用 Docker（推荐）

```bash
# 拉取镜像
docker pull ghcr.io/yosebyte/boardcast:latest

# 运行容器
docker run -d \
  --name boardcast \
  -p 8080:8080 \
  ghcr.io/yosebyte/boardcast:latest \
  --password "your-secure-password"
```

### 直接运行二进制文件

```bash
# 下载最新版本
wget https://github.com/yosebyte/boardcast/releases/latest/download/boardcast-linux-amd64.tar.gz

# 解压并运行
tar -xzf boardcast-linux-amd64.tar.gz
./boardcast --password "your-secure-password"
```

访问 `http://localhost:8080` 即可开始使用。

## 安装

### 二进制发布版本

从 [GitHub Releases](https://github.com/yosebyte/boardcast/releases) 页面下载适合您系统的预编译二进制文件。

支持的平台：
- **Linux**: amd64, arm64, arm, 386, mips 等
- **Windows**: amd64, arm64, 386
- **macOS**: amd64, arm64 (Apple Silicon)
- **FreeBSD**: amd64, arm64

### 从源码编译

```bash
# 克隆仓库
git clone https://github.com/yosebyte/boardcast.git
cd boardcast

# 编译
go mod download
go build -o boardcast ./cmd/boardcast

# 运行
./boardcast --password "your-secure-password"
```

### Docker 镜像

```bash
# 从 GitHub Container Registry 拉取
docker pull ghcr.io/yosebyte/boardcast:latest

# 或者自己构建
docker build -t boardcast .
```

## 配置

### 命令行参数

```bash
./boardcast [选项]
```

**可用选项：**

| 参数 | 类型 | 默认值 | 描述 |
|------|------|--------|------|
| `--password` | string | 无 | **必需** - 访问密码 |
| `--port` | string | `8080` | 服务器监听端口 (1-65535) |
| `--version` | bool | `false` | 显示版本信息并退出 |

**示例：**

```bash
# 基本使用
./boardcast --password "mypassword"

# 自定义端口
./boardcast --password "mypassword" --port 3000

# 查看版本
./boardcast --version
```

### 环境变量

目前版本仅支持命令行参数配置，未来版本将支持环境变量配置。

## 部署

### Docker Compose

创建 `docker-compose.yml` 文件：

```yaml
version: '3.8'

services:
  boardcast:
    image: ghcr.io/yosebyte/boardcast:latest
    container_name: boardcast
    ports:
      - "8080:8080"
    command: ["--password", "your-secure-password"]
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "wget", "--quiet", "--tries=1", "--spider", "http://localhost:8080"]
      interval: 30s
      timeout: 10s
      retries: 3
```

运行：

```bash
docker-compose up -d
```

### Kubernetes

创建 Kubernetes 部署文件：

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: boardcast
spec:
  replicas: 1
  selector:
    matchLabels:
      app: boardcast
  template:
    metadata:
      labels:
        app: boardcast
    spec:
      containers:
      - name: boardcast
        image: ghcr.io/yosebyte/boardcast:latest
        args: ["--password", "your-secure-password"]
        ports:
        - containerPort: 8080
        resources:
          requests:
            memory: "64Mi"
            cpu: "50m"
          limits:
            memory: "128Mi"
            cpu: "100m"
---
apiVersion: v1
kind: Service
metadata:
  name: boardcast-service
spec:
  selector:
    app: boardcast
  ports:
  - port: 80
    targetPort: 8080
  type: LoadBalancer
```

### 反向代理

#### Nginx

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

#### Caddy

```caddyfile
your-domain.com {
    reverse_proxy localhost:8080
}
```

## 开发

### 项目结构

```
boardcast/
├── cmd/boardcast/         # 应用程序入口点
│   └── main.go            # 主函数
├── internal/              # 内部包
│   ├── server.go          # HTTP 服务器
│   ├── auth/              # 认证模块
│   │   └── auth.go        # 认证管理器
│   ├── config/            # 配置管理
│   │   └── config.go      # 配置解析和验证
│   ├── handler/           # HTTP 处理器
│   │   └── handlers.go    # 路由处理函数
│   ├── template/          # HTML 模板
│   │   └── whiteboard.go  # 白板界面模板
│   └── websocket/         # WebSocket 管理
│       └── hub.go         # WebSocket 连接管理器
├── .github/workflows/     # GitHub Actions 工作流
├── Dockerfile             # Docker 构建文件
├── .goreleaser.yml        # GoReleaser 配置
├── go.mod                 # Go 模块定义
├── go.sum                 # Go 模块校验和
└── README.md              # 项目文档
```

### 核心组件

#### 1. 服务器 (`internal/server.go`)
- HTTP 服务器配置和生命周期管理
- 路由注册和中间件
- 优雅关闭处理

#### 2. 认证 (`internal/auth/`)
- 基于 bcrypt 的密码验证
- Session 管理
- 认证中间件

#### 3. WebSocket 管理 (`internal/websocket/`)
- 客户端连接管理
- 实时消息广播
- 内容同步

#### 4. 处理器 (`internal/handler/`)
- HTTP 路由处理
- 请求验证和响应
- 静态文件服务

### 技术栈

- **后端**: Go 1.25+
- **WebSocket**: Gorilla WebSocket
- **认证**: Gorilla Sessions + bcrypt
- **前端**: 原生 HTML/CSS/JavaScript
- **容器化**: Docker + 多阶段构建
- **CI/CD**: GitHub Actions

### 开发环境设置

```bash
# 1. 克隆仓库
git clone https://github.com/yosebyte/boardcast.git
cd boardcast

# 2. 安装依赖
go mod download

# 3. 运行开发服务器
go run cmd/boardcast/main.go --password "dev-password"

# 4. 访问应用
open http://localhost:8080
```

### 构建

```bash
# 本地构建
go build -o boardcast ./cmd/boardcast

# 交叉编译（Linux）
GOOS=linux GOARCH=amd64 go build -o boardcast-linux-amd64 ./cmd/boardcast

# Docker 构建
docker build -t boardcast .
```

### 测试

```bash
# 运行测试
go test ./...

# 运行测试并显示覆盖率
go test -cover ./...

# 生成测试覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## API 接口

### 路由列表

| 路径 | 方法 | 描述 | 认证 |
|------|------|------|------|
| `/` | GET | 白板主页面 | 否 |
| `/auth` | POST | 用户认证 | 否 |
| `/logout` | POST | 用户登出 | 是 |
| `/ws` | WebSocket | WebSocket 连接 | 是 |
| `/content` | GET | 获取当前内容 | 是 |

### 认证接口

#### POST /auth
认证用户身份

**请求体：**
```json
{
  "password": "your-password"
}
```

**响应：**
- `200 OK`: 认证成功
- `400 Bad Request`: 请求格式错误
- `401 Unauthorized`: 密码错误

#### POST /logout
用户登出

**响应：**
- `200 OK`: 登出成功

### WebSocket 接口

#### WebSocket /ws
实时通信连接

**连接要求：** 必须先通过 `/auth` 认证

**消息格式：** 纯文本

- **接收消息**: 其他用户的内容更新
- **发送消息**: 当前用户的内容更新

## 许可证

本项目使用 [BSD 3-Clause License](LICENSE) 许可证。
