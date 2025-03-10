package service

import (
	"context"

	"github.com/codepzj/Stellux/server/internal/pkg/wrap"
	"github.com/codepzj/Stellux/server/internal/posts/internal/domain"
	"github.com/codepzj/Stellux/server/internal/posts/internal/repo"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostsService interface {
	CreatePosts(ctx context.Context, posts *domain.Posts) error
	FindPostById(ctx context.Context, id bson.ObjectID) (*PostsDTO, error)
	FindAllPosts(ctx context.Context) ([]*PostsDTO, error)
	FindPostsByCondition(ctx context.Context, page *wrap.Page) ([]*PostsDTO, int64, int64, error)
	UpdatePostsPublishStatus(ctx context.Context, id bson.ObjectID, isPublish *bool) error
	DeletePostSoftById(ctx context.Context, id bson.ObjectID) error
	ResumePostSoftById(ctx context.Context, id bson.ObjectID) error
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

func (p *PostsService) UpdatePostsPublishStatus(ctx context.Context, id bson.ObjectID, isPublish *bool) error {
	return p.repo.UpdateStatus(ctx, id, isPublish)
}

func (p *PostsService) DeletePostSoftById(ctx context.Context, id bson.ObjectID) error {
	return p.repo.DeletePostSoftById(ctx, id)
}
func (p *PostsService) ResumePostSoftById(ctx context.Context, id bson.ObjectID) error {
	return p.repo.ResumePostById(ctx, id)
}
