package user

import (
	"context"
	"github.com/gotomicro/ego/core/elog"
	"github.com/sony/sonyflake"
	"loverrecipe/internal/domain"
	"loverrecipe/internal/repository"
	"loverrecipe/internal/token"
	"loverrecipe/internal/utils"
)

type Service interface {
	Create(ctx context.Context, user *domain.CreateUserInput) (domain.CreateUserOutput, error)
}

type service struct {
	repo repository.UserRepository
	jwt  *token.JwtTokenHandler
	id   *sonyflake.Sonyflake
}

func NewService(repo repository.UserRepository, jwt *token.JwtTokenHandler, id *sonyflake.Sonyflake) Service {
	return &service{
		repo: repo,
		jwt:  jwt,
		id:   id,
	}
}

func (s *service) Create(ctx context.Context, user *domain.CreateUserInput) (domain.CreateUserOutput, error) {
	user.Password = utils.HashPassword(user.Password)

	id, err := s.id.NextID()
	if err != nil {
		elog.Error("生成用户ID失败", elog.FieldErr(err))
		return domain.CreateUserOutput{}, err
	}
	user.ID = id

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		elog.Error("创建用户失败", elog.FieldErr(err))
		return domain.CreateUserOutput{}, err
	}

	tokenStr, err := s.jwt.GenerateRefreshToken(token.BaseClaims{})
	if err != nil {
		elog.Error("生成token失败", elog.FieldErr(err))
		return domain.CreateUserOutput{}, err
	}

	return domain.CreateUserOutput{
		Token: tokenStr,
	}, nil

}
