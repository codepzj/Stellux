package api

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
		group.GET("/list", h.FindAllPosts)
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
	wrap.SuccessWithDetail(ctx, toPostsVO(posts), "获取文章详情成功")
}

func (h *PostsHandler) FindAllPosts(ctx *gin.Context) {
	posts, err := h.serv.FindAllPosts(ctx)
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, err.Error())
		return
	}
	postsVO := make([]*PostsVO, 0)
	for _, post := range posts {
		postsVO = append(postsVO, toPostsVO(post))
	}
	wrap.SuccessWithDetail(ctx, postsVO, "获取文章列表成功")
}
