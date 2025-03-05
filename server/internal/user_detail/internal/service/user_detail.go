package service

import (
	"context"

	"server/internal/user_detail/internal/domain"
	"server/internal/user_detail/internal/repo"
)

type IUserDetailService interface {
	CreateUserDetail(ctx context.Context, userDetail *domain.UserDetail) error
	UpdateUserDetail(ctx context.Context, userDetail *domain.UserDetail) error
}

type UserDetailService struct {
	userDetailRepo repo.IUserDetailRepo
}

func NewUserDetailService(userDetailRepo repo.IUserDetailRepo) *UserDetailService {
	return &UserDetailService{userDetailRepo: userDetailRepo}
}

func (u *UserDetailService) CreateUserDetail(ctx context.Context, userDetail *domain.UserDetail) error {
	return u.userDetailRepo.CreateUserDetail(ctx, userDetail)
}

func (u *UserDetailService) UpdateUserDetail(ctx context.Context, userDetail *domain.UserDetail) error {
	return u.userDetailRepo.UpdateUserDetail(ctx, userDetail)
}
