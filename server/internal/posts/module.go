package posts

import (
	"github.com/codepzj/Stellux/server/internal/posts/internal/repo"
	"github.com/codepzj/Stellux/server/internal/posts/internal/service"
	"github.com/codepzj/Stellux/server/internal/posts/internal/web"
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
