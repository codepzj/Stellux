package web

import (
	"net/http"
	"server/internal/pkg/wrap"
	"server/internal/user_detail/internal/service"

	"github.com/gin-gonic/gin"
)

type IUserDetailHandler interface {
	UpdateUserDetail(ctx *gin.Context) error
}

type UserDetailHandler struct {
	userDetailService service.IUserDetailService
}

func NewUserDetailHandler(userDetailService service.IUserDetailService) *UserDetailHandler {
	return &UserDetailHandler{userDetailService: userDetailService}
}

func (u *UserDetailHandler) RegisterGinRoutes(router *gin.Engine) {
	group := router.Group("/user_detail")
	{
		group.PUT("/update", u.UpdateOne)
	}
}

func (u *UserDetailHandler) UpdateOne(ctx *gin.Context) {
	var userDetailRequest UserDetailRequest
	if err := ctx.ShouldBindJSON(&userDetailRequest); err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, "参数错误")
		return
	}
	userDetail := UserDetailReqToDO(&userDetailRequest)
	if userDetail == nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, "用户ID格式错误")
		return
	}
	err := u.userDetailService.UpdateUserDetail(ctx, userDetail)
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	wrap.SuccessWithMsg(ctx, "更新用户信息成功")
}
