package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	. "github.com/codepzj/Stellux/server/internal/pkg/logger"
)

// GinLogger gin日志中间件
func GinLogger() gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now() // 记录请求开始时间

		// 获取请求的路径和查询参数等
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		date := time.Now().Format("2006-01-02 15:04:05")

		// 处理请求
		c.Next()

		// 计算请求耗时
		duration := time.Since(start)

		// 记录日志
		Logger.Info(
			"请求日志",
			zap.Int("status", status),
			zap.String("method", method),
			zap.Duration("duration", duration),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("clientIP", clientIP),
			zap.String("userAgent", userAgent),
			zap.String("date", date),
		)
	}
}
