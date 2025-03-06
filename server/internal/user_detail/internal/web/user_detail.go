package web

import (
	"net/http"

	"github.com/codepzj/Stellux/server/internal/pkg/wrap"
	"github.com/codepzj/Stellux/server/internal/user_detail/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type IUserDetailHandler interface {
	UpdateUserDetail(ctx *gin.Context, userDetailRequest UserDetailRequest) (*wrap.Response[any], error)
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
		group.PUT("/update", wrap.WrapWithBody(u.UpdateOne))
	}
}

func (u *UserDetailHandler) UpdateOne(ctx *gin.Context, userDetailRequest UserDetailRequest) (*wrap.Response[any], error) {
	userDetail := UserDetailReqToDO(&userDetailRequest)
	if userDetail == nil {
		return wrap.Fail[any](http.StatusBadRequest, nil, "用户ID格式错误"), errors.New("用户ID格式错误")
	}
	err := u.userDetailService.UpdateUserDetail(ctx, userDetail)
	if err != nil {
		return wrap.Fail[any](http.StatusInternalServerError, nil, err.Error()), err
	}
	return wrap.Success[any](nil, "更新用户信息成功"), nil
}
