package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type UserDetail struct {
	UserID      bson.ObjectID `bson:"uid"`
	NickName    string        `bson:"nick_name"`
	Avatar      string        `bson:"avatar"`
	Description string        `bson:"description"`
	Email       string        `bson:"email"`
}
