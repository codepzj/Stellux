package repo

import (
	"context"

	"github.com/codepzj/Stellux/server/internal/file/internal/domain"
	"github.com/codepzj/Stellux/server/internal/file/internal/repo/dao"

	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type IFileRepo interface {
	CreateMany(ctx context.Context, files []*domain.File) error
	DeleteMany(ctx context.Context, files []*domain.File) error
	FindByPage(ctx context.Context, page int64, pageSize int64) ([]*domain.FileDTO, error)
	GetAllCount(ctx context.Context) (int64, error)
	DeleteByUid(ctx context.Context, uid string) error
}

type FileRepo struct {
	dao dao.IFileDao
}

var _ IFileRepo = (*FileRepo)(nil)

func NewFileRepo(fileDao dao.IFileDao) *FileRepo {
	return &FileRepo{fileDao}
}

func (p *FileRepo) CreateMany(ctx context.Context, files []*domain.File) error {
	return p.dao.Create(ctx, files)
}

func (p *FileRepo) DeleteMany(ctx context.Context, files []*domain.File) error {
	fileIds := lo.Map(files, func(file *domain.File, _ int) bson.ObjectID {
		return file.ID
	})
	return p.dao.Delete(ctx, fileIds)
}

func (p *FileRepo) FindByPage(ctx context.Context, page int64, pageSize int64) ([]*domain.FileDTO, error) {
	files, err := p.dao.FindListByPage(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}
	return domain.ToFilesDTO(files), nil
}

func (p *FileRepo) GetAllCount(ctx context.Context) (int64, error) {
	return p.dao.FindCount(ctx)
}

func (p *FileRepo) DeleteByUid(ctx context.Context, uid string) error {
	return p.dao.DeleteByUid(ctx, uid)
}
