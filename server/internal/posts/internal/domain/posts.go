package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Posts struct {
	ID          bson.ObjectID `bson:"_id"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
	DeletedAt   time.Time     `bson:"deleted_at,omitempty"`
	Title       string
	Content     string
	Author      string
	Description string
	Category    string
	Tags        []string
	Cover       string
	IsTop       bool `bson:"is_top"`
}

type PageDTO struct {
	// 当前页
	PageNo int64 `json:"page_no" binding:"required"`
	// 每页数量
	PageSize int64 `json:"page_size" binding:"required"`
	// 排序字段
	Field string `json:"sort_field,omitempty"`
	// 排序规则
	Order string `json:"sort_order,omitempty"`
	// 搜索内容
	Keyword string `json:"keyword,omitempty"`
}

func (p *PageDTO) OrderConvertToInt() int {
	switch p.Order {
	case "ASC":
		return 1
	case "DESC":
		return -1
	default:
		return -1
	}
}
