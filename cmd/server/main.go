// Package main 用户食谱管理系统 API 服务
// @title 用户食谱管理系统 API
// @version 1.0
// @description 这是一个用户食谱管理系统的后端 API 服务
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description 请输入 "Bearer " 加上 JWT token
package main

import (
	"context"
	"loverrecipe/cmd/server/ioc"
	_ "loverrecipe/docs"
	ioc2 "loverrecipe/internal/ioc"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server"
	"github.com/gotomicro/ego/server/egovernor"
	"go.opentelemetry.io/otel/sdk/trace"
)

// @title 用户食谱管理系统 API
// @version 1.0
// @description 用户食谱管理系统的后端 API 服务
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description 请输入 "Bearer " 加上 JWT token
func main() {
	// 创建 ego 应用实例
	egoApp := ego.New()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tp := ioc2.InitZipkinTracer()
	defer func(tp *trace.TracerProvider, ctx context.Context) {
		err := tp.Shutdown(ctx)
		if err != nil {
			elog.Error("Shutdown zipkinTracer", elog.FieldErr(err))
		}
	}(tp, ctx)

	app := ioc.InitHttpServer()

	// 启动服务
	if err := egoApp.Serve(
		egovernor.Load("server.governor").Build(),
		func() server.Server {
			return app.HttpServer
		}(),
	).Cron(app.Crons...).
		Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}
