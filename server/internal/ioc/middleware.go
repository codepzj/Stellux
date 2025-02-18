package ioc

import (
	"server/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitMiddleWare() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Recovery(),
		middleware.GinLogger(),
		middleware.Cors(),
		middleware.JWT(),
	}
}
