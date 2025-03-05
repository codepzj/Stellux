package service

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"server/internal/user/internal/domain"
	"server/internal/user/internal/repo"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *domain.User) (bson.ObjectID, error)
	FindUserIsExist(ctx context.Context, user *domain.User) (*UserDto, bool)
	FindAllUsers(ctx context.Context) ([]*UserDto, error)
}

var _ IUserService = (*UserService)(nil)

type UserService struct {
	repo repo.IUserRepo
}

func NewUserService(repo repo.IUserRepo) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (bson.ObjectID, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) FindUserIsExist(ctx context.Context, user *domain.User) (*UserDto, bool) {
	user, ok := s.repo.FindUserIsExist(ctx, user)
	if !ok {
		return nil, false
	}
	return DoToDTO(user), true
}

func (s *UserService) FindAllUsers(ctx context.Context) ([]*UserDto, error) {
	users, err := s.repo.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return DOsToDTOs(users), nil
}
