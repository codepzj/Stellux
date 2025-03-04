package domain

import (
	"time"

	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// User 用户表 RoleId 为 0 时为管理员，为 1 时为普通用户
type User struct {
	ID        bson.ObjectID `bson:"_id"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
	DeletedAt time.Time     `bson:"deleted_at,omitempty"`
	Username  string        `bson:"username"`
	Password  string        `bson:"password"`
	RoleId    int           `bson:"role_id"`
}

func ToPtr(users []User) []*User {
	return lo.Map(users, func(item User, index int) *User {
		return &item
	})
}
