package dao

import (
	"context"
	"server/global"
	"time"

	"server/internal/user/internal/domain"

	"go.mongodb.org/mongo-driver/v2/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/pkg/errors"
)

type IUserDao interface {
	CreateOne(ctx context.Context, user *domain.User) (bson.ObjectID, error)
	FindByName(ctx context.Context, username string) (*domain.User, error)
	FindById(ctx context.Context, userId bson.ObjectID) (*domain.User, error)
	FindAll(ctx context.Context) ([]*domain.User, error)
	FindAllByRoleID(ctx context.Context, roleId int) ([]*domain.User, error)
	FindUserIsExist(ctx context.Context, user *domain.User) (*domain.User, bool)
}

var _ IUserDao = (*UserDao)(nil)

type UserDao struct {
	userColl *mongo.Collection
}

func NewUserDao() *UserDao {
	return &UserDao{userColl: global.DB.Collection("user")}
}

func (u *UserDao) CreateOne(ctx context.Context, user *domain.User) (bson.ObjectID, error) {
	user.ID = bson.NewObjectID()
	user.CreatedAt = time.Now().Local()
	user.UpdatedAt = time.Now().Local()
	result, err := u.userColl.InsertOne(ctx, user)
	if err != nil {
		return bson.ObjectID{}, err
	}
	id, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return bson.ObjectID{}, err
	}
	return id, nil
}

func (u *UserDao) FindByName(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := u.userColl.FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDao) FindById(ctx context.Context, userId bson.ObjectID) (*domain.User, error) {
	var user domain.User
	err := u.userColl.FindOne(ctx, bson.D{{Key: "_id", Value: userId}}).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDao) FindAll(ctx context.Context) ([]*domain.User, error) {
	cursor, err := u.userColl.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)
	var users []domain.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return domain.ToPtr(users), nil
}

func (u *UserDao) FindAllByRoleID(ctx context.Context, roleId int) ([]*domain.User, error) {
	cursor, err := u.userColl.Find(ctx, bson.D{{Key: "role_id", Value: bson.M{"$gte": roleId}}})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	var users []domain.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return domain.ToPtr(users), nil
}

func (u *UserDao) FindUserIsExist(ctx context.Context, user *domain.User) (*domain.User, bool) {
	var usr domain.User
	err := u.userColl.FindOne(ctx, bson.D{{Key: "username", Value: user.Username}, {Key: "password", Value: user.Password}}).Decode(&usr)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, false
	} else if err != nil {
		return nil, false
	}

	return &usr, true
}
