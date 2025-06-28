package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应格式
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  GetErrorMessage(CodeSuccess),
		Data: data,
	})
}

// SuccessWithMsg 带自定义消息的成功响应
func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  msg,
		Data: data,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  GetErrorMessage(code),
	})
}

// ErrorWithMsg 带自定义消息的错误响应
func ErrorWithMsg(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
	})
}

// ErrorWithData 带数据的错误响应
func ErrorWithData(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  GetErrorMessage(code),
		Data: data,
	})
}

// AppErrorResponse 应用错误响应
func AppErrorResponse(c *gin.Context, err error) {
	if appErr, ok := IsAppError(err); ok {
		c.JSON(http.StatusOK, Response{
			Code: appErr.Code,
			Msg:  appErr.Message,
		})
		return
	}

	// 如果不是应用错误，返回通用错误
	c.JSON(http.StatusOK, Response{
		Code: CodeInternalServerError,
		Msg:  err.Error(),
	})
}

// BadRequest 请求参数错误
func BadRequest(c *gin.Context, msg ...string) {
	message := GetErrorMessage(CodeBadRequest)
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(http.StatusOK, Response{
		Code: CodeBadRequest,
		Msg:  message,
	})
}

// Unauthorized 未授权
func Unauthorized(c *gin.Context, msg ...string) {
	message := GetErrorMessage(CodeUnauthorized)
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(http.StatusOK, Response{
		Code: CodeUnauthorized,
		Msg:  message,
	})
}

// Forbidden 禁止访问
func Forbidden(c *gin.Context, msg ...string) {
	message := GetErrorMessage(CodeForbidden)
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(http.StatusOK, Response{
		Code: CodeForbidden,
		Msg:  message,
	})
}

// NotFound 资源不存在
func NotFound(c *gin.Context, msg ...string) {
	message := GetErrorMessage(CodeNotFound)
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(http.StatusOK, Response{
		Code: CodeNotFound,
		Msg:  message,
	})
}

// InternalServerError 服务器内部错误
func InternalServerError(c *gin.Context, msg ...string) {
	message := GetErrorMessage(CodeInternalServerError)
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(http.StatusOK, Response{
		Code: CodeInternalServerError,
		Msg:  message,
	})
}

// ValidationError 验证错误
func ValidationError(c *gin.Context, field string, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: CodeUnprocessableEntity,
		Msg:  field + ": " + msg,
	})
}

// DatabaseError 数据库错误
func DatabaseError(c *gin.Context, msg ...string) {
	message := GetErrorMessage(CodeDatabaseError)
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(http.StatusOK, Response{
		Code: CodeDatabaseError,
		Msg:  message,
	})
}

// CacheError 缓存错误
func CacheError(c *gin.Context, msg ...string) {
	message := GetErrorMessage(CodeCacheError)
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(http.StatusOK, Response{
		Code: CodeCacheError,
		Msg:  message,
	})
}

// ThirdPartyError 第三方服务错误
func ThirdPartyError(c *gin.Context, msg ...string) {
	message := GetErrorMessage(CodeThirdPartyError)
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(http.StatusOK, Response{
		Code: CodeThirdPartyError,
		Msg:  message,
	})
}

// FileUploadError 文件上传错误
func FileUploadError(c *gin.Context, msg ...string) {
	message := GetErrorMessage(CodeFileUploadFailed)
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(http.StatusOK, Response{
		Code: CodeFileUploadFailed,
		Msg:  message,
	})
}

// UserNotFound 用户不存在
func UserNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeUserNotFound,
		Msg:  GetErrorMessage(CodeUserNotFound),
	})
}

// UserAlreadyExists 用户已存在
func UserAlreadyExists(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeUserAlreadyExists,
		Msg:  GetErrorMessage(CodeUserAlreadyExists),
	})
}

// InvalidCredentials 无效凭据
func InvalidCredentials(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeInvalidCredentials,
		Msg:  GetErrorMessage(CodeInvalidCredentials),
	})
}

// TokenExpired 令牌过期
func TokenExpired(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeTokenExpired,
		Msg:  GetErrorMessage(CodeTokenExpired),
	})
}

// TokenInvalid 令牌无效
func TokenInvalid(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeTokenInvalid,
		Msg:  GetErrorMessage(CodeTokenInvalid),
	})
}

// PermissionDenied 权限不足
func PermissionDenied(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodePermissionDenied,
		Msg:  GetErrorMessage(CodePermissionDenied),
	})
}

// DishNotFound 菜品不存在
func DishNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeDishNotFound,
		Msg:  GetErrorMessage(CodeDishNotFound),
	})
}

// DishNameEmpty 菜品名称不能为空
func DishNameEmpty(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeDishNameEmpty,
		Msg:  GetErrorMessage(CodeDishNameEmpty),
	})
}

// DishPriceInvalid 菜品价格无效
func DishPriceInvalid(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeDishPriceInvalid,
		Msg:  GetErrorMessage(CodeDishPriceInvalid),
	})
}

// DishTypeInvalid 菜品类型无效
func DishTypeInvalid(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeDishTypeInvalid,
		Msg:  GetErrorMessage(CodeDishTypeInvalid),
	})
}

// DishUserMismatch 菜品不属于该用户
func DishUserMismatch(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeDishUserMismatch,
		Msg:  GetErrorMessage(CodeDishUserMismatch),
	})
}

// DishTypeNotFound 菜品种类不存在
func DishTypeNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeDishTypeNotFound,
		Msg:  GetErrorMessage(CodeDishTypeNotFound),
	})
}

// DishTypeNameEmpty 菜品种类名称不能为空
func DishTypeNameEmpty(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeDishTypeNameEmpty,
		Msg:  GetErrorMessage(CodeDishTypeNameEmpty),
	})
}

// DishTypeUserMismatch 菜品种类不属于该用户
func DishTypeUserMismatch(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeDishTypeUserMismatch,
		Msg:  GetErrorMessage(CodeDishTypeUserMismatch),
	})
}

// FileTypeNotAllowed 文件类型不允许
func FileTypeNotAllowed(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeFileTypeNotAllowed,
		Msg:  GetErrorMessage(CodeFileTypeNotAllowed),
	})
}

// FileSizeExceeded 文件大小超限
func FileSizeExceeded(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeFileSizeExceeded,
		Msg:  GetErrorMessage(CodeFileSizeExceeded),
	})
}

// FileNotFound 文件不存在
func FileNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeFileNotFound,
		Msg:  GetErrorMessage(CodeFileNotFound),
	})
}

// TooManyRequests 请求过于频繁
func TooManyRequests(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeTooManyRequests,
		Msg:  GetErrorMessage(CodeTooManyRequests),
	})
}

// ServiceUnavailable 服务不可用
func ServiceUnavailable(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeServiceUnavailable,
		Msg:  GetErrorMessage(CodeServiceUnavailable),
	})
}
