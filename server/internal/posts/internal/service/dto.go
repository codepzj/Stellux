package service

import (
	"github.com/codepzj/Stellux/server/internal/pkg/wrap"
	"github.com/codepzj/Stellux/server/internal/posts/internal/domain"
	"time"

	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

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
	IsTop       *bool `bson:"is_top"`
	IsPublish   *bool `bson:"is_publish"`
}

type PageDTO struct {
	// 当前页
	PageNo int64
	// 每页数量
	PageSize int64
	// 排序字段
	Field string
	// 排序规则
	Order string
	// 搜索内容
	Keyword string
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

func PageToPageDTO(page *wrap.Page) *PageDTO {
	if page.Field == "" {
		page.Field = "created_at"
		page.Order = "DESC"
	}
	return &PageDTO{
		PageNo:   page.PageNo,
		PageSize: page.Size,
		Field:    page.Field,
		Order:    page.Order,
		Keyword:  page.Keyword,
	}
}

func DoToDTO(posts *domain.Posts) *PostsDTO {
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
		IsPublish:   posts.IsPublish,
	}
}

func DOsToDTOs(posts []*domain.Posts) []*PostsDTO {
	return lo.Map(posts, func(post *domain.Posts, _ int) *PostsDTO {
		return DoToDTO(post)
	})
}
