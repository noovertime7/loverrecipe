package ioc

import (
	"github.com/gotomicro/ego/server/egin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"loverrecipe/internal/controller"
)

func InitHTTP(d *controller.DishController, user *controller.UserController) *egin.Component {
	server := egin.Load("server.http").Build()
	// 添加 Swagger 路由
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	dishesGroup := server.Group("/api/v1/dishes")
	{
		// 创建菜品
		dishesGroup.POST("", d.CreateDishes)

		// 获取菜品详情
		dishesGroup.GET("/:id", d.GetDishesByID)

		// 更新菜品
		dishesGroup.PUT("/:id", d.UpdateDishes)

		// 删除菜品
		dishesGroup.DELETE("/:id", d.DeleteDishes)

		// 获取菜品列表
		dishesGroup.GET("", d.ListDishes)

		// 搜索菜品
		dishesGroup.GET("/search", d.SearchDishes)

		// 获取菜品统计
		dishesGroup.GET("/statistics", d.GetDishesStatistics)

		// 按种类获取菜品
		dishesGroup.GET("/type/:typeId", d.GetDishesByType)

		// 获取带种类信息的菜品
		dishesGroup.GET("/with-type", d.GetDishesWithTypeInfo)
	}

	{
		// 用户注册
		usersGroup := server.Group("/api/v1/user")
		usersGroup.POST("/register", user.Register)
	}

	return server
}
