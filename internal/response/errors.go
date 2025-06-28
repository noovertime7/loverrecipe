package response

import "errors"

// 错误码定义
const (
	// 成功
	CodeSuccess = 200

	// 客户端错误 (4xx)
	CodeBadRequest          = 400 // 请求参数错误
	CodeUnauthorized        = 401 // 未授权
	CodeForbidden           = 403 // 禁止访问
	CodeNotFound            = 404 // 资源不存在
	CodeMethodNotAllowed    = 405 // 方法不允许
	CodeConflict            = 409 // 资源冲突
	CodeUnprocessableEntity = 422 // 请求格式正确但语义错误
	CodeTooManyRequests     = 429 // 请求过于频繁

	// 服务器错误 (5xx)
	CodeInternalServerError = 500 // 服务器内部错误
	CodeNotImplemented      = 501 // 功能未实现
	CodeBadGateway          = 502 // 网关错误
	CodeServiceUnavailable  = 503 // 服务不可用

	// 业务错误码 (1000-9999)
	CodeUserNotFound       = 1001 // 用户不存在
	CodeUserAlreadyExists  = 1002 // 用户已存在
	CodeInvalidCredentials = 1003 // 无效凭据
	CodeTokenExpired       = 1004 // 令牌过期
	CodeTokenInvalid       = 1005 // 令牌无效
	CodePermissionDenied   = 1006 // 权限不足

	// 菜品相关错误码 (2000-2999)
	CodeDishNotFound     = 2001 // 菜品不存在
	CodeDishNameEmpty    = 2002 // 菜品名称不能为空
	CodeDishPriceInvalid = 2003 // 菜品价格无效
	CodeDishTypeInvalid  = 2004 // 菜品类型无效
	CodeDishUserMismatch = 2005 // 菜品不属于该用户

	// 菜品种类相关错误码 (3000-3999)
	CodeDishTypeNotFound     = 3001 // 菜品种类不存在
	CodeDishTypeNameEmpty    = 3002 // 菜品种类名称不能为空
	CodeDishTypeUserMismatch = 3003 // 菜品种类不属于该用户

	// 文件上传相关错误码 (4000-4999)
	CodeFileUploadFailed   = 4001 // 文件上传失败
	CodeFileTypeNotAllowed = 4002 // 文件类型不允许
	CodeFileSizeExceeded   = 4003 // 文件大小超限
	CodeFileNotFound       = 4004 // 文件不存在

	// 数据库相关错误码 (5000-5999)
	CodeDatabaseError      = 5001 // 数据库错误
	CodeDatabaseConnection = 5002 // 数据库连接错误
	CodeDatabaseTimeout    = 5003 // 数据库超时

	// 缓存相关错误码 (6000-6999)
	CodeCacheError      = 6001 // 缓存错误
	CodeCacheConnection = 6002 // 缓存连接错误
	CodeCacheTimeout    = 6003 // 缓存超时

	// 第三方服务错误码 (7000-7999)
	CodeThirdPartyError       = 7001 // 第三方服务错误
	CodeThirdPartyTimeout     = 7002 // 第三方服务超时
	CodeThirdPartyUnavailable = 7003 // 第三方服务不可用
)

// 错误信息映射
var errorMessages = map[int]string{
	// 成功
	CodeSuccess: "成功",

	// 客户端错误
	CodeBadRequest:          "请求参数错误",
	CodeUnauthorized:        "未授权",
	CodeForbidden:           "禁止访问",
	CodeNotFound:            "资源不存在",
	CodeMethodNotAllowed:    "方法不允许",
	CodeConflict:            "资源冲突",
	CodeUnprocessableEntity: "请求格式正确但语义错误",
	CodeTooManyRequests:     "请求过于频繁",

	// 服务器错误
	CodeInternalServerError: "服务器内部错误",
	CodeNotImplemented:      "功能未实现",
	CodeBadGateway:          "网关错误",
	CodeServiceUnavailable:  "服务不可用",

	// 业务错误
	CodeUserNotFound:       "用户不存在",
	CodeUserAlreadyExists:  "用户已存在",
	CodeInvalidCredentials: "无效凭据",
	CodeTokenExpired:       "令牌过期",
	CodeTokenInvalid:       "令牌无效",
	CodePermissionDenied:   "权限不足",

	// 菜品相关错误
	CodeDishNotFound:     "菜品不存在",
	CodeDishNameEmpty:    "菜品名称不能为空",
	CodeDishPriceInvalid: "菜品价格无效",
	CodeDishTypeInvalid:  "菜品类型无效",
	CodeDishUserMismatch: "菜品不属于该用户",

	// 菜品种类相关错误
	CodeDishTypeNotFound:     "菜品种类不存在",
	CodeDishTypeNameEmpty:    "菜品种类名称不能为空",
	CodeDishTypeUserMismatch: "菜品种类不属于该用户",

	// 文件上传相关错误
	CodeFileUploadFailed:   "文件上传失败",
	CodeFileTypeNotAllowed: "文件类型不允许",
	CodeFileSizeExceeded:   "文件大小超限",
	CodeFileNotFound:       "文件不存在",

	// 数据库相关错误
	CodeDatabaseError:      "数据库错误",
	CodeDatabaseConnection: "数据库连接错误",
	CodeDatabaseTimeout:    "数据库超时",

	// 缓存相关错误
	CodeCacheError:      "缓存错误",
	CodeCacheConnection: "缓存连接错误",
	CodeCacheTimeout:    "缓存超时",

	// 第三方服务错误
	CodeThirdPartyError:       "第三方服务错误",
	CodeThirdPartyTimeout:     "第三方服务超时",
	CodeThirdPartyUnavailable: "第三方服务不可用",
}

