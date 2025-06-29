package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

const salt = "lover"

// 通过MD5加盐的方式加密用户密码

// HashPassword 对密码进行加盐哈希
func HashPassword(password string) string {
	hash := md5.New()
	io.WriteString(hash, password+salt)
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

func ValidatePassword(inPass, md5Pass string) bool {
	return md5Pass == HashPassword(inPass)
}
