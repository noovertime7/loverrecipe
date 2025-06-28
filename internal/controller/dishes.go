package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"loverrecipe/internal/domain"
	"loverrecipe/internal/response"
	"loverrecipe/internal/services/dishes"
)

type DishController struct {
	service dishes.Service
}

func NewDishControllerWithRegister(service dishes.Service) *DishController {
	controller := &DishController{
		service: service,
	}

	return controller
}

// CreateDishes 创建菜品
// @Summary 创建菜品
// @Description 创建新的菜品
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param dishes body domain.CreateDishesRequest true "菜品信息"
// @Success 200 {object} response.Response{data=domain.Dishes} "创建成功"
// @Failure 400 {object} response.Response{msg=string} "请求参数错误"
// @Failure 401 {object} response.Response{msg=string} "未授权"
// @Failure 500 {object} response.Response{msg=string} "服务器内部错误"
// @Router /api/v1/dishes [post]
func (c *DishController) CreateDishes(ctx *gin.Context) {
	var req domain.CreateDishesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "请求参数错误: "+err.Error())
		return
	}

	// 从JWT中获取用户ID（这里需要根据您的JWT实现调整）
	userID := c.getUserIDFromContext(ctx)
	req.UserID = userID

	dishes, err := c.service.CreateDishes(ctx.Request.Context(), req)
	if err != nil {
		response.AppErrorResponse(ctx, err)
		return
	}

	response.SuccessWithMsg(ctx, "创建成功", dishes)
}

// GetDishesByID 根据ID获取菜品详情
// @Summary 获取菜品详情
// @Description 根据菜品ID获取菜品详细信息
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "菜品ID"
// @Success 200 {object} response.Response{data=domain.Dishes} "获取成功"
// @Failure 400 {object} response.Response{msg=string} "请求参数错误"
// @Failure 404 {object} response.Response{msg=string} "菜品不存在"
// @Failure 500 {object} response.Response{msg=string} "服务器内部错误"
// @Router /api/v1/dishes/{id} [get]
func (c *DishController) GetDishesByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "无效的菜品ID")
		return
	}

	dishes, err := c.service.GetDishesByID(ctx.Request.Context(), id)
	if err != nil {
		if err == domain.ErrDishesNotFound {
			response.DishNotFound(ctx)
			return
		}
		response.AppErrorResponse(ctx, err)
		return
	}

	response.Success(ctx, dishes)
}

// UpdateDishes 更新菜品
// @Summary 更新菜品
// @Description 更新菜品信息
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "菜品ID"
// @Param dishes body domain.UpdateDishesRequest true "菜品更新信息"
// @Success 200 {object} response.Response{data=domain.Dishes} "更新成功"
// @Failure 400 {object} response.Response{msg=string} "请求参数错误"
// @Failure 401 {object} response.Response{msg=string} "未授权"
// @Failure 404 {object} response.Response{msg=string} "菜品不存在"
// @Failure 500 {object} response.Response{msg=string} "服务器内部错误"
// @Router /api/v1/dishes/{id} [put]
func (c *DishController) UpdateDishes(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "无效的菜品ID")
		return
	}

	var req domain.UpdateDishesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "请求参数错误: "+err.Error())
		return
	}

	req.ID = id
	req.UserID = c.getUserIDFromContext(ctx)

	dishes, err := c.service.UpdateDishes(ctx.Request.Context(), req)
	if err != nil {
		if err == domain.ErrDishesNotFound {
			response.DishNotFound(ctx)
			return
		}
		if err == domain.ErrDishesUserMismatch {
			response.DishUserMismatch(ctx)
			return
		}
		response.AppErrorResponse(ctx, err)
		return
	}

	response.SuccessWithMsg(ctx, "更新成功", dishes)
}

// DeleteDishes 删除菜品
// @Summary 删除菜品
// @Description 删除指定菜品
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "菜品ID"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Failure 400 {object} response.Response{msg=string} "请求参数错误"
// @Failure 401 {object} response.Response{msg=string} "未授权"
// @Failure 403 {object} response.Response{msg=string} "无权限"
// @Failure 404 {object} response.Response{msg=string} "菜品不存在"
// @Failure 500 {object} response.Response{msg=string} "服务器内部错误"
// @Router /api/v1/dishes/{id} [delete]
func (c *DishController) DeleteDishes(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "无效的菜品ID")
		return
	}

	userID := c.getUserIDFromContext(ctx)
	err = c.service.DeleteDishes(ctx.Request.Context(), id, userID)
	if err != nil {
		if err == domain.ErrDishesNotFound {
			response.DishNotFound(ctx)
			return
		}
		if err == domain.ErrDishesUserMismatch {
			response.DishUserMismatch(ctx)
			return
		}
		response.AppErrorResponse(ctx, err)
		return
	}

	response.SuccessWithMsg(ctx, "删除成功", nil)
}

