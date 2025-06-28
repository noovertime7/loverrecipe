package repository

import (
	"context"
	"loverrecipe/internal/domain"
	"loverrecipe/internal/repository/dao"

	"github.com/ego-component/egorm"
)

type DishesRepository interface {
	Create(ctx context.Context, req domain.CreateDishesRequest) (*domain.Dishes, error)
	GetByID(ctx context.Context, id int64) (*domain.Dishes, error)
	GetByUserID(ctx context.Context, userID int64) ([]domain.Dishes, error)
	GetByType(ctx context.Context, typeID int64) ([]domain.Dishes, error)
	GetByUserIDAndType(ctx context.Context, userID int64, typeID int64) ([]domain.Dishes, error)
	GetDishesWithTypeInfo(ctx context.Context, userID int64) ([]domain.DishesWithType, error)
	Update(ctx context.Context, req domain.UpdateDishesRequest) (*domain.Dishes, error)
	Delete(ctx context.Context, id int64, userID int64) error
	List(ctx context.Context, query domain.DishesQuery) (*domain.DishesListResponse, error)
	Count(ctx context.Context) (int64, error)
}

type dishesRepository struct {
	dishesDao   dao.DishesDao
	dishTypeDao dao.DishTypeDao
}

func NewDishesRepository(db *egorm.Component) DishesRepository {
	return &dishesRepository{
		dishesDao:   dao.NewDishesDao(db),
		dishTypeDao: dao.NewDishTypeDao(db),
	}
}

// Create 创建菜品
func (r *dishesRepository) Create(ctx context.Context, req domain.CreateDishesRequest) (*domain.Dishes, error) {
	// 创建领域对象
	dishes, err := domain.NewDishes(req)
	if err != nil {
		return nil, err
	}

	// 转换为DAO对象
	daoDishes := dao.Dishes{
		UserID:  dishes.UserID,
		Name:    dishes.Name,
		Desc:    dishes.Desc,
		Price:   dishes.Price,
		Img:     dishes.Img,
		Type:    dishes.Type,
		Calorie: dishes.Calorie,
		Ctime:   dishes.Ctime,
		Utime:   dishes.Utime,
	}

	// 保存到数据库
	savedDishes, err := r.dishesDao.Save(ctx, daoDishes)
	if err != nil {
		return nil, err
	}

	// 转换回领域对象
	dishes.ID = savedDishes.ID
	return dishes, nil
}

// GetByID 根据ID获取菜品
func (r *dishesRepository) GetByID(ctx context.Context, id int64) (*domain.Dishes, error) {
	daoDishes, err := r.dishesDao.GetByID(ctx, id)
	if err != nil {
		return nil, domain.ErrDishesNotFound
	}

	return r.daoToDomain(daoDishes), nil
}

// GetByUserID 根据用户ID获取菜品列表
func (r *dishesRepository) GetByUserID(ctx context.Context, userID int64) ([]domain.Dishes, error) {
	daoDishes, err := r.dishesDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return r.daoListToDomainList(daoDishes), nil
}

// GetByType 根据菜品种类获取菜品列表
func (r *dishesRepository) GetByType(ctx context.Context, typeID int64) ([]domain.Dishes, error) {
	daoDishes, err := r.dishesDao.GetByType(ctx, typeID)
	if err != nil {
		return nil, err
	}

	return r.daoListToDomainList(daoDishes), nil
}

// GetByUserIDAndType 根据用户ID和菜品种类获取菜品列表
func (r *dishesRepository) GetByUserIDAndType(ctx context.Context, userID int64, typeID int64) ([]domain.Dishes, error) {
	daoDishes, err := r.dishesDao.GetByUserIDAndType(ctx, userID, typeID)
	if err != nil {
		return nil, err
	}

	return r.daoListToDomainList(daoDishes), nil
}

// GetDishesWithTypeInfo 获取菜品及其种类信息
func (r *dishesRepository) GetDishesWithTypeInfo(ctx context.Context, userID int64) ([]domain.DishesWithType, error) {
	daoDishesWithType, err := r.dishesDao.GetDishesWithTypeInfo(ctx, userID)
	if err != nil {
		return nil, err
	}

	var result []domain.DishesWithType
	for _, dt := range daoDishesWithType {
		dishesWithType := domain.DishesWithType{
			Dishes: domain.Dishes{
				ID:      dt.ID,
				UserID:  dt.UserID,
				Name:    dt.Name,
				Desc:    dt.Desc,
				Price:   dt.Price,
				Img:     dt.Img,
				Type:    dt.Type,
				Calorie: dt.Calorie,
				Ctime:   dt.Ctime,
				Utime:   dt.Utime,
			},
			TypeName:        dt.TypeName,
			TypeDescription: dt.TypeDescription,
			TypeIcon:        dt.TypeIcon,
			TypeColor:       dt.TypeColor,
		}
		result = append(result, dishesWithType)
	}

	return result, nil
}

