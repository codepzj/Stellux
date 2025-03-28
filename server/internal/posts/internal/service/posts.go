package service

import (
	"context"

	"github.com/codepzj/Stellux/server/internal/pkg/wrap"
	"github.com/codepzj/Stellux/server/internal/posts/internal/domain"
	"github.com/codepzj/Stellux/server/internal/posts/internal/repo"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostsService interface {
	FindPostById(ctx context.Context, id bson.ObjectID) (*PostsDTO, error)
	FindAllPosts(ctx context.Context) ([]*PostsDTO, error)
	FindPostByCondition(ctx context.Context, page *wrap.Page) ([]*PostsDTO, int64, int64, error)
	AdminFindPostByCondition(ctx context.Context, page *wrap.Page) ([]*PostsDTO, int64, int64, error)
	AdminUpdatePostStatus(ctx context.Context, id bson.ObjectID, isPublish *bool) error
	AdminDeletePostSoftById(ctx context.Context, id bson.ObjectID) error
	AdminCreatePost(ctx context.Context, posts *domain.Posts) error
	AdminUpdatePost(ctx context.Context, posts *domain.Posts) error
	AdminResumePostSoftById(ctx context.Context, id bson.ObjectID) error
}

type PostsService struct {
	repo repo.IPostsRepo
}

var _ IPostsService = (*PostsService)(nil)

func NewPostsService(repo repo.IPostsRepo) *PostsService {
	return &PostsService{repo}
}

// FindPostById 获取特定文章
func (p *PostsService) FindPostById(ctx context.Context, id bson.ObjectID) (*PostsDTO, error) {
	posts, err := p.repo.FindPostById(ctx, id)
	if err != nil {
		return nil, err
	}
	return DoToDTO(posts), nil
}

// FindAllPosts 获取所有文章
func (p *PostsService) FindAllPosts(ctx context.Context) ([]*PostsDTO, error) {
	post, err := p.repo.FindAllPosts(ctx)
	if err != nil {
		return nil, err
	}
	return DOsToDTOs(post), nil
}

// FindPostByCondition 根据条件获取文章（分页，排序，关键词）
func (p *PostsService) FindPostByCondition(ctx context.Context, page *wrap.Page) ([]*PostsDTO, int64, int64, error) {
	pageDTO := PageToPageDTO(page)
	posts, totalCount, totalPage, err := p.repo.FindPostsByCondition(ctx, pageDTO.PageNo, pageDTO.PageSize, pageDTO.Keyword, pageDTO.Field, pageDTO.OrderConvertToInt())
	if err != nil {
		return make([]*PostsDTO, 0), 0, 0, err
	}

	return DOsToDTOs(posts), totalCount, totalPage, nil

}

// AdminFindPostByCondition 管理员根据条件获取文章（分页，排序，关键词）
func (p *PostsService) AdminFindPostByCondition(ctx context.Context, page *wrap.Page) ([]*PostsDTO, int64, int64, error) {
	pageDTO := PageToPageDTO(page)
	posts, totalCount, totalPage, err := p.repo.AdminFindPostsByCondition(ctx, pageDTO.PageNo, pageDTO.PageSize, pageDTO.Keyword, pageDTO.Field, pageDTO.OrderConvertToInt())
	if err != nil {
		return make([]*PostsDTO, 0), 0, 0, err
	}
	return DOsToDTOs(posts), totalCount, totalPage, nil
}

// AdminCreatePost 管理员创建文章
func (p *PostsService) AdminCreatePost(ctx context.Context, posts *domain.Posts) error {
	return p.repo.AdminCreatePost(ctx, posts)
}

// AdminUpdatePost 管理员更新文章
func (p *PostsService) AdminUpdatePost(ctx context.Context, posts *domain.Posts) error {
	return p.repo.AdminUpdatePost(ctx, posts)
}

// AdminUpdatePostStatus 管理员上下架文章
func (p *PostsService) AdminUpdatePostStatus(ctx context.Context, id bson.ObjectID, isPublish *bool) error {
	return p.repo.AdminUpdatePostStatus(ctx, id, isPublish)
}

// AdminDeletePostSoftById 管理员软删除文章
func (p *PostsService) AdminDeletePostSoftById(ctx context.Context, id bson.ObjectID) error {
	return p.repo.AdminDeletePostSoftById(ctx, id)
}

// AdminResumePostSoftById 管理员恢复文章
func (p *PostsService) AdminResumePostSoftById(ctx context.Context, id bson.ObjectID) error {
	return p.repo.AdminResumePostById(ctx, id)
}
