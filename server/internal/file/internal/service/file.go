package service

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"server/internal/file/internal/domain"
	"server/internal/file/internal/repo"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/google/uuid"
)

type IFileService interface {
	UploadFilesLocal(ctx context.Context, uids []string, fileNames []string, files []*multipart.FileHeader) ([]*domain.File, error)
	GetPhotosByPage(ctx context.Context, page int64, pageSize int64) ([]*domain.FileDTO, int64, int64, error)
	DeletePhotoByUid(ctx context.Context, uid string) error
}

type FileService struct {
	repo repo.IFileRepo
}

var _ IFileService = (*FileService)(nil)

func NewFileService(fileRepo repo.IFileRepo) *FileService {
	return &FileService{fileRepo}
}

func (s *FileService) UploadFilesLocal(ctx context.Context, uids []string, fileNames []string, files []*multipart.FileHeader) ([]*domain.File, error) {
	var pictures []*domain.File
	os.MkdirAll("static/images", os.ModePerm)
	for i := range files {
		uid := uids[i]
		fileName := fileNames[i]
		fileNameList := strings.Split(fileName, ".")
		// 生成新的文件名
		newFileName := fileNameList[0] + strings.ReplaceAll(uuid.New().String(), "-", "") + "." + fileNameList[1]
		networkPath := "/images/" + newFileName
		filePath := "static/images/" + newFileName
		picture := &domain.File{
			Uid:       uid,
			FileName:  fileName,
			Type:      "local",
			Url:       networkPath,
			Dst:       filePath,
			CreatedAt: time.Now().Local(),
		}
		pictures = append(pictures, picture)
	}
	// 保存到数据库
	if err := s.repo.CreateMany(ctx, pictures); err != nil {
		// 回滚
		for _, picture := range pictures {
			os.Remove(picture.Dst)
		}
		s.repo.DeleteMany(ctx, pictures)
		return nil, errors.Wrap(err, "保存到数据库失败")
	}

	// 保存文件到本地
	for i, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, errors.Wrap(err, "打开文件失败")
		}
		defer file.Close()
		// 创建目标文件
		dst, err := os.Create(pictures[i].Dst)
		if err != nil {
			return nil, errors.Wrap(err, "创建目标文件失败")
		}
		defer dst.Close()
		// 保存文件
		if _, err := io.Copy(dst, file); err != nil {
			return nil, errors.Wrap(err, "保存文件失败")
		}
	}

	return pictures, nil
}

func (s *FileService) GetPhotosByPage(ctx context.Context, page int64, pageSize int64) ([]*domain.FileDTO, int64, int64, error) {
	photos, err := s.repo.FindByPage(ctx, page, pageSize)
	if err != nil {
		return nil, 0, 0, err
	}
	totalCount, err := s.repo.GetAllCount(ctx)
	if err != nil {
		return nil, 0, 0, err
	}
	totalPage := totalCount / pageSize
	if totalCount%pageSize != 0 {
		totalPage++
	}
	return photos, totalCount, totalPage, nil
}

func (s *FileService) DeletePhotoByUid(ctx context.Context, uid string) error {
	return s.repo.DeleteByUid(ctx, uid)
}
