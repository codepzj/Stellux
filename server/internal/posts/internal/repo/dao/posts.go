package dao

import (
	"context"
	"time"

	"github.com/codepzj/Stellux/server/global"
	"github.com/codepzj/Stellux/server/internal/posts/internal/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type IPostsDao interface {
	Create(ctx context.Context, posts *domain.Posts) error
	FindAll(ctx context.Context) ([]*domain.Posts, error)
	FindListByCondition(ctx context.Context, skip int64, limit int64, keyword string, field string, order int) ([]*domain.Posts, error)
	FindById(ctx context.Context, id bson.ObjectID) (*domain.Posts, error)
	GetAllCount(ctx context.Context) (int64, error)
	GetAllCountByKeyword(ctx context.Context, keyword string) (int64, error)
	FindOneAndUpdateStatus(ctx context.Context, id bson.ObjectID, isPublish *bool) error
	ResumePostById(ctx context.Context, id bson.ObjectID) error
	DeleteSoftById(ctx context.Context, id bson.ObjectID) error
	DeleteById(ctx context.Context, id bson.ObjectID) error
}

type PostsDao struct {
	postColl *mongo.Collection
}

var _ IPostsDao = (*PostsDao)(nil)

func NewPostsDao() *PostsDao {
	return &PostsDao{postColl: global.DB.Collection("posts")}
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

func (p *PostsDao) FindListByCondition(ctx context.Context, skip int64, limit int64, keyword string, field string, order int) ([]*domain.Posts, error) {
	findOptions := options.Find().SetSkip(skip).SetLimit(limit).SetSort(bson.M{field: order})
	// 查询搜索内容是否在title、description中出现
	filter := bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: "title", Value: bson.D{{Key: "$regex", Value: keyword}}}},
		bson.D{{Key: "description", Value: bson.D{{Key: "$regex", Value: keyword}}}},
	}}}
	cursor, err := p.postColl.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	var posts []domain.Posts
	if err := cursor.All(ctx, &posts); err != nil {
		return nil, err
	}
	return domain.ToPtr(posts), nil
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

func (p *PostsDao) FindOneAndUpdateStatus(ctx context.Context, id bson.ObjectID, isPublish *bool) error {
	result := p.postColl.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"is_publish": isPublish}})
	return result.Err()
}

func (p *PostsDao) DeleteSoftById(ctx context.Context, id bson.ObjectID) error {
	result := p.postColl.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"deleted_at": time.Now().Local()}})
	return result.Err()
}

func (p *PostsDao) ResumePostById(ctx context.Context, id bson.ObjectID) error {
	// 删除deleted_at字段
	result := p.postColl.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$unset": bson.M{"deleted_at": nil}})
	return result.Err()
}

func (p *PostsDao) DeleteById(ctx context.Context, id bson.ObjectID) error {
	result, err := p.postColl.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errors.Wrap(err, "删除失败")
	}
	if result.DeletedCount == 0 {
		return errors.Wrap(err, "删除条数为0，删除失败")
	}
	return nil
}
