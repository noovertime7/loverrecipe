package dao

import (
	"context"

	"github.com/ego-component/egorm"
)

type Dishes struct {
	ID      int64 `gorm:"primaryKey;type:BIGINT;comment:'业务标识'"`
	UserID  int64 `gorm:"type:BIGINT;comment:'用户ID'"`
	Ctime   int64
	Utime   int64
	Name    string `gorm:"type:VARCHAR(100);comment:'菜名'"`
	Desc    string `gorm:"type:VARCHAR(200);comment:'菜描述'"`
	Price   int64  `gorm:"type:BIGINT;comment:'价格'"`
	Img     string `gorm:"type:VARCHAR(200);comment:'菜图片'"`
	Type    int64  `gorm:"type:BIGINT;comment:'菜类别';index:idx_type"`
	Calorie int64  `gorm:"type:BIGINT;comment:'卡路里'"`
}

// TableName 重命名表
func (Dishes) TableName() string {
	return "dishes"
}

// DishesWithType 包含菜品和种类信息的结构体
type DishesWithType struct {
	Dishes
	TypeName        string `json:"type_name"`
	TypeDescription string `json:"type_description"`
	TypeIcon        string `json:"type_icon"`
	TypeColor       string `json:"type_color"`
}

type DishesDao interface {
	GetByIDs(ctx context.Context, id []int64) (map[int64]Dishes, error)
	GetByID(ctx context.Context, id int64) (Dishes, error)
	GetByUserID(ctx context.Context, userID int64) ([]Dishes, error)
	GetByType(ctx context.Context, typeID int64) ([]Dishes, error)
	GetByUserIDAndType(ctx context.Context, userID int64, typeID int64) ([]Dishes, error)
	GetDishesWithTypeInfo(ctx context.Context, userID int64) ([]DishesWithType, error)
	Delete(ctx context.Context, id int64) error
	Save(ctx context.Context, config Dishes) (Dishes, error)
	Find(ctx context.Context, offset int, limit int) ([]Dishes, error)
	Count(ctx context.Context) (int64, error)
}

// Implementation of the DishesDao interface
type dishesDAO struct {
	db *egorm.Component
}

// NewDishesDao creates a new instance of DishesDao
func NewDishesDao(db *egorm.Component) DishesDao {
	return &dishesDAO{db: db}
}

// GetByIDs 根据ID列表批量获取菜品信息
func (d *dishesDAO) GetByIDs(ctx context.Context, ids []int64) (map[int64]Dishes, error) {
	var dishes []Dishes
	result := make(map[int64]Dishes)

	if len(ids) == 0 {
		return result, nil
	}

	err := d.db.WithContext(ctx).Where("id IN ?", ids).Find(&dishes).Error
	if err != nil {
		return nil, err
	}

	for _, dish := range dishes {
		result[dish.ID] = dish
	}

	return result, nil
}

// GetByID 根据ID获取单个菜品信息
func (d *dishesDAO) GetByID(ctx context.Context, id int64) (Dishes, error) {
	var dish Dishes
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&dish).Error
	return dish, err
}

// Delete 根据ID删除菜品
func (d *dishesDAO) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Where("id = ?", id).Delete(&Dishes{}).Error
}

// Save 保存或更新菜品信息
func (d *dishesDAO) Save(ctx context.Context, dish Dishes) (Dishes, error) {
	if dish.ID == 0 {
		// 新增菜品
		err := d.db.WithContext(ctx).Create(&dish).Error
		return dish, err
	} else {
		// 更新菜品
		err := d.db.WithContext(ctx).Save(&dish).Error
		return dish, err
	}
}

// Find 分页查询菜品列表
func (d *dishesDAO) Find(ctx context.Context, offset int, limit int) ([]Dishes, error) {
	var dishes []Dishes
	err := d.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&dishes).Error
	return dishes, err
}

// GetByUserID 根据用户ID获取菜品列表
func (d *dishesDAO) GetByUserID(ctx context.Context, userID int64) ([]Dishes, error) {
	var dishes []Dishes
	err := d.db.WithContext(ctx).Where("user_id = ?", userID).Find(&dishes).Error
	return dishes, err
}

// GetByType 根据菜品种类获取菜品列表
func (d *dishesDAO) GetByType(ctx context.Context, typeID int64) ([]Dishes, error) {
	var dishes []Dishes
	err := d.db.WithContext(ctx).Where("type = ?", typeID).Find(&dishes).Error
	return dishes, err
}

// GetByUserIDAndType 根据用户ID和菜品种类获取菜品列表
func (d *dishesDAO) GetByUserIDAndType(ctx context.Context, userID int64, typeID int64) ([]Dishes, error) {
	var dishes []Dishes
	err := d.db.WithContext(ctx).Where("user_id = ? AND type = ?", userID, typeID).Find(&dishes).Error
	return dishes, err
}

// Count 统计菜品总数
func (d *dishesDAO) Count(ctx context.Context) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&Dishes{}).Count(&count).Error
	return count, err
}

// GetDishesWithTypeInfo 获取菜品及其种类信息
func (d *dishesDAO) GetDishesWithTypeInfo(ctx context.Context, userID int64) ([]DishesWithType, error) {
	var result []DishesWithType

	// 使用JOIN查询获取菜品和种类信息
	err := d.db.WithContext(ctx).
		Table("dishes").
		Select("dishes.*, dish_types.name as type_name, dish_types.description as type_description, dish_types.icon as type_icon, dish_types.color as type_color").
		Joins("LEFT JOIN dish_types ON dishes.type = dish_types.id").
		Where("dishes.user_id = ?", userID).
		Find(&result).Error

	return result, err
}
