package app

import (
	"fmt"
	"log"

	"github.com/codepzj/Stellux/server/global"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
}

func NewHttpServer(engine *gin.Engine) *HttpServer {
	return &HttpServer{
		engine: engine,
	}
}

func (s *HttpServer) Start() {
	port := global.Env.Port
	if port == "" {
		log.Fatal("未配置端口，请检查.env文件中的PORT设置")
	}
	log.Printf("Server is running on port %s", port)
	if err := s.engine.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
