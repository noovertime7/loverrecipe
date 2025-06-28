package dao

import (
	"context"
	"time"

	"github.com/ego-component/egorm"
)

type User struct {
	ID        int64  `gorm:"primaryKey;type:BIGINT;comment:'用户ID'"`
	Username  string `gorm:"type:VARCHAR(50);uniqueIndex:idx_username;comment:'用户名'"`
	Password  string `gorm:"type:VARCHAR(255);comment:'密码(加密后)'"`
	Avatar    string `gorm:"type:VARCHAR(200);comment:'头像URL'"`
	Status    int64  `gorm:"type:BIGINT;default:1;comment:'状态 1:正常 0:禁用'"`
	LastLogin int64  `gorm:"type:BIGINT;comment:'最后登录时间'"`
	Ctime     int64  `gorm:"comment:'创建时间'"`
	Utime     int64  `gorm:"comment:'更新时间'"`
}

// TableName 重命名表
func (User) TableName() string {
	return "users"
}

type UserDao interface {
	Create(ctx context.Context, user User) (User, error)
	GetByID(ctx context.Context, id int64) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	Update(ctx context.Context, user User) (User, error)
	UpdatePassword(ctx context.Context, id int64, newPassword string) error
	UpdateLastLogin(ctx context.Context, id int64) error
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, offset int, limit int) ([]User, error)
	Count(ctx context.Context) (int64, error)
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
}

// Implementation of the UserDao interface
type userDAO struct {
	db *egorm.Component
}

// NewUserDao creates a new instance of UserDao
func NewUserDao(db *egorm.Component) UserDao {
	return &userDAO{db: db}
}

// Create 创建新用户
func (u *userDAO) Create(ctx context.Context, user User) (User, error) {
	// 设置创建和更新时间
	now := time.Now().Unix()
	user.Ctime = now
	user.Utime = now

	err := u.db.WithContext(ctx).Create(&user).Error
	return user, err
}

// GetByID 根据ID获取用户信息
func (u *userDAO) GetByID(ctx context.Context, id int64) (User, error) {
	var user User
	err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return user, err
}

// GetByUsername 根据用户名获取用户信息
func (u *userDAO) GetByUsername(ctx context.Context, username string) (User, error) {
	var user User
	err := u.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return user, err
}

// Update 更新用户信息
func (u *userDAO) Update(ctx context.Context, user User) (User, error) {
	user.Utime = time.Now().Unix()
	err := u.db.WithContext(ctx).Save(&user).Error
	return user, err
}

// UpdatePassword 更新用户密码
func (u *userDAO) UpdatePassword(ctx context.Context, id int64, newPassword string) error {
	return u.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"password": newPassword,
		"utime":    time.Now().Unix(),
	}).Error
}

// UpdateLastLogin 更新最后登录时间
func (u *userDAO) UpdateLastLogin(ctx context.Context, id int64) error {
	return u.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"last_login": time.Now().Unix(),
		"utime":      time.Now().Unix(),
	}).Error
}

// Delete 删除用户
func (u *userDAO) Delete(ctx context.Context, id int64) error {
	return u.db.WithContext(ctx).Where("id = ?", id).Delete(&User{}).Error
}

// Find 分页查询用户列表
func (u *userDAO) Find(ctx context.Context, offset int, limit int) ([]User, error) {
	var users []User
	err := u.db.WithContext(ctx).Offset(offset).Limit(limit).Order("ctime DESC").Find(&users).Error
	return users, err
}

// Count 统计用户总数
func (u *userDAO) Count(ctx context.Context) (int64, error) {
	var count int64
	err := u.db.WithContext(ctx).Model(&User{}).Count(&count).Error
	return count, err
}

// CheckUsernameExists 检查用户名是否已存在
func (u *userDAO) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	var count int64
	err := u.db.WithContext(ctx).Model(&User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}
