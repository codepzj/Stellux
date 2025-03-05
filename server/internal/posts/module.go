package posts

import (
	"server/internal/posts/internal/service"
	"server/internal/posts/internal/web"
)

type (
	Handler = web.PostsHandler
	Service = service.IPostsService
	Module  struct {
		Hdl *Handler
		Svc Service
	}
)
