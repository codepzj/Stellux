package service

import (
	"context"
	"fmt"

	"github.com/codepzj/Stellux/server/global"
	"github.com/codepzj/Stellux/server/internal/pkg/utils"
	"github.com/codepzj/Stellux/server/internal/user/internal/domain"
	"github.com/codepzj/Stellux/server/internal/user/internal/repo"
	"github.com/codepzj/Stellux/server/internal/user_detail"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/writeconcern"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *domain.User) (bson.ObjectID, error)
	FindUserIsExist(ctx context.Context, user *domain.User) (*UserDto, bool)
	FindAllUsers(ctx context.Context) ([]*UserDto, error)
}

var _ IUserService = (*UserService)(nil)

type UserService struct {
	repo           repo.IUserRepo
	userDetailRepo user_detail.Repository
}

func NewUserService(repo repo.IUserRepo, userDetailRepo user_detail.Repository) *UserService {
	return &UserService{
		repo:           repo,
		userDetailRepo: userDetailRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (bson.ObjectID, error) {
	wc := writeconcern.Majority()
	txnOptions := options.Transaction().SetWriteConcern(wc)
	session, err := global.Client.StartSession()
	if err != nil {
		return bson.ObjectID{}, err
	}
	defer session.EndSession(context.TODO())
	result, err := session.WithTransaction(ctx, func(ctx context.Context) (any, error) {
		// 判断用户是否存在
		if _, isExist := s.repo.FindUserIsExist(ctx, user); isExist {
			return nil, errors.New("用户已存在")
		}
		id, err := s.repo.CreateUser(ctx, user)
		if err != nil {
			return nil, err
		}
		// 创建用户详情
		ud := user_detail.UserDetail{
			UserID:      id,
			NickName:    "浩瀚星河",
			Avatar:      "https://github.com/account",
			Description: "低级的欲望通过放纵就可获得；高级的欲望通过自律方可获得；顶级的欲望通过煎熬才可获得。所谓自由，不是随心所欲，而是自我主宰。",
			Email:       "email@codepzj.cn",
		}
		fmt.Println(s.userDetailRepo)
		err = s.userDetailRepo.CreateUserDetail(ctx, &ud)
		if err != nil {
			return nil, err
		}
		// 为用户添加权限
		added, err := global.Enforcer.AddRoleForUser(id.Hex(), roleIdConvertToString(*user.RoleId))
		if added && err == nil {
			return id, nil
		}
		return bson.ObjectID{}, err
	}, txnOptions)
	if err != nil {
		return bson.ObjectID{}, err
	}
	if bid, ok := result.(bson.ObjectID); ok {
		return bid, nil
	}
	return bson.ObjectID{}, errors.New("不是有效的ObjectId类型")
}

func (s *UserService) FindUserIsExist(ctx context.Context, user *domain.User) (*UserDto, bool) {
	u, ok := s.repo.FindUserIsExist(ctx, user)
	if !ok {
		return nil, false
	}

	// TODO: 为了开发环境方便，生产环境应该让后台初始化账号密码
	// 如果密码为初始化密码，跳过校验
	if user.Password == "123456" && u.Password == "123456" {
		return DoToDTO(u), true
	}

	// bcrypt校验密码是否有效
	if utils.CompareHashAndPassword(u.Password, user.Password) {
		return DoToDTO(u), true
	}
	return nil, false
}

func (s *UserService) FindAllUsers(ctx context.Context) ([]*UserDto, error) {
	users, err := s.repo.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return DOsToDTOs(users), nil
}
