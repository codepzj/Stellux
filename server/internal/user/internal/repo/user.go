package repo

import (
	"context"

	"github.com/codepzj/Stellux/server/internal/pkg/utils"
	"github.com/codepzj/Stellux/server/internal/user/internal/domain"
	"github.com/codepzj/Stellux/server/internal/user/internal/repo/dao"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user *domain.User) (bson.ObjectID, error)
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

func (u *UserRepo) CreateUser(ctx context.Context, user *domain.User) (bson.ObjectID, error) {
	hashPassword, err := utils.GenerateHashPassword(user.Password)
	// bcrypt加密密码
	if err != nil {
		return bson.ObjectID{}, errors.New("系统错误，密码加密失败")
	}
	user.Password = hashPassword
	return u.dao.CreateOne(ctx, user)
}

func (u *UserRepo) FindUserIsExist(ctx context.Context, user *domain.User) (*domain.User, bool) {
	return u.dao.FindUserIsExist(ctx, user)

}

func (u *UserRepo) FindAllUsers(ctx context.Context) ([]*domain.User, error) {
	return u.dao.FindAll(ctx)
}
