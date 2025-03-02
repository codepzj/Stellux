package api

import (
	"fmt"
	"net/http"
	"server/internal/file/internal/service"
	"server/internal/pkg/wrap"

	"github.com/gin-gonic/gin"
)

type IFileHandler interface {
	UploadFilesLocal(ctx *gin.Context)
	GetPhotosLocalByPage(ctx *gin.Context)
	DeletePhotoByUid(ctx *gin.Context)
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
		group.GET("/list", h.GetPhotosLocalByPage)
		group.POST("/local/upload", h.UploadFilesLocal)
		group.DELETE("/local/delete", h.DeletePhotoByUid)
	}
}

// UploadFilesLocal 上传图片到本地
func (h *FileHandler) UploadFilesLocal(ctx *gin.Context) {
	var uploadFilesRequest UploadFilesRequest
	if err := ctx.ShouldBind(&uploadFilesRequest); err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, err.Error())
		return
	}
	uids := uploadFilesRequest.Uids
	fileNames := uploadFilesRequest.FileNames
	files := uploadFilesRequest.Files
	if len(uids) == 0 || len(fileNames) == 0 || len(files) == 0 {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, "未找到上传的文件")
		return
	}
	if len(uids) != len(fileNames) || len(uids) != len(files) {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, "文件数量不匹配")
		return
	}
	pictures, err := h.serv.UploadFilesLocal(ctx, uids, fileNames, files)
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	count := len(pictures)
	wrap.SuccessWithMsg(ctx, fmt.Sprintf("上传成功，共%d张图片", count))
}

func (h *FileHandler) GetPhotosLocalByPage(ctx *gin.Context) {
	var page wrap.Page
	if err := ctx.ShouldBind(&page); err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, err.Error())
		return
	}
	pageNo := page.PageNo
	size := page.Size

	photos, totalCount, totalPage, err := h.serv.GetPhotosByPage(ctx, pageNo, size)

	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusInternalServerError, err.Error())
	}

	filePageDTO := wrap.ToPageVO(int64(pageNo), int64(size), totalCount, totalPage, photos)
	wrap.SuccessWithData(ctx, filePageDTO)
}

func (h *FileHandler) DeletePhotoByUid(ctx *gin.Context) {
	uid := ctx.DefaultQuery("uid", "")
	if uid == "" {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, "uid不能为空")
		return
	}
	err := h.serv.DeletePhotoByUid(ctx, uid)
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	wrap.SuccessWithMsg(ctx, "删除成功")
}
