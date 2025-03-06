package web

import (
	"fmt"
	"net/http"
	"server/internal/file/internal/service"
	"server/internal/pkg/wrap"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type IFileHandler interface {
	UploadFilesLocal(ctx *gin.Context, uploadFilesRequest UploadFilesRequest) (*wrap.Response[any], error)
	GetPhotosLocalByPage(ctx *gin.Context, page wrap.Page) (*wrap.Response[any], error)
	DeletePhotoByUid(ctx *gin.Context) (*wrap.Response[any], error)
}

type FileHandler struct {
	serv service.IFileService
}

var _ IFileHandler = (*FileHandler)(nil)

func NewFileHandler(serv service.IFileService) *FileHandler {
	return &FileHandler{serv}
}

func (h *FileHandler) RegisterGinRoutes(router *gin.Engine) {
	router.Static("/images", "./static/images")
	group := router.Group("/picture")
	{
		group.GET("/list", wrap.WrapWithBody(h.GetPhotosLocalByPage))
		group.POST("/local/upload", wrap.WrapWithBody(h.UploadFilesLocal))
		group.DELETE("/local/delete", wrap.Wrap(h.DeletePhotoByUid))
	}
}

// UploadFilesLocal 上传图片到本地
func (h *FileHandler) UploadFilesLocal(ctx *gin.Context, uploadFilesRequest UploadFilesRequest) (*wrap.Response[any], error) {
	uids := uploadFilesRequest.Uids
	fileNames := uploadFilesRequest.FileNames
	files := uploadFilesRequest.Files
	if len(uids) == 0 || len(fileNames) == 0 || len(files) == 0 {
		return wrap.Fail[any](http.StatusInternalServerError, nil, "未找到上传的文件"), errors.New("未找到上传的文件")
	}
	if len(uids) != len(fileNames) || len(uids) != len(files) {
		return wrap.Fail[any](http.StatusInternalServerError, nil, "文件数量不匹配"), errors.New("文件数量不匹配")
	}
	pictures, err := h.serv.UploadFilesLocal(ctx, uids, fileNames, files)
	if err != nil {
		return wrap.Fail[any](http.StatusInternalServerError, nil, err.Error()), err
	}
	return wrap.Fail[any](http.StatusInternalServerError, nil, fmt.Sprintf("上传成功，共%d张图片", len(pictures))), nil
}

func (h *FileHandler) GetPhotosLocalByPage(ctx *gin.Context, page wrap.Page) (*wrap.Response[any], error) {
	pageNo := page.PageNo
	size := page.Size

	photos, totalCount, totalPage, err := h.serv.GetPhotosByPage(ctx, pageNo, size)

	if err != nil {
		return wrap.Fail[any](http.StatusInternalServerError, nil, err.Error()), err
	}

	filePageDTO := wrap.ToPageVO(pageNo, size, totalCount, totalPage, photos)
	return wrap.Success[any](filePageDTO, "分页获取图片成功"), nil
}

func (h *FileHandler) DeletePhotoByUid(ctx *gin.Context) (*wrap.Response[any], error) {
	uid := ctx.DefaultQuery("uid", "")
	if uid == "" {
		return wrap.Fail[any](http.StatusBadRequest, nil, "uid不能为空"), errors.New("uid不能为空")
	}
	err := h.serv.DeletePhotoByUid(ctx, uid)
	if err != nil {
		return wrap.Fail[any](http.StatusInternalServerError, nil, err.Error()), err
	}
	return wrap.Success[any](nil, "删除成功"), nil
}
