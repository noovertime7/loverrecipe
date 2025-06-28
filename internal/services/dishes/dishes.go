package dishes

import (
	"context"
	"loverrecipe/internal/domain"
	"loverrecipe/internal/repository"
)

type Service interface {
	CreateDishes(ctx context.Context, req domain.CreateDishesRequest) (*domain.Dishes, error)
	GetDishesByID(ctx context.Context, id int64) (*domain.Dishes, error)
	GetDishesByUserID(ctx context.Context, userID int64) ([]domain.Dishes, error)
	GetDishesByType(ctx context.Context, typeID int64) ([]domain.Dishes, error)
	GetDishesByUserIDAndType(ctx context.Context, userID int64, typeID int64) ([]domain.Dishes, error)
	GetDishesWithTypeInfo(ctx context.Context, userID int64) ([]domain.DishesWithType, error)
	UpdateDishes(ctx context.Context, req domain.UpdateDishesRequest) (*domain.Dishes, error)
	DeleteDishes(ctx context.Context, id int64, userID int64) error
	ListDishes(ctx context.Context, query domain.DishesQuery) (*domain.DishesListResponse, error)
	GetDishesCount(ctx context.Context) (int64, error)
	SearchDishes(ctx context.Context, userID int64, keyword string, offset int, limit int) (*domain.DishesListResponse, error)
	GetDishesStatistics(ctx context.Context, userID int64) (*DishesStatistics, error)
}

type service struct {
	repo repository.DishesRepository
}

// NewService 创建菜品服务实例
func NewService(repo repository.DishesRepository) Service {
	return &service{
		repo: repo,
	}
}

