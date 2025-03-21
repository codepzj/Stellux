package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		log.Println("请求来源origin:", origin)
		if origin != "" {
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true") // 允许携带cookie
			c.Header("Access-Control-Max-Age", "86400")
			c.Set("content-type", "application/json")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
