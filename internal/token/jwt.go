package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JwtTokenHandler 结构体定义
type JwtTokenHandler struct {
	secret string
}

var (
	ExpireTime        = time.Minute * 10
	RefreshExpireTime = time.Hour * 24 * 7
)

// RegisterJwt 注册 JWT secret
func RegisterJwt() *JwtTokenHandler {
	return &JwtTokenHandler{
		secret: "lover_recipe",
	}
}

// BaseClaims 基本声明结构体
type BaseClaims struct {
	UserId   uint
	Username string
}

// CustomClaims 自定义声明结构体
type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

// GenerateToken 生成主 Token
func (j *JwtTokenHandler) GenerateToken(baseClaims BaseClaims) (string, error) {
	return j.generateToken(baseClaims, ExpireTime)
}

// GenerateRefreshToken 生成刷新 Token
func (j *JwtTokenHandler) GenerateRefreshToken(baseClaims BaseClaims) (string, error) {
	return j.generateToken(baseClaims, RefreshExpireTime)
}

// generateToken 生成 Token 的通用方法
func (j *JwtTokenHandler) generateToken(baseClaims BaseClaims, duration time.Duration) (string, error) {
	expireTime := time.Now().Add(duration)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000 * time.Millisecond)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(expireTime),                               // 过期时间
			Issuer:    "lover",                                                      // 签名的发行者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

// ParseToken 解析 Token
func (j *JwtTokenHandler) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		err := parseTokenError(err)
		return nil, err
	}

	// 解析成功
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("解析 token 失败")
}

// parseTokenError 处理 Token 解析错误
func parseTokenError(err error) error {
	if errors.Is(err, jwt.ErrTokenExpired) {
		return errors.New("登录过期，请重新登录")
	}
	return errors.New("token 不可用: " + err.Error())
}
