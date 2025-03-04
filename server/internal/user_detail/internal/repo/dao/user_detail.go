package dao

import (
	"context"

	"server/internal/user_detail/internal/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IUserDetailDao interface {
	CreateOne(ctx context.Context, userDetail *domain.UserDetail) error
	UpdateOne(ctx context.Context, userDetail *domain.UserDetail) error
}

var _ IUserDetailDao = (*UserDetailDao)(nil)

type UserDetailDao struct {
	userDetailColl *mongo.Collection
}

func NewUserDetailDao(db *mongo.Database) *UserDetailDao {
	return &UserDetailDao{userDetailColl: db.Collection("user_detail")}
}

func (u *UserDetailDao) CreateOne(ctx context.Context, userDetail *domain.UserDetail) error {
	result, err := u.userDetailColl.InsertOne(ctx, userDetail)
	if err != nil {
		return err
	}
	if result.InsertedID == nil {
		return errors.New("insert failed")
	}
	return nil
}

func (u *UserDetailDao) UpdateOne(ctx context.Context, userDetail *domain.UserDetail) error {
	result, err := u.userDetailColl.UpdateOne(ctx, bson.M{"user_id": userDetail.UserID}, bson.M{"$set": userDetail})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("update failed")
	}
	return nil
}
