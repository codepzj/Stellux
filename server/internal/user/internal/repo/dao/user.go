package dao

import (
	"context"
	"time"

	"server/internal/user/internal/domain"

	"go.mongodb.org/mongo-driver/v2/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/pkg/errors"
)

type IUserDao interface {
	CreateOne(ctx context.Context, user *domain.User) error
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

func NewUserDao(database *mongo.Database) *UserDao {
	return &UserDao{userColl: database.Collection("user")}
}

func (u *UserDao) CreateOne(ctx context.Context, user *domain.User) error {
	user.CreatedAt = time.Now().Local()
	user.UpdatedAt = time.Now().Local()
	_, err := u.userColl.InsertOne(ctx, user)
	if err != nil {
		return errors.Wrapf(err, "添加user失败,%+v", user)
	}
	return err
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
	defer cursor.Close(ctx)
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
	defer cursor.Close(ctx)

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
