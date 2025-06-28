//go:build wireinject

package ioc

import (
	"github.com/google/wire"
	"loverrecipe/internal/ioc"
)

var (
	BaseSet = wire.NewSet(
		ioc.InitDB,
		ioc.InitRedisCmd,
		ioc.InitRedisClient,
		ioc.InitIDGenerator,
	)
)

func InitHttpServer() *ioc.App {
	wire.Build(
		//BaseSet,
		ioc.Crons,
		ioc.InitHTTP,
		ioc.InitTasks,
		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
