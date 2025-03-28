package repo

import (
	"context"

	"github.com/codepzj/Stellux/server/internal/posts/internal/domain"
	"github.com/codepzj/Stellux/server/internal/posts/internal/repo/dao"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostsRepo interface {
	FindPostById(ctx context.Context, id bson.ObjectID) (*domain.Posts, error)
	FindAllPosts(ctx context.Context) ([]*domain.Posts, error)
	FindPostsByCondition(ctx context.Context, pageNo int64, pageSize int64, keyword string, field string, order int) ([]*domain.Posts, int64, int64, error)
	GetAllCount(ctx context.Context) (int64, error)
	GetAllCountByKeyword(ctx context.Context, keyword string) (int64, error)
	AdminFindPostsByCondition(ctx context.Context, pageNo int64, pageSize int64, keyword string, field string, order int) ([]*domain.Posts, int64, int64, error)
	AdminGetAllCountByKeyword(ctx context.Context, keyword string) (int64, error)
	AdminCreatePost(ctx context.Context, posts *domain.Posts) error
	AdminUpdatePost(ctx context.Context, posts *domain.Posts) error
	AdminUpdatePostStatus(ctx context.Context, id bson.ObjectID, isPublish *bool) error
	AdminResumePostById(ctx context.Context, id bson.ObjectID) error
	AdminDeletePostSoftById(ctx context.Context, id bson.ObjectID) error
	AdminDeletePostById(ctx context.Context, id bson.ObjectID) error
}

type PostsRepo struct {
	dao dao.IPostsDao
}

var _ IPostsRepo = (*PostsRepo)(nil)

func NewPostsRepo(dao dao.IPostsDao) *PostsRepo {
	return &PostsRepo{dao}
}

// FindPostById 获取特定文章
func (p *PostsRepo) FindPostById(ctx context.Context, id bson.ObjectID) (*domain.Posts, error) {
	return p.dao.FindById(ctx, id)
}

// FindAllPosts 获取所有文章
func (p *PostsRepo) FindAllPosts(ctx context.Context) ([]*domain.Posts, error) {
	return p.dao.FindAll(ctx)
}

// FindPostsByCondition 根据条件获取文章（分页，排序，关键词）
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

// GetAllCount 获取文章总数
func (p *PostsRepo) GetAllCount(ctx context.Context) (int64, error) {
	return p.dao.GetAllCount(ctx)
}

// GetAllCountByKeyword 用户通过关键词获取文章总数
func (p *PostsRepo) GetAllCountByKeyword(ctx context.Context, keyword string) (int64, error) {
	return p.dao.GetAllCountByKeyword(ctx, keyword)
}

// AdminFindPostsByCondition 管理员分页，关键词查询文章列表，筛除删除的文章,不区分是否发布
func (p *PostsRepo) AdminFindPostsByCondition(ctx context.Context, pageNo int64, pageSize int64, keyword string, field string, order int) ([]*domain.Posts, int64, int64, error) {
	totalCount, err := p.AdminGetAllCountByKeyword(ctx, keyword)
	if err != nil {
		return nil, 0, 0, err
	}
	totalPage := (totalCount + pageSize - 1) / pageSize
	list, err := p.dao.AdminFindListByCondition(ctx, (pageNo-1)*pageSize, pageSize, keyword, field, order)
	if err != nil {
		return nil, 0, 0, err
	}
	return list, totalCount, totalPage, nil
}

// AdminGetAllCountByKeyword 管理员通过关键词获取文章总数
func (p *PostsRepo) AdminGetAllCountByKeyword(ctx context.Context, keyword string) (int64, error) {
	return p.dao.AdminGetAllCountByKeyword(ctx, keyword)
}

// AdminCreatePost 管理员创建文章
func (p *PostsRepo) AdminCreatePost(ctx context.Context, posts *domain.Posts) error {
	return p.dao.AdminCreate(ctx, posts)
}

// AdminUpdatePost 管理员更新文章
func (p *PostsRepo) AdminUpdatePost(ctx context.Context, posts *domain.Posts) error {
	return p.dao.AdminUpdate(ctx, posts)
}

// AdminUpdatePostStatus 管理员上下架文章
func (p *PostsRepo) AdminUpdatePostStatus(ctx context.Context, id bson.ObjectID, isPublish *bool) error {
	return p.dao.AdminFindOneAndUpdateStatus(ctx, id, isPublish)
}

// AdminDeletePostSoftById 管理员软删除文章
func (p *PostsRepo) AdminDeletePostSoftById(ctx context.Context, id bson.ObjectID) error {
	return p.dao.AdminDeleteSoftById(ctx, id)
}

// AdminResumePostById 管理员恢复文章
func (p *PostsRepo) AdminResumePostById(ctx context.Context, id bson.ObjectID) error {
	return p.dao.AdminResumePostById(ctx, id)
}

// AdminDeletePostById 管理员硬删除文章
func (p *PostsRepo) AdminDeletePostById(ctx context.Context, id bson.ObjectID) error {
	return p.dao.AdminDeletePostById(ctx, id)
}
