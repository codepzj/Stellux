package web

import (
	"net/http"

	"github.com/codepzj/Stellux/server/internal/pkg/wrap"
	"github.com/codepzj/Stellux/server/internal/posts/internal/service"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostHandler interface {
	CreatePosts(ctx *gin.Context, postsReq PostsReq) (*wrap.Response[any], error)
	FindPostById(ctx *gin.Context) (*wrap.Response[any], error)
	FindAllPosts(ctx *gin.Context) (*wrap.Response[any], error)
	FindPostsByCondition(ctx *gin.Context, page wrap.Page) (*wrap.Response[any], error)
	UpdatePublishStatus(ctx *gin.Context, updatePublishStatusReq UpdatePublishStatusReq) (*wrap.Response[any], error)
	DeletePostSoftById(ctx *gin.Context) (*wrap.Response[any], error)
	ResumePostSoftById(ctx *gin.Context) (*wrap.Response[any], error)
}
type PostsHandler struct {
	serv service.IPostsService
}

func NewPostHandler(serv service.IPostsService) *PostsHandler {
	return &PostsHandler{serv}
}

func (h *PostsHandler) RegisterGinRoutes(router *gin.Engine) {
	group := router.Group("/posts")
	{
		group.GET("/:id", wrap.Wrap(h.FindPostById))
		group.GET("/list/all", wrap.Wrap(h.FindAllPosts))
		group.GET("/list", wrap.WrapWithBody(h.FindPostsByCondition))
		group.POST("/create", wrap.WrapWithBody(h.CreatePosts))
		group.PUT("/update/status", wrap.WrapWithBody(h.UpdatePublishStatus))
		group.PUT("/resume/:id", wrap.Wrap(h.ResumePostSoftById))
		group.DELETE("/soft-delete/:id", wrap.Wrap(h.DeletePostSoftById))
	}
}

func (h *PostsHandler) CreatePosts(ctx *gin.Context, postsReq PostsReq) (*wrap.Response[any], error) {
	err := h.serv.CreatePosts(ctx, toPosts(postsReq))
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	return wrap.Success[any](nil, "新增文章成功"), nil
}

func (h *PostsHandler) FindPostById(ctx *gin.Context) (*wrap.Response[any], error) {
	id := ctx.Param("id")
	if id == "" {
		return wrap.Fail[any](http.StatusBadRequest, nil, "参数错误"), errors.New("参数错误")
	}
	idObj, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	posts, err := h.serv.FindPostById(ctx, idObj)
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	return wrap.Success[any](DTOToVO(posts), "获取文章详情成功"), nil
}

func (h *PostsHandler) FindAllPosts(ctx *gin.Context) (*wrap.Response[any], error) {
	posts, err := h.serv.FindAllPosts(ctx)
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	return wrap.Success[any](DTOsToVOs(posts), "获取文章列表成功"), nil
}

func (h *PostsHandler) FindPostsByCondition(ctx *gin.Context, page wrap.Page) (*wrap.Response[any], error) {
	posts, totalCount, totalPage, err := h.serv.FindPostsByCondition(ctx, &page)
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}

	pageVO := wrap.PageVO[PostsVO]{
		Page:       page,
		TotalCount: totalCount,
		TotalPage:  totalPage,
		List:       DTOsToVOs(posts),
	}
	return wrap.Success[any](pageVO, "获取文章列表成功"), nil
}

func (h *PostsHandler) UpdatePublishStatus(ctx *gin.Context, updatePublishStatusReq UpdatePublishStatusReq) (*wrap.Response[any], error) {
	id := updatePublishStatusReq.ID
	isPublish := updatePublishStatusReq.IsPublish
	idObj, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	err = h.serv.UpdatePostsPublishStatus(ctx, idObj, isPublish)
	if err != nil {
		return wrap.Fail[any](http.StatusInternalServerError, nil, err.Error()), err
	}
	return wrap.Success[any](nil, "更新文章状态成功"), nil
}

func (h *PostsHandler) DeletePostSoftById(ctx *gin.Context) (*wrap.Response[any], error) {
	id := ctx.Param("id")
	if id == "" {
		return wrap.Fail[any](http.StatusBadRequest, nil, "id值为空"), errors.New("id值为空")
	}
	idObj, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	err = h.serv.DeletePostSoftById(ctx, idObj)
	if err != nil {
		return wrap.Fail[any](http.StatusInternalServerError, nil, err.Error()), err
	}
	return wrap.Success[any](nil, "软删除文章成功"), nil
}

func (h *PostsHandler) ResumePostSoftById(ctx *gin.Context) (*wrap.Response[any], error) {
	id := ctx.Param("id")
	if id == "" {
		return wrap.Fail[any](http.StatusBadRequest, nil, "id值为空"), errors.New("id值为空")
	}
	idObj, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	err = h.serv.ResumePostSoftById(ctx, idObj)
	if err != nil {
		return wrap.Fail[any](http.StatusInternalServerError, nil, err.Error()), err
	}
	return wrap.Success[any](nil, "恢复文章成功"), nil
}
