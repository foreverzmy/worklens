package routes

import (
	"bufio"
	"log"

	"github.com/foreverzmy/worklens/convert"
	"github.com/foreverzmy/worklens/utils"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/valyala/fasthttp"
)

func RepoInfo(ctx *fasthttp.RequestCtx) {
	r, err := utils.GetRepo(ctx)
	if err != nil {
		ctx.Error(err.Error(), 500)
		return
	}

	ctx.SetBodyStreamWriter(func(w *bufio.Writer) {
		remotes, err := r.Remotes()
		if err != nil {
			log.Fatalf("无法获取远程信息: %v", err)
		}

		utils.WriteStreamMessage(w, "remote", convert.ConvertToRemotesConfigJSON(remotes))

		// 获取HEAD引用
		head, err := r.Head()
		if err != nil {
			log.Fatalf("无法获取HEAD引用: %v", err)
		}

		utils.WriteStreamMessage(w, "head", convert.ConvertToReferenceJSON(head))

		commit, err := r.CommitObject(head.Hash())
		if err != nil {
			log.Fatalf("无法获取提交: %v", err)
		}

		utils.WriteStreamMessage(w, "commit", commit)

		branches, err := r.Branches()
		if err != nil {
			log.Fatalf("无法获取分支列表: %v", err)
		}

		err = branches.ForEach(func(ref *plumbing.Reference) error {
			utils.WriteStreamMessage(w, "branch", convert.ConvertToReferenceJSON(ref))
			return nil
		})
		if err != nil {
			log.Fatalf("遍历分支时出错: %v", err)
		}
	})
}
