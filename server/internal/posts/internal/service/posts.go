package service

import (
	"context"
	"server/internal/pkg/wrap"
	"server/internal/posts/internal/domain"
	"server/internal/posts/internal/repo"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostsService interface {
	CreatePosts(ctx context.Context, posts *domain.Posts) error
	FindPostById(ctx context.Context, id bson.ObjectID) (*PostsDTO, error)
	FindAllPosts(ctx context.Context) ([]*PostsDTO, error)
	FindPostsByCondition(ctx context.Context, page *wrap.Page) ([]*PostsDTO, int64, int64, error)
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

func (p *PostsService) FindPostById(ctx context.Context, id bson.ObjectID) (*PostsDTO, error) {
	posts, err := p.repo.FindPostById(ctx, id)
	if err != nil {
		return nil, err
	}
	return DoToDTO(posts), nil
}

func (p *PostsService) FindAllPosts(ctx context.Context) ([]*PostsDTO, error) {
	post, err := p.repo.FindAllPosts(ctx)
	if err != nil {
		return nil, err
	}
	return DOsToDTOs(post), nil
}

func (p *PostsService) FindPostsByCondition(ctx context.Context, page *wrap.Page) ([]*PostsDTO, int64, int64, error) {
	pageDTO := PageToPageDTO(page)
	posts, totalCount, totalPage, err := p.repo.FindPostsByCondition(ctx, pageDTO.PageNo, pageDTO.PageSize, pageDTO.Keyword, pageDTO.Field, pageDTO.OrderConvertToInt())
	if err != nil {
		return make([]*PostsDTO, 0), 0, 0, err
	}

	return DOsToDTOs(posts), totalCount, totalPage, nil

}
