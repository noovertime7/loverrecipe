//go:build wireinject

package ioc

import (
	"github.com/google/wire"
	"loverrecipe/internal/controller"
	"loverrecipe/internal/ioc"
	"loverrecipe/internal/repository"
	"loverrecipe/internal/repository/dao"
	"loverrecipe/internal/services/dishes"
	"loverrecipe/internal/services/user"
	"loverrecipe/internal/token"
)

var (
	BaseSet = wire.NewSet(
		ioc.InitDB,
		ioc.InitRedisCmd,
		ioc.InitRedisClient,
		ioc.InitIDGenerator,
		token.RegisterJwt,
	)
	dishesSet = wire.NewSet(
		dao.NewDishesDao,
		repository.NewDishesRepository,
		dishes.NewService,
		controller.NewDishControllerWithRegister,
	)
	userSet = wire.NewSet(
		dao.NewUserDao,
		repository.NewUserRepository,
		user.NewService,
		controller.NewUserController,
	)
)

func InitHttpServer() *ioc.App {
	wire.Build(
		BaseSet,
		dishesSet,
		userSet,
		ioc.Crons,
		ioc.InitHTTP,
		ioc.InitTasks,
		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
