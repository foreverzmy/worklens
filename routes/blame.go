package routes

import (
	"bufio"
	"time"

	"github.com/foreverzmy/worklens/utils"
	"github.com/valyala/fasthttp"
)

func Blame(ctx *fasthttp.RequestCtx) {
	r, err := utils.GetRepo(ctx)
	if err != nil {
		ctx.Error(err.Error(), 500)
		return
	}

	ctx.SetBodyStreamWriter(func(w *bufio.Writer) {
		head, err := r.Head()
		if err != nil {
			ctx.Error(err.Error(), 500)
		}

		headInfo := map[string]interface{}{
			"name": head.Name().Short(),
			"hash": head.Hash().String(),
			"type": head.Type().String(),
		}

		utils.WriteStreamMessage(w, "head", headInfo)

		time.Sleep(1 * time.Second)

		commit, _, err := r.GetBranchFiles(head.Name().Short())
		if err != nil {
			ctx.Error(err.Error(), 500)
		}

		utils.WriteStreamMessage(w, "commit", commit)
	})
}
