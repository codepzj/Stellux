package web

import (
	"github.com/codepzj/Stellux/server/internal/posts/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostsReq struct {
	Title       string   `json:"title" binding:"required"`
	Content     string   `json:"content" binding:"required"`
	Author      string   `json:"author" binding:"required"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Cover       string   `json:"cover"`
	IsTop       *bool    `json:"is_top" binding:"required"`
	IsPublish   *bool    `json:"is_publish" binding:"required"`
}

type UpdatePostReq struct {
	ID string `json:"id" binding:"required"`
	PostsReq
}

type UpdatePublishStatusReq struct {
	ID        string `json:"id" binding:"required"`
	IsPublish *bool  `json:"is_publish" binding:"required"`
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

func updatePostReqToPosts(updatePostReq UpdatePostReq) *domain.Posts {
	idObj, _ := bson.ObjectIDFromHex(updatePostReq.ID)
	return &domain.Posts{
		ID:          idObj,
		Title:       updatePostReq.Title,
		Content:     updatePostReq.Content,
		Author:      updatePostReq.Author,
		Description: updatePostReq.Description,
		Category:    updatePostReq.Category,
		Tags:        updatePostReq.Tags,
		Cover:       updatePostReq.Cover,
		IsTop:       updatePostReq.IsTop,
		IsPublish:   updatePostReq.IsPublish,
	}
}
