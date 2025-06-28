# loverrecipe

## 技术栈
- Go 1.20
- gin（HTTP框架）
- wire（依赖注入）
- ego（微服务框架）

## 目录结构
```
cmd/server         # 程序入口
internal/handler   # 路由与handler
internal/service   # 业务逻辑
internal/repository# 数据访问
internal/model     # 数据模型
internal/config    # 配置
pkg/               # 公共包
wire/              # wire依赖注入相关
```

## 安装依赖
```
go mod tidy
```

## 生成依赖注入代码
```
wire gen ./wire
```

## 启动服务
```
go run ./cmd/server/main.go
```

## 访问测试
浏览器或curl访问：http://localhost:9001/ping
返回：
```
{"message": "pong"}
```

## 开发建议
- 新增依赖注入请在对应ProviderSet中注册
- 业务逻辑建议分层（handler/service/repository）
- 配置建议放在internal/config 