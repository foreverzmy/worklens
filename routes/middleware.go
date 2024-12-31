package routes

import (
	"github.com/foreverzmy/worklens/repo"
	"github.com/foreverzmy/worklens/utils"
	"github.com/valyala/fasthttp"
)

func RepoMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Method()) == fasthttp.MethodPost {
			repoRoot, err := utils.GetRepoRoot(ctx)
			if err != nil {
				ctx.Error(err.Error(), 500)
				return
			}

			r, err := repo.PlainOpen(repoRoot)
			if err != nil {
				ctx.Error(err.Error(), 500)
				return
			}

			ctx.SetUserValue("repo", r)
		}

		next(ctx)
	}
}
