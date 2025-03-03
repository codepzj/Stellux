package domain

import (
	"time"

	"github.com/samber/lo"
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

type PostsDTO struct {
	ID          bson.ObjectID `bson:"_id"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
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

type PostsVO struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Tags        []string  `json:"tags"`
	Cover       string    `json:"cover"`
	IsTop       bool      `json:"is_top"`
}

func ToPostsPtr(posts []Posts) []*Posts {
	return lo.Map(posts, func(item Posts, _ int) *Posts {
		return &item
	})
}

func ToPostsDTO(posts *Posts) *PostsDTO {
	return &PostsDTO{
		ID:          posts.ID,
		CreatedAt:   posts.CreatedAt,
		UpdatedAt:   posts.UpdatedAt,
		Title:       posts.Title,
		Content:     posts.Content,
		Author:      posts.Author,
		Description: posts.Description,
		Category:    posts.Category,
		Tags:        posts.Tags,
		Cover:       posts.Cover,
		IsTop:       posts.IsTop,
	}
}

func ToPostsDTOs(posts []*Posts) []*PostsDTO {
	return lo.Map(posts, func(item *Posts, _ int) *PostsDTO {
		return &PostsDTO{
			ID:          item.ID,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			Title:       item.Title,
			Content:     item.Content,
			Author:      item.Author,
			Description: item.Description,
			Category:    item.Category,
			Tags:        item.Tags,
			Cover:       item.Cover,
			IsTop:       item.IsTop,
		}
	})
}

func ToPostsVO(posts *PostsDTO) *PostsVO {
	return &PostsVO{
		ID:          posts.ID.Hex(),
		CreatedAt:   posts.CreatedAt,
		UpdatedAt:   posts.UpdatedAt,
		Title:       posts.Title,
		Content:     posts.Content,
		Author:      posts.Author,
		Description: posts.Description,
		Category:    posts.Category,
		Tags:        posts.Tags,
		Cover:       posts.Cover,
		IsTop:       posts.IsTop,
	}
}

func ToPostsVOs(posts []*PostsDTO) []*PostsVO {
	return lo.Map(posts, func(item *PostsDTO, _ int) *PostsVO {
		return ToPostsVO(item)
	})
}
