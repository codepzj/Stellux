package web

import (
	"server/internal/pkg/utils"
	"server/internal/pkg/wrap"
	"server/internal/user/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
		userRouter.POST("/login", wrap.WrapWithBody(h.Login))
		userRouter.POST("/create", wrap.WrapWithBody(h.CreateUser))
		userRouter.POST("/list", wrap.Wrap(h.FindAllUsers))
	}
}

func (h *UserHandler) Login(ctx *gin.Context, loginReq LoginReq) (*wrap.Response[any], error) {
	user := LoginReqToDO(&loginReq)
	if u, isExist := h.serv.FindUserIsExist(ctx, user); isExist {
		token, _ := utils.GenerateJwt(u.ID)
		loginVo := LoginVO{
			User:  DTOToVO(u),
			Token: token,
		}
		return wrap.Success[any](loginVo, "登录成功"), nil
	}
	return wrap.Fail[any](http.StatusBadRequest, nil, "用户名或者密码错误"), errors.New("用户名或者密码错误")
}

func (h *UserHandler) CreateUser(ctx *gin.Context, createUserReq CreateUserReq) (*wrap.Response[any], error) {
	user := CreateUserReqToDO(&createUserReq)
	_, err := h.serv.CreateUser(ctx, user)
	if err != nil {
		return wrap.Fail[any](http.StatusInternalServerError, nil, err.Error()), err
	}
	return wrap.Success[any](nil, "创建用户成功"), nil
}

func (h *UserHandler) FindAllUsers(ctx *gin.Context) (*wrap.Response[any], error) {
	users, err := h.serv.FindAllUsers(ctx)
	if err != nil {
		return wrap.Fail[any](http.StatusInternalServerError, nil, err.Error()), err
	}
	return wrap.Success[any](DTOsToVOs(users), "查询成功"), nil
}
