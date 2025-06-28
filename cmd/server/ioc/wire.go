//go:build wireinject

package ioc

import (
	"github.com/google/wire"
	"loverrecipe/internal/controller"
	"loverrecipe/internal/ioc"
	"loverrecipe/internal/repository"
	"loverrecipe/internal/repository/dao"
	"loverrecipe/internal/services/dishes"
)

var (
	BaseSet = wire.NewSet(
		ioc.InitDB,
		ioc.InitRedisCmd,
		ioc.InitRedisClient,
		ioc.InitIDGenerator,
	)
	dishesSet = wire.NewSet(
		dao.NewDishesDao,
		repository.NewDishesRepository,
		dishes.NewService,
		controller.NewDishControllerWithRegister,
	)
)

func InitHttpServer() *ioc.App {
	wire.Build(
		BaseSet,
		dishesSet,
		ioc.Crons,
		ioc.InitHTTP,
		ioc.InitTasks,
		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
