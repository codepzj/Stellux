package middleware

import (
	"log"
	"net/http"
	"server/global"

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
