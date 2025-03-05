package posts

import (
	"server/internal/posts/internal/repo"
	"server/internal/posts/internal/service"
	"server/internal/posts/internal/web"
)

type (
	Handler    = web.PostsHandler
	Service    = service.IPostsService
	Repository = repo.IPostsRepo
	Module     struct {
		Hdl  *Handler
		Svc  Service
		Repo Repository
	}
)
