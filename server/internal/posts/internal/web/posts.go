package web

import (
	"net/http"
	"server/internal/pkg/wrap"
	"server/internal/posts/internal/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostHandler interface {
	CreatePosts(ctx *gin.Context)
	FindPostById(ctx *gin.Context)
	FindAllPosts(ctx *gin.Context)
	FindPostsByCondition(ctx *gin.Context)
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
		group.GET("/:id", h.FindPostById)
		group.GET("/list/all", h.FindAllPosts)
		group.GET("/list", h.FindPostsByCondition)
		group.POST("/create", h.CreatePosts)
	}
}

func (h *PostsHandler) CreatePosts(ctx *gin.Context) {
	var postsReq PostsReq
	if err := ctx.ShouldBindJSON(&postsReq); err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, err.Error()+" 参数错误")
		return
	}
	err := h.serv.CreatePosts(ctx, toPosts(postsReq))
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, err.Error())
		return
	}
	wrap.SuccessWithMsg(ctx, "新增文章成功")
}

func (h *PostsHandler) FindPostById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, "参数错误")
		return
	}
	idObj, err := bson.ObjectIDFromHex(id)
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, err.Error())
		return
	}
	posts, err := h.serv.FindPostById(ctx, idObj)
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, err.Error())
		return
	}
	wrap.SuccessWithDetail(ctx, DTOToVO(posts), "获取文章详情成功")
}

func (h *PostsHandler) FindAllPosts(ctx *gin.Context) {
	posts, err := h.serv.FindAllPosts(ctx)
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, err.Error())
		return
	}
	wrap.SuccessWithDetail(ctx, DTOsToVOs(posts), "获取文章列表成功")
}

func (h *PostsHandler) FindPostsByCondition(ctx *gin.Context) {
	var page wrap.Page
	if err := ctx.ShouldBind(&page); err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, err.Error()+"，参数错误")
		return
	}
	posts, totalCount, totalPage, err := h.serv.FindPostsByCondition(ctx, &page)
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, err.Error())
		return
	}

	pageVO := wrap.PageVO[PostsVO]{
		Page:       page,
		TotalCount: totalCount,
		TotalPage:  totalPage,
		List:       DTOsToVOs(posts),
	}
	wrap.SuccessWithDetail(ctx, pageVO, "获取文章列表成功")
}