// 自定义错误类型
var (
	ErrUserNotFound       = NewError(CodeUserNotFound, "用户不存在")
	ErrUserAlreadyExists  = NewError(CodeUserAlreadyExists, "用户已存在")
	ErrInvalidCredentials = NewError(CodeInvalidCredentials, "无效凭据")
	ErrTokenExpired       = NewError(CodeTokenExpired, "令牌过期")
	ErrTokenInvalid       = NewError(CodeTokenInvalid, "令牌无效")
	ErrPermissionDenied   = NewError(CodePermissionDenied, "权限不足")

	ErrDishNotFound     = NewError(CodeDishNotFound, "菜品不存在")
	ErrDishNameEmpty    = NewError(CodeDishNameEmpty, "菜品名称不能为空")
	ErrDishPriceInvalid = NewError(CodeDishPriceInvalid, "菜品价格无效")
	ErrDishTypeInvalid  = NewError(CodeDishTypeInvalid, "菜品类型无效")
	ErrDishUserMismatch = NewError(CodeDishUserMismatch, "菜品不属于该用户")

	ErrDishTypeNotFound     = NewError(CodeDishTypeNotFound, "菜品种类不存在")
	ErrDishTypeNameEmpty    = NewError(CodeDishTypeNameEmpty, "菜品种类名称不能为空")
	ErrDishTypeUserMismatch = NewError(CodeDishTypeUserMismatch, "菜品种类不属于该用户")

	ErrFileUploadFailed   = NewError(CodeFileUploadFailed, "文件上传失败")
	ErrFileTypeNotAllowed = NewError(CodeFileTypeNotAllowed, "文件类型不允许")
	ErrFileSizeExceeded   = NewError(CodeFileSizeExceeded, "文件大小超限")
	ErrFileNotFound       = NewError(CodeFileNotFound, "文件不存在")

	ErrDatabaseError      = NewError(CodeDatabaseError, "数据库错误")
	ErrDatabaseConnection = NewError(CodeDatabaseConnection, "数据库连接错误")
	ErrDatabaseTimeout    = NewError(CodeDatabaseTimeout, "数据库超时")

	ErrCacheError      = NewError(CodeCacheError, "缓存错误")
	ErrCacheConnection = NewError(CodeCacheConnection, "缓存连接错误")
	ErrCacheTimeout    = NewError(CodeCacheTimeout, "缓存超时")

	ErrThirdPartyError       = NewError(CodeThirdPartyError, "第三方服务错误")
	ErrThirdPartyTimeout     = NewError(CodeThirdPartyTimeout, "第三方服务超时")
	ErrThirdPartyUnavailable = NewError(CodeThirdPartyUnavailable, "第三方服务不可用")
)

// AppError 应用错误结构
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error 实现error接口
func (e *AppError) Error() string {
	return e.Message
}

// NewError 创建新的应用错误
func NewError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// NewErrorWithCode 根据错误码创建错误
func NewErrorWithCode(code int) *AppError {
	message, exists := errorMessages[code]
	if !exists {
		message = "未知错误"
	}
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// GetErrorMessage 获取错误信息
func GetErrorMessage(code int) string {
	message, exists := errorMessages[code]
	if !exists {
		return "未知错误"
	}
	return message
}

// IsAppError 判断是否为应用错误
func IsAppError(err error) (*AppError, bool) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}

// GetErrorCode 获取错误码
func GetErrorCode(err error) int {
	if appErr, ok := IsAppError(err); ok {
		return appErr.Code
	}
	return CodeInternalServerError
}
