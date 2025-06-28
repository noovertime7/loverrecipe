package main

import (
	"context"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server"
	"github.com/gotomicro/ego/server/egovernor"
	"go.opentelemetry.io/otel/sdk/trace"
	"loverrecipe/cmd/server/ioc"
	ioc2 "loverrecipe/internal/ioc"
)

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
