package repo

import (
	"context"

	"server/internal/user_detail/internal/domain"
	"server/internal/user_detail/internal/repo/dao"
)

type IUserDetailRepo interface {
	CreateOne(ctx context.Context, userDetail *domain.UserDetail) error
	UpdateOne(ctx context.Context, userDetail *domain.UserDetail) error
}

var _ IUserDetailRepo = (*UserDetailRepo)(nil)

type UserDetailRepo struct {
	userDetailDao dao.IUserDetailDao
}

func NewUserDetailRepo(userDetailDao dao.IUserDetailDao) *UserDetailRepo {
	return &UserDetailRepo{userDetailDao: userDetailDao}
}

func (u *UserDetailRepo) CreateOne(ctx context.Context, userDetail *domain.UserDetail) error {
	return u.userDetailDao.CreateOne(ctx, userDetail)
}

func (u *UserDetailRepo) UpdateOne(ctx context.Context, userDetail *domain.UserDetail) error {
	return u.userDetailDao.UpdateOne(ctx, userDetail)
}
