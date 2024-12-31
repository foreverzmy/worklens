package utils

import (
	"encoding/json"
	"errors"

	"github.com/foreverzmy/worklens/repo"
	"github.com/valyala/fasthttp"
)

func GetBody(ctx *fasthttp.RequestCtx) (map[string]interface{}, error) {
	body := ctx.PostBody()
	var params map[string]interface{}
	if err := json.Unmarshal(body, &params); err != nil {
		return nil, err
	}

	return params, nil
}

func GetRepoRoot(ctx *fasthttp.RequestCtx) (repoRoot string, err error) {
	params, err := GetBody(ctx)
	if err != nil {
		return
	}

	repoRoot, ok := params["repo"].(string)
	if !ok {
		return repoRoot, errors.New("missing or invalid 'repo' parameter")
	}
	return
}

func GetRepo(ctx *fasthttp.RequestCtx) (r *repo.Repo, err error) {
	if r, ok := ctx.UserValue("repo").(*repo.Repo); ok {
		return r, nil
	}

	return nil, errors.New("repo not found")
}
