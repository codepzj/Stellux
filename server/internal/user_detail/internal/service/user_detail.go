package service

import (
	"context"

	"server/internal/user_detail/internal/domain"
	"server/internal/user_detail/internal/repo"
)

type IUserDetailService interface {
	CreateOne(ctx context.Context, userDetail *domain.UserDetail) error
	UpdateOne(ctx context.Context, userDetail *domain.UserDetail) error
}

type UserDetailService struct {
	userDetailRepo repo.IUserDetailRepo
}

func NewUserDetailService(userDetailRepo repo.IUserDetailRepo) *UserDetailService {
	return &UserDetailService{userDetailRepo: userDetailRepo}
}

func (u *UserDetailService) CreateOne(ctx context.Context, userDetail *domain.UserDetail) error {
	return u.userDetailRepo.CreateOne(ctx, userDetail)
}

func (u *UserDetailService) UpdateOne(ctx context.Context, userDetail *domain.UserDetail) error {
	return u.userDetailRepo.UpdateOne(ctx, userDetail)
}