// Update 更新菜品
func (r *dishesRepository) Update(ctx context.Context, req domain.UpdateDishesRequest) (*domain.Dishes, error) {
	// 先获取现有菜品
	existingDishes, err := r.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// 更新菜品信息
	if err := existingDishes.Update(req); err != nil {
		return nil, err
	}

	// 转换为DAO对象
	daoDishes := dao.Dishes{
		ID:      existingDishes.ID,
		UserID:  existingDishes.UserID,
		Name:    existingDishes.Name,
		Desc:    existingDishes.Desc,
		Price:   existingDishes.Price,
		Img:     existingDishes.Img,
		Type:    existingDishes.Type,
		Calorie: existingDishes.Calorie,
		Ctime:   existingDishes.Ctime,
		Utime:   existingDishes.Utime,
	}

	// 保存到数据库
	_, err = r.dishesDao.Save(ctx, daoDishes)
	if err != nil {
		return nil, err
	}

	return existingDishes, nil
}

// Delete 删除菜品
func (r *dishesRepository) Delete(ctx context.Context, id int64, userID int64) error {
	// 先获取现有菜品
	existingDishes, err := r.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// 检查删除权限
	if err := existingDishes.CanDelete(userID); err != nil {
		return err
	}

	// 执行删除
	return r.dishesDao.Delete(ctx, id)
}

// List 分页查询菜品列表
func (r *dishesRepository) List(ctx context.Context, query domain.DishesQuery) (*domain.DishesListResponse, error) {
	var daoDishes []dao.Dishes
	var err error

	// 根据查询条件获取数据
	if query.Type > 0 {
		daoDishes, err = r.dishesDao.GetByUserIDAndType(ctx, query.UserID, query.Type)
	} else {
		daoDishes, err = r.dishesDao.GetByUserID(ctx, query.UserID)
	}
	if err != nil {
		return nil, err
	}

	// 转换为领域对象
	dishesList := r.daoListToDomainList(daoDishes)

	// 手动分页
	total := int64(len(dishesList))
	start := query.Offset
	end := start + query.Limit
	if end > len(dishesList) {
		end = len(dishesList)
	}
	if start > len(dishesList) {
		start = len(dishesList)
	}

	pagedList := dishesList[start:end]

	// 获取带种类信息的菜品列表
	dishesWithType, err := r.GetDishesWithTypeInfo(ctx, query.UserID)
	if err != nil {
		return nil, err
	}

	// 过滤分页后的菜品
	var result []domain.DishesWithType
	for _, dish := range pagedList {
		for _, dt := range dishesWithType {
			if dt.ID == dish.ID {
				result = append(result, dt)
				break
			}
		}
	}

	return &domain.DishesListResponse{
		List:  result,
		Total: total,
		Page:  query.Offset/query.Limit + 1,
		Size:  query.Limit,
	}, nil
}

// Count 统计菜品总数
func (r *dishesRepository) Count(ctx context.Context) (int64, error) {
	return r.dishesDao.Count(ctx)
}

// daoToDomain 将DAO对象转换为领域对象
func (r *dishesRepository) daoToDomain(daoDishes dao.Dishes) *domain.Dishes {
	return &domain.Dishes{
		ID:      daoDishes.ID,
		UserID:  daoDishes.UserID,
		Name:    daoDishes.Name,
		Desc:    daoDishes.Desc,
		Price:   daoDishes.Price,
		Img:     daoDishes.Img,
		Type:    daoDishes.Type,
		Calorie: daoDishes.Calorie,
		Ctime:   daoDishes.Ctime,
		Utime:   daoDishes.Utime,
	}
}

// daoListToDomainList 将DAO对象列表转换为领域对象列表
func (r *dishesRepository) daoListToDomainList(daoDishes []dao.Dishes) []domain.Dishes {
	var result []domain.Dishes
	for _, d := range daoDishes {
		result = append(result, *r.daoToDomain(d))
	}
	return result
}
