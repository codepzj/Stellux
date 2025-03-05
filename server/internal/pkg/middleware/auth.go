package middleware

import (
	"log"
	"net/http"
	"server/global"
	"server/internal/pkg/wrap"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.GetString("userId")
		requestURI := ctx.Request.RequestURI
		method := ctx.Request.Method
		log.Println("requestURI为:", requestURI, "method为:", method)
		ok, err := global.Enforcer.Enforce(userId, requestURI, method)
		if err != nil {
			wrap.FailWithMsg(ctx, http.StatusInternalServerError, "权限校验失败")
			ctx.Abort()
			return
		}
		if ok {
			log.Println("userId为:", userId, "允许访问")
			ctx.Next()
		} else {
			log.Println("userId为:", userId, "禁止访问")
			ctx.AbortWithStatus(http.StatusForbidden)
		}
	}
}
