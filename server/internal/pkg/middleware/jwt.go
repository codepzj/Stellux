package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/codepzj/Stellux/server/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 排除GET请求
		if ctx.Request.Method == "GET" {
			ctx.Next()
			return
		}
		// 排除接口
		excludeUri := []string{
			"/user/login",
		}
		for _, uri := range excludeUri {
			if ctx.Request.RequestURI == uri {
				ctx.Next()
				return
			}
		}
		token := ctx.Request.Header.Get("Authorization")
		log.Println("用户携带的token:", token)
		// 若非GET请求的token为空
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			log.Println("用户未携带token")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, err := utils.ParseJwt(strings.TrimPrefix(token, "Bearer "))
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		log.Println("解析后的claims:", claims)
		ctx.Set("userId", claims.ID)
		ctx.Next()
	}
}
