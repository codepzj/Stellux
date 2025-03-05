package web

import "mime/multipart"

type UploadFilesRequest struct {
	Uids      []string                `form:"uids" binding:"required"`
	FileNames []string                `form:"file_names" binding:"required"`
	Files     []*multipart.FileHeader `form:"files" binding:"required"`
}
