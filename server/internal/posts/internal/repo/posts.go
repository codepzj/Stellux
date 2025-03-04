package repo

import (
	"context"
	"server/internal/posts/internal/domain"
	"server/internal/posts/internal/repo/dao"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostsRepo interface {
	CreatePost(ctx context.Context, posts *domain.Posts) error
	FindPostById(ctx context.Context, id bson.ObjectID) (*domain.Posts, error)
	FindAllPosts(ctx context.Context) ([]*domain.Posts, error)
	FindPostsByCondition(ctx context.Context, pageNo int64, pageSize int64, keyword string, field string, order int) ([]*domain.Posts, int64, int64, error)
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

func (p *PostsRepo) FindPostById(ctx context.Context, id bson.ObjectID) (*domain.Posts, error) {
	return p.dao.FindById(ctx, id)
}

func (p *PostsRepo) FindAllPosts(ctx context.Context) ([]*domain.Posts, error) {
	return p.dao.FindAll(ctx)
}

func (p *PostsRepo) FindPostsByCondition(ctx context.Context, pageNo int64, pageSize int64, keyword string, field string, order int) ([]*domain.Posts, int64, int64, error) {
	totalCount, err := p.GetAllCountByKeyword(ctx, keyword)
	if err != nil {
		return nil, 0, 0, err
	}
	totalPage := (totalCount + pageSize - 1) / pageSize
	list, err := p.dao.FindListByCondition(ctx, (pageNo-1)*pageSize, pageSize, keyword, field, order)
	if err != nil {
		return nil, 0, 0, err
	}
	return list, totalCount, totalPage, nil
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
