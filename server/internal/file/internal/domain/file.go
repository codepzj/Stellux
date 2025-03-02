package domain

import (
	"time"

	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type File struct {
	ID        bson.ObjectID `bson:"_id"`
	Uid       string        `bson:"uid"`
	FileName  string        `bson:"file_name"`
	CreatedAt time.Time     `bson:"created_at"`
	Type      string        `bson:"type"`
	Url       string        `bson:"url"`
	Dst       string        `bson:"dst"`
}

type FileDTO struct {
	ID        bson.ObjectID `json:"id"`
	Uid       string        `json:"uid"`
	FileName  string        `json:"file_name"`
	CreatedAt time.Time     `json:"created_at"`
	Type      string        `json:"type"`
	Url       string        `json:"url"`
}

func GetFileFromTime(file *File) *File {
	return &File{
		ID:        bson.NewObjectID(),
		Uid:       file.Uid,
		FileName:  file.FileName,
		CreatedAt: time.Now().Local(),
		Type:      file.Type,
		Url:       file.Url,
		Dst:       file.Dst,
	}
}
func GetFilesFromTime(files []*File) []*File {
	return lo.Map(files, func(item *File, _ int) *File {
		return GetFileFromTime(item)
	})
}

func ToFilePtrSlice(files []File) []*File {
	return lo.Map(files, func(item File, _ int) *File {
		return &item
	})
}

func ToFileDTO(file *File) *FileDTO {
	return &FileDTO{
		ID:        file.ID,
		Uid:       file.Uid,
		FileName:  file.FileName,
		CreatedAt: file.CreatedAt,
		Type:      file.Type,
		Url:       file.Url,
	}
}

func ToFilesDTO(files []*File) []*FileDTO {
	return lo.Map(files, func(item *File, _ int) *FileDTO {
		return ToFileDTO(item)
	})
}
