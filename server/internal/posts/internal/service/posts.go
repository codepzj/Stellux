package service

import (
	"context"
	"server/internal/posts/internal/domain"
	"server/internal/posts/internal/repo"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostsService interface {
	CreatePosts(ctx context.Context, posts *domain.Posts) error
	FindPostById(ctx context.Context, id bson.ObjectID) (*domain.PostsDTO, error) // 修改返回类型为 *domain.PostsDTO
	FindAllPosts(ctx context.Context) ([]*domain.PostsDTO, error)
	FindPostsByPage(ctx context.Context, pageDTO *domain.PageDTO) ([]*domain.PostsDTO, int64, int64, error)
}

type PostsService struct {
	repo repo.IPostsRepo
}

var _ IPostsService = (*PostsService)(nil)

func NewPostsService(repo repo.IPostsRepo) *PostsService {
	return &PostsService{repo}
}

func (p *PostsService) CreatePosts(ctx context.Context, posts *domain.Posts) error {
	return p.repo.CreatePost(ctx, posts)
}

func (p *PostsService) FindPostById(ctx context.Context, id bson.ObjectID) (*domain.PostsDTO, error) { // 修改返回类型为 *domain.PostsDTO
	posts, err := p.repo.FindPostById(ctx, id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostsService) FindAllPosts(ctx context.Context) ([]*domain.PostsDTO, error) {
	return p.repo.FindAllPosts(ctx)
}

func (p *PostsService) FindPostsByPage(ctx context.Context, pageDTO *domain.PageDTO) ([]*domain.PostsDTO, int64, int64, error) {
	return p.repo.FindPostsByPage(ctx, pageDTO)
}
