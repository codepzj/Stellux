package repo

import (
	"context"

	"server/internal/user_detail/internal/domain"
	"server/internal/user_detail/internal/repo/dao"
)

type IUserDetailRepo interface {
	CreateUserDetail(ctx context.Context, userDetail *domain.UserDetail) error
	UpdateUserDetail(ctx context.Context, userDetail *domain.UserDetail) error
}

var _ IUserDetailRepo = (*UserDetailRepo)(nil)

type UserDetailRepo struct {
	userDetailDao dao.IUserDetailDao
}

func NewUserDetailRepo(userDetailDao dao.IUserDetailDao) *UserDetailRepo {
	return &UserDetailRepo{userDetailDao: userDetailDao}
}

func (u *UserDetailRepo) CreateUserDetail(ctx context.Context, userDetail *domain.UserDetail) error {
	return u.userDetailDao.CreateUserDetail(ctx, userDetail)
}

func (u *UserDetailRepo) UpdateUserDetail(ctx context.Context, userDetail *domain.UserDetail) error {
	return u.userDetailDao.UpdateUserDetail(ctx, userDetail)
}
