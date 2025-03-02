package dao

import (
	"context"
	"server/internal/file/internal/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type IFileDao interface {
	Create(ctx context.Context, files []*domain.File) error
	Delete(ctx context.Context, filesId []bson.ObjectID) error
	FindListByPage(ctx context.Context, page int64, pageSize int64) ([]*domain.File, error)
	FindCount(ctx context.Context) (int64, error)
	DeleteByUid(ctx context.Context, uid string) error
}

type FileDao struct {
	fileColl *mongo.Collection
}

var _ IFileDao = (*FileDao)(nil)

func NewFileDao(database *mongo.Database) *FileDao {
	return &FileDao{
		fileColl: database.Collection("file"),
	}
}

func (p *FileDao) Create(ctx context.Context, files []*domain.File) error {
	result, err := p.fileColl.InsertMany(ctx, domain.GetFilesFromTime(files))
	if err != nil {
		return err
	}
	if result.InsertedIDs == nil {
		return err
	}
	return nil
}

func (p *FileDao) Delete(ctx context.Context, filesId []bson.ObjectID) error {
	result, err := p.fileColl.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": filesId}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.Wrap(err, "删除失败")
	}
	return nil
}

func (p *FileDao) FindListByPage(ctx context.Context, page int64, pageSize int64) ([]*domain.File, error) {
	// 设置分页查询
	skip := (page - 1) * pageSize
	limit := pageSize
	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)
	// 降序
	findOptions.SetSort(bson.M{"created_at": -1})

	cursor, err := p.fileColl.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, errors.Wrap(err, "查询文件失败")
	}
	var files []domain.File
	if err := cursor.All(ctx, &files); err != nil {
		return nil, errors.Wrap(err, "解码文件失败")
	}
	return domain.ToFilePtrSlice(files), nil
}

func (p *FileDao) FindCount(ctx context.Context) (int64, error) {
	return p.fileColl.CountDocuments(ctx, bson.M{})
}

func (p *FileDao) DeleteByUid(ctx context.Context, uid string) error {
	result, err := p.fileColl.DeleteOne(ctx, bson.M{"uid": uid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.Wrap(err, "删除失败，文件不存在")
	}
	return nil
}
