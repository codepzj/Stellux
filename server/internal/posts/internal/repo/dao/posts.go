package dao

import (
	"context"
	"server/internal/posts/internal/domain"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type IPostsDao interface {
	Create(ctx context.Context, posts *domain.Posts) error
	FindAll(ctx context.Context) ([]*domain.Posts, error)
	FindListByPage(ctx context.Context, pageDTO *domain.PageDTO) ([]*domain.Posts, error)
	FindById(ctx context.Context, id bson.ObjectID) (*domain.Posts, error)
	GetAllCount(ctx context.Context) (int64, error)
	GetAllCountByKeyword(ctx context.Context, keyword string) (int64, error)
	DeleteById(ctx context.Context, Id bson.ObjectID) error
}

type PostsDao struct {
	postColl *mongo.Collection
}

var _ IPostsDao = (*PostsDao)(nil)

func NewPostsDao(database *mongo.Database) *PostsDao {
	return &PostsDao{postColl: database.Collection("posts")}
}

func (p *PostsDao) Create(ctx context.Context, posts *domain.Posts) error {
	posts.ID = bson.NewObjectID()
	posts.CreatedAt = time.Now()
	posts.UpdatedAt = time.Now()
	_, err := p.postColl.InsertOne(ctx, posts)
	if err != nil {
		return errors.Wrap(err, "添加文章失败")
	}
	return nil
}

func (p *PostsDao) FindById(ctx context.Context, id bson.ObjectID) (*domain.Posts, error) {
	var posts domain.Posts
	err := p.postColl.FindOne(ctx, bson.M{"_id": id}).Decode(&posts)
	if err != nil {
		return nil, errors.Wrap(err, "查询文章失败")
	}
	return &posts, nil
}

func (p *PostsDao) FindAll(ctx context.Context) ([]*domain.Posts, error) {
	cursor, err := p.postColl.Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, "查询文章失败")
	}
	var posts []*domain.Posts
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, errors.Wrap(err, "查询文章失败")
	}
	return posts, nil
}

func (p *PostsDao) FindListByPage(ctx context.Context, pageDTO *domain.PageDTO) ([]*domain.Posts, error) {
	skip := (pageDTO.PageNo - 1) * pageDTO.PageSize
	limit := pageDTO.PageSize
	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)
	// 查询搜索内容是否在title、description中出现
	if pageDTO.Field == "" {
		pageDTO.Field = "created_at"
	}
	filter := bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: "title", Value: bson.D{{Key: "$regex", Value: pageDTO.Keyword}}}},
		bson.D{{Key: "description", Value: bson.D{{Key: "$regex", Value: pageDTO.Keyword}}}},
	}}}
	cursor, err := p.postColl.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	var posts []domain.Posts
	if err := cursor.All(ctx, &posts); err != nil {
		return nil, err
	}
	return domain.ToPostsPtr(posts), nil
}

func (p *PostsDao) GetAllCount(ctx context.Context) (int64, error) {
	return p.postColl.CountDocuments(ctx, bson.M{})
}

func (p *PostsDao) GetAllCountByKeyword(ctx context.Context, keyword string) (int64, error) {
	// 查询搜索内容是否在title、description、content中出现
	filter := bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: "title", Value: bson.D{{Key: "$regex", Value: keyword}}}},
		bson.D{{Key: "description", Value: bson.D{{Key: "$regex", Value: keyword}}}},
	}}}
	return p.postColl.CountDocuments(ctx, filter)
}

func (p *PostsDao) DeleteById(ctx context.Context, Id bson.ObjectID) error {
	result, err := p.postColl.DeleteOne(ctx, bson.M{"_id": Id})
	if err != nil {
		return errors.Wrap(err, "删除失败")
	}
	if result.DeletedCount == 0 {
		return errors.Wrap(err, "删除条数为0，删除失败")
	}
	return nil
}
