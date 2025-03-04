package api

import (
	"server/internal/pkg/utils"
	"server/internal/pkg/wrap"
	"server/internal/user/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	serv service.IUserService
}

func NewUserHandler(serv service.IUserService) *UserHandler {
	return &UserHandler{serv}
}

func (h *UserHandler) RegisterGinRoutes(router *gin.Engine) {
	userRouter := router.Group("/user")
	{
		userRouter.POST("/login", h.Login)
		userRouter.POST("/create", h.CreateUser)
		userRouter.POST("/list", h.FindAllUsers)
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var loginReq LoginReq
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, "参数错误")
		return
	}
	user := LoginReqToDO(&loginReq)
	if u, isExist := h.serv.FindUserIsExist(ctx, user); isExist {
		token, _ := utils.GenerateJwt(u.ID)
		loginVo := LoginVO{
			User:  DTOToVO(u),
			Token: token,
		}
		wrap.SuccessWithDetail(ctx, loginVo, "登录成功")
		return
	}
	wrap.FailWithMsg(ctx, http.StatusBadRequest, "用户名或者密码错误")
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var createUserReq *CreateUserReq
	if err := ctx.ShouldBindJSON(createUserReq); err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, "参数错误")
		return
	}
	user := CreateUserReqToDO(createUserReq)
	if err := h.serv.CreateUser(ctx, user); err != nil {
		wrap.FailWithMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	wrap.SuccessWithMsg(ctx, "创建用户成功")
}

func (h *UserHandler) FindAllUsers(ctx *gin.Context) {
	users, err := h.serv.FindAllUsers(ctx)
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	wrap.SuccessWithDetail(ctx, DTOsToVOs(users), "查询成功")
}
