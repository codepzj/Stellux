package api

import (
	"github.com/samber/lo"
	"server/internal/posts/internal/service"
	"time"
)

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

func DTOToVO(posts *service.PostsDTO) *PostsVO {
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

func DTOsToVOs(posts []*service.PostsDTO) []*PostsVO {
	return lo.Map(posts, func(item *service.PostsDTO, _ int) *PostsVO {
		return DTOToVO(item)
	})
}
