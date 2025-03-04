package repo

import (
	"context"
	"server/internal/user/internal/domain"
	"server/internal/user/internal/repo/dao"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user *domain.User) error
	FindUserIsExist(ctx context.Context, user *domain.User) (*domain.User, bool)
	FindAllUsers(ctx context.Context) ([]*domain.User, error)
}

var _ IUserRepo = (*UserRepo)(nil)

type UserRepo struct {
	dao dao.IUserDao
}

func NewUserRepo(dao dao.IUserDao) *UserRepo {
	return &UserRepo{dao}
}

func (u *UserRepo) CreateUser(ctx context.Context, user *domain.User) error {
	return u.dao.CreateOne(ctx, user)
}

func (u *UserRepo) FindUserIsExist(ctx context.Context, user *domain.User) (*domain.User, bool) {
	return u.dao.FindUserIsExist(ctx, user)

}

func (u *UserRepo) FindAllUsers(ctx context.Context) ([]*domain.User, error) {
	return u.dao.FindAll(ctx)
}
