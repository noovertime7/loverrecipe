package repository

import (
	"context"
	"github.com/pkg/errors"
	"loverrecipe/internal/domain"
	"loverrecipe/internal/repository/dao"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.CreateUserInput) error
}

type userRepository struct {
	dao dao.UserDao
}

func NewUserRepository(dao dao.UserDao) UserRepository {
	return &userRepository{dao: dao}
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.CreateUserInput) error {
	du := dao.User{
		ID:       user.ID,
		Username: user.Name,
		Password: user.Password,
		Avatar:   user.Avatar,
		Status:   user.Status,
	}

	_, err := r.dao.Create(ctx, du)
	if err != nil {
		return errors.Wrap(err, "create user failed")
	}
	return nil
}
