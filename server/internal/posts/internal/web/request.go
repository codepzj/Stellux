package web

import "github.com/codepzj/Stellux/server/internal/posts/internal/domain"

type PostsReq struct {
	Title       string   `json:"title" binding:"required"`
	Content     string   `json:"content" binding:"required"`
	Author      string   `json:"author" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Category    string   `json:"category" binding:"required"`
	Tags        []string `json:"tags" binding:"required"`
	Cover       string   `json:"cover"`
	IsTop       *bool    `json:"is_top" binding:"required"`
	IsPublish   *bool    `json:"is_publish" binding:"required"`
}

type UpdatePublishStatusReq struct {
	ID        string `json:"id" binding:"required"`
	IsPublish *bool   `json:"is_publish" binding:"required"`
}

type PostIDReq struct {
	ID string `uri:"id" binding:"required"`
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
		IsPublish:   postsReq.IsPublish,
	}
}
