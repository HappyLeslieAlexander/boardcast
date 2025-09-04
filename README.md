<div align="center">
  <img src="https://cdn.yobc.de/assets/boardcast.png" alt="boardcast" width="300">
</div>

<div align="center">
  <p><strong>轻量极简的实时协作文字白板</strong></p>
</div>

<div align="center">
  <a href="#✨-特性">特性</a> •
  <a href="#📦-安装">安装</a> •
  <a href="#⚙️-配置">配置</a> •
  <a href="#🛠️-开发">开发</a>
</div>

## 📖 概述

BoardCast 是一个轻量级的实时协作白板应用，使用 Go 语言开发。它允许多个用户通过密码认证后在同一个白板上实时协作编辑文本内容。应用具有简洁的用户界面，支持 WebSocket 实时同步，内置快照功能可随时保存和恢复内容，是团队协作、会议记录、头脑风暴的理想工具。

## 🎬 演示

<div align="center">
  <img src="https://cdn.yobc.de/assets/boardcast.gif" alt="BoardCast" width="1280">
</div>

## ✨ 特性

### 🔐 安全认证
- 基于密码的访问控制
- 会话管理和安全的 Cookie 存储
- bcrypt 密码加密
- 可选生成随机密码

### 🔄 实时协作
- WebSocket 实时通信
- 多用户同步编辑
- 异常断线自动重连
- 自动内容保存和恢复

### 📸 快照管理
- 一键保存白板内容快照
- 快速恢复历史快照
- 本地文件持久化存储
- 支持覆盖更新

### 📱 响应式设计
- 适配桌面和移动设备
- 简洁直观的用户界面
- 现代化的设计风格
- 支持暗色主题

### 🚀 轻量高效
- 单文件部署
- 开箱即用
- 低资源占用
- 无数据库依赖

### 🛠️ 运维友好
- Docker 容器化支持
- 多平台二进制发布
- 优雅关闭和错误处理
- 结构化日志

## 📦 安装

### 📋 二进制发布版本

从 [GitHub Releases](https://github.com/yosebyte/boardcast/releases) 页面下载适合您系统的预编译二进制文件。

📱 **支持的平台**：
- **🐧 Linux**: amd64, arm64, arm, 386, mips 等
- **🪟 Windows**: amd64, arm64, 386
- **🍎 macOS**: amd64, arm64 (Apple Silicon)
- **🔥 FreeBSD**: amd64, arm64

### 🔧 从源码编译

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

### 🐳 Docker 镜像

```bash
# 从 GitHub Container Registry 拉取
docker pull ghcr.io/yosebyte/boardcast:latest

# 或者自己构建
docker build -t boardcast .
```

### 📝 使用指南

1. **🌐 访问应用**: 在浏览器中打开应用地址
2. **🔑 身份认证**: 在密码框输入密码进行登录
3. **✏️ 开始协作**: 在白板区域输入和编辑文本内容
4. **💾 保存快照**: 点击保存按钮保存当前内容到本地文件
5. **🔄 恢复快照**: 点击恢复按钮从快照文件恢复内容
6. **🎨 主题切换**: 点击切换按钮切换明暗主题
7. **🚪 断开连接**: 点击退出按钮登出并断开连接

**📊 连接状态指示**：
- **🔴 红色密码框**: 未连接状态，需要输入密码进行认证
- **🟡 黄色密码框**: 连接中状态，正在建立WebSocket连接
- **🟢 绿色密码框**: 已连接状态，可以正常进行实时协作

**📸 快照功能说明**：
- 快照文件保存为 `boardcast.txt`，位于应用运行目录
- 保存快照会覆盖之前的快照文件
- 恢复快照会将内容同步到所有在线用户

## ⚙️ 配置

### 🚀 命令行参数

```bash
./boardcast [选项]
```

**📝 可用选项：**

| 参数 | 类型 | 默认值 | 描述 |
|------|------|--------|------|
| `--password` | string | 随机生成 | 🔐 访问密码 |
| `--port` | string | `8200` | 🌐 服务器监听端口 |
| `--version` | bool | `false` | ℹ️ 显示版本信息并退出 |

**💡 示例：**

```bash
# 使用随机密码
./boardcast

# 自定义密码和端口
./boardcast --password "secret" --port 3000

# 查看版本
./boardcast --version
```

## 🛠️ 开发

### 📁 项目结构

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
│   ├── docker.yml         # Docker 镜像构建工作流
│   └── release.yml        # 二进制包发布工作流
├── Dockerfile             # Docker 构建文件
├── .goreleaser.yml        # GoReleaser 配置
├── go.mod                 # Go 模块定义
├── go.sum                 # Go 模块校验和
└── README.md              # 项目文档
```

### 🧩 核心组件

#### 1. 🌐 服务器 (`internal/server.go`)
- HTTP 服务器配置和生命周期管理
- 路由注册和中间件
- 优雅关闭处理

#### 2. 🔐 认证 (`internal/auth/`)
- 基于 bcrypt 的密码验证
- Session 管理
- 认证中间件

#### 3. 🔌 WebSocket 管理 (`internal/websocket/`)
- 客户端连接管理
- 实时消息广播
- 内容同步
- 快照保存和恢复功能

#### 4. 🎯 处理器 (`internal/handler/`)
- HTTP 路由处理
- 请求验证和响应
- 静态文件服务
- 快照API端点处理

### 🛠️ 技术栈

- **🐹 后端**: Go 1.25+
- **🔌 WebSocket**: Gorilla WebSocket
- **🔐 认证**: Gorilla Sessions + bcrypt
- **🎨 前端**: 原生 HTML/CSS/JavaScript
- **🐳 容器化**: Docker + 多阶段构建
- **🔄 CI/CD**: GitHub Actions

### 💻 开发环境设置

```bash
# 1. 克隆仓库
git clone https://github.com/yosebyte/boardcast.git
cd boardcast

# 2. 安装依赖
go mod download

# 3. 运行开发服务器
go run cmd/boardcast/main.go --password "dev-password"

# 4. 访问应用
open http://localhost:8200
```

### 🔨 构建

```bash
# 本地构建
go build -o boardcast ./cmd/boardcast

# 交叉编译（Linux）
GOOS=linux GOARCH=amd64 go build -o boardcast-linux-amd64 ./cmd/boardcast

# Docker 构建
docker build -t boardcast .
```

## 🌐 API 接口

| 路径 | 方法 | 描述 | 认证 |
|------|------|------|------|
| `/` | GET | 🏠 白板主页面 | 否 |
| `/auth` | POST | 🔐 用户认证 | 否 |
| `/logout` | POST | 🚪 用户登出 | 是 |
| `/ws` | WebSocket | 🔌 WebSocket 连接 | 是 |
| `/content` | GET | 📄 获取当前内容 | 是 |
| `/save` | POST | 💾 保存内容快照 | 是 |
| `/restore` | POST | 🔄 恢复内容快照 | 是 |

## 📄 许可证

本项目使用 [BSD 3-Clause License](LICENSE) 许可证
