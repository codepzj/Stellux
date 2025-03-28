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
	FindPostById(ctx *gin.Context) (*wrap.Response[any], error)
	FindAllPosts(ctx *gin.Context) (*wrap.Response[any], error)
	FindPostsByCondition(ctx *gin.Context, page wrap.Page) (*wrap.Response[any], error)
	AdminFindPostsByCondition(ctx *gin.Context, page wrap.Page) (*wrap.Response[any], error)
	AdminUpdatePostsStatus(ctx *gin.Context, updatePublishStatusReq UpdatePublishStatusReq) (*wrap.Response[any], error)
	AdminUpdatePost(ctx *gin.Context, updatePostReq UpdatePostReq) (*wrap.Response[any], error)
	AdminCreatePost(ctx *gin.Context, postsReq PostsReq) (*wrap.Response[any], error)
	AdminResumePostSoftById(ctx *gin.Context) (*wrap.Response[any], error)
	DeletePostSoftById(ctx *gin.Context) (*wrap.Response[any], error)
}
type PostsHandler struct {
	serv service.IPostsService
}

func NewPostHandler(serv service.IPostsService) *PostsHandler {
	return &PostsHandler{serv}
}

func (h *PostsHandler) RegisterGinRoutes(router *gin.Engine) {
	postsGroup := router.Group("/posts")
	{
		postsGroup.GET("/:id", wrap.Wrap(h.FindPostById))                  // 获取特定文章
		postsGroup.GET("/list/all", wrap.Wrap(h.FindAllPosts))             // 获取所有文章
		postsGroup.GET("/list", wrap.WrapWithBody(h.FindPostsByCondition)) // 根据条件获取文章（分页，排序，关键词，发布）
	}
	adminPostsGroup := router.Group("/admin-api/posts")
	{
		adminPostsGroup.POST("/list", wrap.WrapWithBody(h.AdminFindPostsByCondition)) // 管理员根据条件获取文章（分页，排序，关键词，不区分发布）
		adminPostsGroup.POST("/create", wrap.WrapWithBody(h.AdminCreatePost))         // 管理员创建文章
		adminPostsGroup.PUT("/update", wrap.WrapWithBody(h.AdminUpdatePost))         // 管理员更新文章
		adminPostsGroup.PATCH("/status", wrap.WrapWithBody(h.AdminUpdatePostsStatus)) // 管理员更新文章发布状态
		adminPostsGroup.PATCH("/resume/:id", wrap.WrapWithUri(h.AdminResumePostSoftById))    // 管理员恢复删除文章
		adminPostsGroup.PATCH("/delete/:id", wrap.WrapWithUri(h.AdminDeletePostSoftById))    // 管理员软删除文章
	}
}

// FindPostById 获取特定文章
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

// FindAllPosts 获取所有文章
func (h *PostsHandler) FindAllPosts(ctx *gin.Context) (*wrap.Response[any], error) {
	posts, err := h.serv.FindAllPosts(ctx)
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	return wrap.Success[any](DTOsToVOs(posts), "获取文章列表成功"), nil
}

// FindPostsByCondition 根据条件获取文章（分页，排序，关键词）
func (h *PostsHandler) FindPostsByCondition(ctx *gin.Context, page wrap.Page) (*wrap.Response[any], error) {
	posts, totalCount, totalPage, err := h.serv.FindPostByCondition(ctx, &page)
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

// AdminFindPostsByCondition 管理员根据条件获取文章（分页，排序，关键词）
func (h *PostsHandler) AdminFindPostsByCondition(ctx *gin.Context, page wrap.Page) (*wrap.Response[any], error) {
	posts, totalCount, totalPage, err := h.serv.AdminFindPostByCondition(ctx, &page)
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

// AdminCreatePost 管理员创建文章
func (h *PostsHandler) AdminCreatePost(ctx *gin.Context, postsReq PostsReq) (*wrap.Response[any], error) {
	err := h.serv.AdminCreatePost(ctx, toPosts(postsReq))
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	return wrap.Success[any](nil, "新增文章成功"), nil
}

// AdminUpdatePost 管理员更新文章
func (h *PostsHandler) AdminUpdatePost(ctx *gin.Context, updatePostReq UpdatePostReq) (*wrap.Response[any], error) {
	err := h.serv.AdminUpdatePost(ctx, updatePostReqToPosts(updatePostReq))
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	return wrap.Success[any](nil, "更新文章成功"), nil
}

// AdminUpdatePostsStatus 管理员更新文章发布状态
func (h *PostsHandler) AdminUpdatePostsStatus(ctx *gin.Context, updatePublishStatusReq UpdatePublishStatusReq) (*wrap.Response[any], error) {
	idObj, err := bson.ObjectIDFromHex(updatePublishStatusReq.ID)
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	err = h.serv.AdminUpdatePostStatus(ctx, idObj, updatePublishStatusReq.IsPublish)
	if err != nil {
		return wrap.Fail[any](http.StatusInternalServerError, nil, err.Error()), err
	}
	return wrap.Success[any](nil, "更新文章状态成功"), nil
}

// AdminDeletePostSoftById 管理员软删除文章
func (h *PostsHandler) AdminDeletePostSoftById(ctx *gin.Context, postIDReq PostIDReq) (*wrap.Response[any], error) {
	idObj, err := bson.ObjectIDFromHex(postIDReq.ID)
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	err = h.serv.AdminDeletePostSoftById(ctx, idObj)
	if err != nil {
		return wrap.Fail[any](http.StatusInternalServerError, nil, err.Error()), err
	}
	return wrap.Success[any](nil, "软删除文章成功"), nil
}

// AdminResumePostSoftById 管理员恢复删除文章
func (h *PostsHandler) AdminResumePostSoftById(ctx *gin.Context, postIDReq PostIDReq) (*wrap.Response[any], error) {
	idObj, err := bson.ObjectIDFromHex(postIDReq.ID)
	if err != nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, err.Error()), err
	}
	err = h.serv.AdminResumePostSoftById(ctx, idObj)
	if err != nil {
		return wrap.Fail[any](http.StatusInternalServerError, nil, err.Error()), err
	}
	return wrap.Success[any](nil, "恢复文章成功"), nil
}