// CreateDishes 创建菜品
func (s *service) CreateDishes(ctx context.Context, req domain.CreateDishesRequest) (*domain.Dishes, error) {
	// 业务逻辑验证
	if err := s.validateCreateRequest(req); err != nil {
		return nil, err
	}

	// 调用仓储层创建菜品
	dishes, err := s.repo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

// GetDishesByID 根据ID获取菜品
func (s *service) GetDishesByID(ctx context.Context, id int64) (*domain.Dishes, error) {
	if id <= 0 {
		return nil, domain.ErrDishesNotFound
	}

	dishes, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

// GetDishesByUserID 根据用户ID获取菜品列表
func (s *service) GetDishesByUserID(ctx context.Context, userID int64) ([]domain.Dishes, error) {
	if userID <= 0 {
		return nil, domain.ErrDishesUserMismatch
	}

	dishes, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

// GetDishesByType 根据菜品种类获取菜品列表
func (s *service) GetDishesByType(ctx context.Context, typeID int64) ([]domain.Dishes, error) {
	if typeID <= 0 {
		return nil, domain.ErrDishesTypeInvalid
	}

	dishes, err := s.repo.GetByType(ctx, typeID)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

// GetDishesByUserIDAndType 根据用户ID和菜品种类获取菜品列表
func (s *service) GetDishesByUserIDAndType(ctx context.Context, userID int64, typeID int64) ([]domain.Dishes, error) {
	if userID <= 0 {
		return nil, domain.ErrDishesUserMismatch
	}
	if typeID <= 0 {
		return nil, domain.ErrDishesTypeInvalid
	}

	dishes, err := s.repo.GetByUserIDAndType(ctx, userID, typeID)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

// GetDishesWithTypeInfo 获取菜品及其种类信息
func (s *service) GetDishesWithTypeInfo(ctx context.Context, userID int64) ([]domain.DishesWithType, error) {
	if userID <= 0 {
		return nil, domain.ErrDishesUserMismatch
	}

	dishesWithType, err := s.repo.GetDishesWithTypeInfo(ctx, userID)
	if err != nil {
		return nil, err
	}

	return dishesWithType, nil
}

// UpdateDishes 更新菜品
func (s *service) UpdateDishes(ctx context.Context, req domain.UpdateDishesRequest) (*domain.Dishes, error) {
	// 业务逻辑验证
	if err := s.validateUpdateRequest(req); err != nil {
		return nil, err
	}

	// 调用仓储层更新菜品
	dishes, err := s.repo.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

// DeleteDishes 删除菜品
func (s *service) DeleteDishes(ctx context.Context, id int64, userID int64) error {
	if id <= 0 {
		return domain.ErrDishesNotFound
	}
	if userID <= 0 {
		return domain.ErrDishesUserMismatch
	}

	err := s.repo.Delete(ctx, id, userID)
	if err != nil {
		return err
	}

	return nil
}

// ListDishes 分页查询菜品列表
func (s *service) ListDishes(ctx context.Context, query domain.DishesQuery) (*domain.DishesListResponse, error) {
	// 验证查询参数
	if err := s.validateQuery(query); err != nil {
		return nil, err
	}

	// 设置默认分页参数
	if query.Limit <= 0 {
		query.Limit = 10
	}
	if query.Limit > 100 {
		query.Limit = 100
	}
	if query.Offset < 0 {
		query.Offset = 0
	}

	result, err := s.repo.List(ctx, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDishesCount 获取菜品总数
func (s *service) GetDishesCount(ctx context.Context) (int64, error) {
	count, err := s.repo.Count(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// SearchDishes 搜索菜品（简单实现，实际项目中可能需要全文搜索）
func (s *service) SearchDishes(ctx context.Context, userID int64, keyword string, offset int, limit int) (*domain.DishesListResponse, error) {
	if userID <= 0 {
		return nil, domain.ErrDishesUserMismatch
	}
	if keyword == "" {
		// 如果关键词为空，返回所有菜品
		query := domain.DishesQuery{
			UserID: userID,
			Offset: offset,
			Limit:  limit,
		}
		return s.ListDishes(ctx, query)
	}

	// 获取用户所有菜品
	allDishes, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 简单关键词匹配
	var matchedDishes []domain.Dishes
	for _, dish := range allDishes {
		if contains(dish.Name, keyword) || contains(dish.Desc, keyword) {
			matchedDishes = append(matchedDishes, dish)
		}
	}

	// 手动分页
	total := int64(len(matchedDishes))
	start := offset
	end := start + limit
	if end > len(matchedDishes) {
		end = len(matchedDishes)
	}
	if start > len(matchedDishes) {
		start = len(matchedDishes)
	}

	pagedDishes := matchedDishes[start:end]

	// 转换为带种类信息的格式
	var result []domain.DishesWithType
	dishesWithType, err := s.repo.GetDishesWithTypeInfo(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, dish := range pagedDishes {
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
		Page:  offset/limit + 1,
		Size:  limit,
	}, nil
}

// GetDishesStatistics 获取菜品统计信息
func (s *service) GetDishesStatistics(ctx context.Context, userID int64) (*DishesStatistics, error) {
	if userID <= 0 {
		return nil, domain.ErrDishesUserMismatch
	}

	// 获取用户所有菜品
	allDishes, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 计算统计信息
	stats := &DishesStatistics{
		TotalDishes:  int64(len(allDishes)),
		TotalPrice:   0,
		TotalCalorie: 0,
		AvgPrice:     0,
		AvgCalorie:   0,
	}

	if stats.TotalDishes > 0 {
		for _, dish := range allDishes {
			stats.TotalPrice += dish.Price
			stats.TotalCalorie += dish.Calorie
		}
		stats.AvgPrice = stats.TotalPrice / stats.TotalDishes
		stats.AvgCalorie = stats.TotalCalorie / stats.TotalDishes
	}

	return stats, nil
}

// validateCreateRequest 验证创建请求
func (s *service) validateCreateRequest(req domain.CreateDishesRequest) error {
	// 基础验证已在domain层完成，这里可以添加服务层特有的验证逻辑
	return nil
}

// validateUpdateRequest 验证更新请求
func (s *service) validateUpdateRequest(req domain.UpdateDishesRequest) error {
	// 基础验证已在domain层完成，这里可以添加服务层特有的验证逻辑
	return nil
}

// validateQuery 验证查询参数
func (s *service) validateQuery(query domain.DishesQuery) error {
	if query.UserID <= 0 {
		return domain.ErrDishesUserMismatch
	}
	return nil
}

// contains 简单的字符串包含检查
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		(len(s) > len(substr) && (s[:len(substr)] == substr ||
			s[len(s)-len(substr):] == substr ||
			func() bool {
				for i := 1; i <= len(s)-len(substr); i++ {
					if s[i:i+len(substr)] == substr {
						return true
					}
				}
				return false
			}())))
}

// DishesStatistics 菜品统计信息
type DishesStatistics struct {
	TotalDishes  int64 `json:"total_dishes"`
	TotalPrice   int64 `json:"total_price"`
	TotalCalorie int64 `json:"total_calorie"`
	AvgPrice     int64 `json:"avg_price"`
	AvgCalorie   int64 `json:"avg_calorie"`
}
