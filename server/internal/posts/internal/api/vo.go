package api

import (
	"server/internal/posts/internal/domain"
	"time"
)

type PostsVO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
}

func toPostsVO(posts *domain.Posts) *PostsVO {
	return &PostsVO{
		ID:        posts.ID.Hex(),
		CreatedAt: posts.CreatedAt,
		UpdatedAt: posts.UpdatedAt,
		Title:     posts.Title,
		Content:   posts.Content,
		Author:    posts.Author,
	}
}
