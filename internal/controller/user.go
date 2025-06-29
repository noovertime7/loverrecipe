package controller

import (
	"loverrecipe/internal/domain"
	"loverrecipe/internal/response"
	"loverrecipe/internal/services/user"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/elog"
)

type UserController struct {
	user.Service // 嵌入用户服务接口，直接使用服务层的方法
}

func NewUserController(s user.Service) *UserController {
	return &UserController{s}
}

// Register 用户注册接口
// @Summary 用户注册
// @Description 处理用户注册请求，验证输入参数并创建新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body domain.CreateUserInput true "用户注册信息"
// @Success 200 {object} response.Response{data=domain.CreateUserOutput} "注册成功"
// @Failure 400 {object} response.Response{msg=string} "请求参数错误"
// @Failure 500 {object} response.Response{msg=string} "服务器内部错误"
// @Router /api/v1/user/register [post]
func (uc *UserController) Register(ctx *gin.Context) {
	// 创建用户注册参数结构体
	params := &domain.CreateUserInput{}

	// 绑定并验证 JSON 请求体到参数结构体
	if err := domain.BindJson(ctx, params); err != nil {
		// 参数绑定失败，返回 400 错误
		response.BadRequest(ctx, err.Error())
		// 记录错误日志
		elog.Error("bind json error", elog.String("error", err.Error()))
		return
	}

	// 调用服务层创建用户
	data, err := uc.Service.Create(ctx.Request.Context(), params)
	if err != nil {
		// 创建用户失败，返回 500 错误
		response.InternalServerError(ctx, err.Error())
		// 记录错误日志
		elog.Error("create user error", elog.String("error", err.Error()))
		return
	}

	// 注册成功，返回用户信息
	response.Success(ctx, data)
}
