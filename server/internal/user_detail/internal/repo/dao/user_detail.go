package dao

import (
	"context"

	"server/internal/user_detail/internal/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IUserDetailDao interface {
	CreateUserDetail(ctx context.Context, userDetail *domain.UserDetail) error
	UpdateUserDetail(ctx context.Context, userDetail *domain.UserDetail) error
}

var _ IUserDetailDao = (*UserDetailDao)(nil)

type UserDetailDao struct {
	userDetailColl *mongo.Collection
}

func NewUserDetailDao(db *mongo.Database) *UserDetailDao {
	return &UserDetailDao{userDetailColl: db.Collection("user_detail")}
}

func (u *UserDetailDao) CreateUserDetail(ctx context.Context, userDetail *domain.UserDetail) error {
	result, err := u.userDetailColl.InsertOne(ctx, userDetail)
	if err != nil {
		return err
	}
	if result.InsertedID == nil {
		return errors.New("insert failed")
	}
	return nil
}

func (u *UserDetailDao) UpdateUserDetail(ctx context.Context, userDetail *domain.UserDetail) error {
	filter := bson.M{"uid": userDetail.UserID}
	update := bson.D{{"$set", bson.D{{"nick_name", userDetail.NickName}, {"avatar", userDetail.Avatar}, {"description", userDetail.Description}, {"email", userDetail.Email}}}}
	result, err := u.userDetailColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("更新失败")
	}
	return nil
}