// ListDishes 获取菜品列表
// @Summary 获取菜品列表
// @Description 分页获取当前用户的菜品列表
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码，默认1"
// @Param size query int false "每页数量，默认10，最大100"
// @Param type query int false "菜品种类ID"
// @Success 200 {object} response.Response{data=domain.DishesListResponse} "获取成功"
// @Failure 400 {object} response.Response{msg=string} "请求参数错误"
// @Failure 401 {object} response.Response{msg=string} "未授权"
// @Failure 500 {object} response.Response{msg=string} "服务器内部错误"
// @Router /api/v1/dishes [get]
func (c *DishController) ListDishes(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	typeID, _ := strconv.ParseInt(ctx.Query("type"), 10, 64)

	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	if size > 100 {
		size = 100
	}

	offset := (page - 1) * size
	userID := c.getUserIDFromContext(ctx)

	query := domain.DishesQuery{
		UserID: userID,
		Type:   typeID,
		Offset: offset,
		Limit:  size,
	}

	result, err := c.service.ListDishes(ctx.Request.Context(), query)
	if err != nil {
		response.AppErrorResponse(ctx, err)
		return
	}

	response.Success(ctx, result)
}

// SearchDishes 搜索菜品
// @Summary 搜索菜品
// @Description 根据关键词搜索菜品
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param keyword query string true "搜索关键词"
// @Param page query int false "页码，默认1"
// @Param size query int false "每页数量，默认10，最大100"
// @Success 200 {object} response.Response{data=domain.DishesListResponse} "搜索成功"
// @Failure 400 {object} response.Response{msg=string} "请求参数错误"
// @Failure 401 {object} response.Response{msg=string} "未授权"
// @Failure 500 {object} response.Response{msg=string} "服务器内部错误"
// @Router /api/v1/dishes/search [get]
func (c *DishController) SearchDishes(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	if size > 100 {
		size = 100
	}

	offset := (page - 1) * size
	userID := c.getUserIDFromContext(ctx)

	result, err := c.service.SearchDishes(ctx.Request.Context(), userID, keyword, offset, size)
	if err != nil {
		response.AppErrorResponse(ctx, err)
		return
	}

	response.SuccessWithMsg(ctx, "搜索成功", result)
}

// GetDishesStatistics 获取菜品统计
// @Summary 获取菜品统计
// @Description 获取当前用户的菜品统计信息
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} response.Response{data=dishes.DishesStatistics} "获取成功"
// @Failure 401 {object} response.Response{msg=string} "未授权"
// @Failure 500 {object} response.Response{msg=string} "服务器内部错误"
// @Router /api/v1/dishes/statistics [get]
func (c *DishController) GetDishesStatistics(ctx *gin.Context) {
	userID := c.getUserIDFromContext(ctx)

	stats, err := c.service.GetDishesStatistics(ctx.Request.Context(), userID)
	if err != nil {
		response.AppErrorResponse(ctx, err)
		return
	}

	response.Success(ctx, stats)
}

// GetDishesByType 按种类获取菜品
// @Summary 按种类获取菜品
// @Description 根据菜品种类ID获取菜品列表
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param typeId path int true "菜品种类ID"
// @Success 200 {object} response.Response{data=[]domain.Dishes} "获取成功"
// @Failure 400 {object} response.Response{msg=string} "请求参数错误"
// @Failure 401 {object} response.Response{msg=string} "未授权"
// @Failure 500 {object} response.Response{msg=string} "服务器内部错误"
// @Router /api/v1/dishes/type/{typeId} [get]
func (c *DishController) GetDishesByType(ctx *gin.Context) {
	typeIDStr := ctx.Param("typeId")
	typeID, err := strconv.ParseInt(typeIDStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "无效的菜品种类ID")
		return
	}

	dishes, err := c.service.GetDishesByType(ctx.Request.Context(), typeID)
	if err != nil {
		response.AppErrorResponse(ctx, err)
		return
	}

	response.Success(ctx, dishes)
}

// GetDishesWithTypeInfo 获取带种类信息的菜品
// @Summary 获取带种类信息的菜品
// @Description 获取当前用户的菜品列表，包含种类详细信息
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} response.Response{data=[]domain.DishesWithType} "获取成功"
// @Failure 401 {object} response.Response{msg=string} "未授权"
// @Failure 500 {object} response.Response{msg=string} "服务器内部错误"
// @Router /api/v1/dishes/with-type [get]
func (c *DishController) GetDishesWithTypeInfo(ctx *gin.Context) {
	userID := c.getUserIDFromContext(ctx)

	dishesWithType, err := c.service.GetDishesWithTypeInfo(ctx.Request.Context(), userID)
	if err != nil {
		response.AppErrorResponse(ctx, err)
		return
	}

	response.Success(ctx, dishesWithType)
}

// getUserIDFromContext 从上下文中获取用户ID
// 这里需要根据您的JWT实现来调整
func (c *DishController) getUserIDFromContext(ctx *gin.Context) int64 {
	// 示例实现，需要根据您的JWT中间件调整
	// 假设JWT中间件将用户ID存储在上下文中
	if userID, exists := ctx.Get("user_id"); exists {
		if id, ok := userID.(int64); ok {
			return id
		}
	}
	// 临时返回默认值，实际项目中应该从JWT中解析
	return 1
}
