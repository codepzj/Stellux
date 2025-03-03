package repo

import (
	"context"
	"server/internal/posts/internal/domain"
	"server/internal/posts/internal/repo/dao"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostsRepo interface {
	CreatePost(ctx context.Context, posts *domain.Posts) error
	FindPostById(ctx context.Context, id bson.ObjectID) (*domain.PostsDTO, error)
	FindAllPosts(ctx context.Context) ([]*domain.PostsDTO, error)
	FindPostsByPage(ctx context.Context, pageDTO *domain.PageDTO) ([]*domain.PostsDTO, int64, int64, error)
	GetAllCount(ctx context.Context) (int64, error)
	GetAllCountByKeyword(ctx context.Context, keyword string) (int64, error)
	DeletePostById(ctx context.Context, Id bson.ObjectID) error
}

type PostsRepo struct {
	dao dao.IPostsDao
}

var _ IPostsRepo = (*PostsRepo)(nil)

func NewPostsRepo(dao dao.IPostsDao) *PostsRepo {
	return &PostsRepo{dao}
}

func (p *PostsRepo) CreatePost(ctx context.Context, posts *domain.Posts) error {
	return p.dao.Create(ctx, posts)
}

func (p *PostsRepo) FindPostById(ctx context.Context, id bson.ObjectID) (*domain.PostsDTO, error) {
	posts, err := p.dao.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.ToPostsDTO(posts), nil
}

func (p *PostsRepo) FindAllPosts(ctx context.Context) ([]*domain.PostsDTO, error) {
	posts, err := p.dao.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return domain.ToPostsDTOs(posts), nil
}

func (p *PostsRepo) FindPostsByPage(ctx context.Context, pageDTO *domain.PageDTO) ([]*domain.PostsDTO, int64, int64, error) {
	total_count, err := p.GetAllCountByKeyword(ctx, pageDTO.Keyword)
	if err != nil {
		return nil, 0, 0, err
	}
	total_page := total_count / pageDTO.PageSize
	if total_count%pageDTO.PageSize != 0 {
		total_page++
	}
	list, err := p.dao.FindListByPage(ctx, pageDTO)
	if err != nil {
		return nil, 0, 0, err
	}
	return domain.ToPostsDTOs(list), total_count, total_page, nil
}

func (p *PostsRepo) GetAllCount(ctx context.Context) (int64, error) {
	return p.dao.GetAllCount(ctx)
}

func (p *PostsRepo) GetAllCountByKeyword(ctx context.Context, keyword string) (int64, error) {
	return p.dao.GetAllCountByKeyword(ctx, keyword)
}

func (p *PostsRepo) DeletePostById(ctx context.Context, Id bson.ObjectID) error {
	return p.dao.DeleteById(ctx, Id)
}
