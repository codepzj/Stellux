package web

import "server/internal/posts/internal/domain"

type PostsReq struct {
	Title       string   `json:"title" binding:"required"`
	Content     string   `json:"content" binding:"required"`
	Author      string   `json:"author" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Category    string   `json:"category" binding:"required"`
	Tags        []string `json:"tags" binding:"required"`
	Cover       string   `json:"cover"`
	IsTop       bool     `json:"isTop"`
}

func toPosts(postsReq PostsReq) *domain.Posts {
	return &domain.Posts{
		Title:       postsReq.Title,
		Content:     postsReq.Content,
		Author:      postsReq.Author,
		Description: postsReq.Description,
		Category:    postsReq.Category,
		Tags:        postsReq.Tags,
		Cover:       postsReq.Cover,
		IsTop:       postsReq.IsTop,
	}
}
