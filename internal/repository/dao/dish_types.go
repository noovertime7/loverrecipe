package dao

import (
	"context"

	"github.com/ego-component/egorm"
)

type DishType struct {
	ID          int64  `gorm:"primaryKey;type:BIGINT;comment:'业务标识'"`
	UserID      int64  `gorm:"type:BIGINT;comment:'用户ID'"`
	Ctime       int64  `gorm:"comment:'创建时间'"`
	Utime       int64  `gorm:"comment:'更新时间'"`
	Name        string `gorm:"type:VARCHAR(50);comment:'种类名称'"`
	Description string `gorm:"type:VARCHAR(200);comment:'种类描述'"`
	Icon        string `gorm:"type:VARCHAR(200);comment:'种类图标'"`
	Color       string `gorm:"type:VARCHAR(20);comment:'种类颜色'"`
	Sort        int64  `gorm:"type:BIGINT;default:0;comment:'排序权重'"`
	Status      int64  `gorm:"type:BIGINT;default:1;comment:'状态 1:启用 0:禁用'"`
}

// TableName 重命名表
func (DishType) TableName() string {
	return "dish_types"
}

type DishTypeDao interface {
	GetByIDs(ctx context.Context, ids []int64) (map[int64]DishType, error)
	GetByID(ctx context.Context, id int64) (DishType, error)
	GetByUserID(ctx context.Context, userID int64) ([]DishType, error)
	Delete(ctx context.Context, id int64) error
	Save(ctx context.Context, dishType DishType) (DishType, error)
	Find(ctx context.Context, offset int, limit int) ([]DishType, error)
	FindByStatus(ctx context.Context, status int64, offset int, limit int) ([]DishType, error)
	Count(ctx context.Context) (int64, error)
}

// Implementation of the DishTypeDao interface
type dishTypeDAO struct {
	db *egorm.Component
}

// NewDishTypeDao creates a new instance of DishTypeDao
func NewDishTypeDao(db *egorm.Component) DishTypeDao {
	return &dishTypeDAO{db: db}
}

// GetByIDs 根据ID列表批量获取菜品种类信息
func (d *dishTypeDAO) GetByIDs(ctx context.Context, ids []int64) (map[int64]DishType, error) {
	var dishTypes []DishType
	result := make(map[int64]DishType)

	if len(ids) == 0 {
		return result, nil
	}

	err := d.db.WithContext(ctx).Where("id IN ?", ids).Find(&dishTypes).Error
	if err != nil {
		return nil, err
	}

	for _, dishType := range dishTypes {
		result[dishType.ID] = dishType
	}

	return result, nil
}

// GetByID 根据ID获取单个菜品种类信息
func (d *dishTypeDAO) GetByID(ctx context.Context, id int64) (DishType, error) {
	var dishType DishType
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&dishType).Error
	return dishType, err
}

// GetByUserID 根据用户ID获取菜品种类列表
func (d *dishTypeDAO) GetByUserID(ctx context.Context, userID int64) ([]DishType, error) {
	var dishTypes []DishType
	err := d.db.WithContext(ctx).Where("user_id = ?", userID).Order("sort DESC, id ASC").Find(&dishTypes).Error
	return dishTypes, err
}

// Delete 根据ID删除菜品种类
func (d *dishTypeDAO) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Where("id = ?", id).Delete(&DishType{}).Error
}

// Save 保存或更新菜品种类信息
func (d *dishTypeDAO) Save(ctx context.Context, dishType DishType) (DishType, error) {
	if dishType.ID == 0 {
		// 新增菜品种类
		err := d.db.WithContext(ctx).Create(&dishType).Error
		return dishType, err
	} else {
		// 更新菜品种类
		err := d.db.WithContext(ctx).Save(&dishType).Error
		return dishType, err
	}
}

// Find 分页查询菜品种类列表
func (d *dishTypeDAO) Find(ctx context.Context, offset int, limit int) ([]DishType, error) {
	var dishTypes []DishType
	err := d.db.WithContext(ctx).Order("sort DESC, id ASC").Offset(offset).Limit(limit).Find(&dishTypes).Error
	return dishTypes, err
}

// FindByStatus 根据状态分页查询菜品种类列表
func (d *dishTypeDAO) FindByStatus(ctx context.Context, status int64, offset int, limit int) ([]DishType, error) {
	var dishTypes []DishType
	err := d.db.WithContext(ctx).Where("status = ?", status).Order("sort DESC, id ASC").Offset(offset).Limit(limit).Find(&dishTypes).Error
	return dishTypes, err
}

// Count 统计菜品种类总数
func (d *dishTypeDAO) Count(ctx context.Context) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&DishType{}).Count(&count).Error
	return count, err
}
