package domain

import (
	"errors"
	"time"
)

// Dishes 菜品领域模型
type Dishes struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Price   int64  `json:"price"`
	Img     string `json:"img"`
	Type    int64  `json:"type"`
	Calorie int64  `json:"calorie"`
	Ctime   int64  `json:"ctime"`
	Utime   int64  `json:"utime"`
}

// DishesWithType 包含种类信息的菜品
type DishesWithType struct {
	Dishes
	TypeName        string `json:"type_name"`
	TypeDescription string `json:"type_description"`
	TypeIcon        string `json:"type_icon"`
	TypeColor       string `json:"type_color"`
}

// CreateDishesRequest 创建菜品请求
type CreateDishesRequest struct {
	UserID  int64  `json:"user_id" validate:"required"`
	Name    string `json:"name" validate:"required,min=1,max=100"`
	Desc    string `json:"desc" validate:"max=200"`
	Price   int64  `json:"price" validate:"min=0"`
	Img     string `json:"img" validate:"max=200"`
	Type    int64  `json:"type" validate:"required"`
	Calorie int64  `json:"calorie" validate:"min=0"`
}

// UpdateDishesRequest 更新菜品请求
type UpdateDishesRequest struct {
	ID      int64  `json:"id" validate:"required"`
	UserID  int64  `json:"user_id" validate:"required"`
	Name    string `json:"name" validate:"required,min=1,max=100"`
	Desc    string `json:"desc" validate:"max=200"`
	Price   int64  `json:"price" validate:"min=0"`
	Img     string `json:"img" validate:"max=200"`
	Type    int64  `json:"type" validate:"required"`
	Calorie int64  `json:"calorie" validate:"min=0"`
}

// DishesQuery 菜品查询条件
type DishesQuery struct {
	UserID int64 `json:"user_id"`
	Type   int64 `json:"type"`
	Offset int   `json:"offset"`
	Limit  int   `json:"limit"`
}

// DishesListResponse 菜品列表响应
type DishesListResponse struct {
	List  []DishesWithType `json:"list"`
	Total int64            `json:"total"`
	Page  int              `json:"page"`
	Size  int              `json:"size"`
}

// 错误定义
var (
	ErrDishesNotFound     = errors.New("菜品不存在")
	ErrDishesNameEmpty    = errors.New("菜品名称不能为空")
	ErrDishesPriceInvalid = errors.New("菜品价格无效")
	ErrDishesTypeInvalid  = errors.New("菜品类型无效")
	ErrDishesUserMismatch = errors.New("菜品不属于该用户")
)

// NewDishes 创建新的菜品实例
func NewDishes(req CreateDishesRequest) (*Dishes, error) {
	// 验证请求参数
	if err := req.Validate(); err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	dishes := &Dishes{
		UserID:  req.UserID,
		Name:    req.Name,
		Desc:    req.Desc,
		Price:   req.Price,
		Img:     req.Img,
		Type:    req.Type,
		Calorie: req.Calorie,
		Ctime:   now,
		Utime:   now,
	}

	return dishes, nil
}

// Validate 验证创建菜品请求
func (req CreateDishesRequest) Validate() error {
	if req.UserID <= 0 {
		return errors.New("用户ID无效")
	}
	if req.Name == "" {
		return ErrDishesNameEmpty
	}
	if len(req.Name) > 100 {
		return errors.New("菜品名称过长")
	}
	if len(req.Desc) > 200 {
		return errors.New("菜品描述过长")
	}
	if req.Price < 0 {
		return ErrDishesPriceInvalid
	}
	if len(req.Img) > 200 {
		return errors.New("图片URL过长")
	}
	if req.Type <= 0 {
		return ErrDishesTypeInvalid
	}
	if req.Calorie < 0 {
		return errors.New("卡路里不能为负数")
	}
	return nil
}

// Validate 验证更新菜品请求
func (req UpdateDishesRequest) Validate() error {
	if req.ID <= 0 {
		return errors.New("菜品ID无效")
	}
	if req.UserID <= 0 {
		return errors.New("用户ID无效")
	}
	if req.Name == "" {
		return ErrDishesNameEmpty
	}
	if len(req.Name) > 100 {
		return errors.New("菜品名称过长")
	}
	if len(req.Desc) > 200 {
		return errors.New("菜品描述过长")
	}
	if req.Price < 0 {
		return ErrDishesPriceInvalid
	}
	if len(req.Img) > 200 {
		return errors.New("图片URL过长")
	}
	if req.Type <= 0 {
		return ErrDishesTypeInvalid
	}
	if req.Calorie < 0 {
		return errors.New("卡路里不能为负数")
	}
	return nil
}

// Update 更新菜品信息
func (d *Dishes) Update(req UpdateDishesRequest) error {
	// 验证用户权限
	if d.UserID != req.UserID {
		return ErrDishesUserMismatch
	}

	// 验证请求参数
	if err := req.Validate(); err != nil {
		return err
	}

	// 更新字段
	d.Name = req.Name
	d.Desc = req.Desc
	d.Price = req.Price
	d.Img = req.Img
	d.Type = req.Type
	d.Calorie = req.Calorie
	d.Utime = time.Now().Unix()

	return nil
}

// CanDelete 检查是否可以删除
func (d *Dishes) CanDelete(userID int64) error {
	if d.UserID != userID {
		return ErrDishesUserMismatch
	}
	return nil
}

// ToResponse 转换为响应格式
func (d *Dishes) ToResponse() map[string]interface{} {
	return map[string]interface{}{
		"id":      d.ID,
		"user_id": d.UserID,
		"name":    d.Name,
		"desc":    d.Desc,
		"price":   d.Price,
		"img":     d.Img,
		"type":    d.Type,
		"calorie": d.Calorie,
		"ctime":   d.Ctime,
		"utime":   d.Utime,
	}
}
