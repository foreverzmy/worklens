package main

import (
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"time"

	"github.com/foreverzmy/worklens/repo"
	"github.com/foreverzmy/worklens/server"
	"github.com/valyala/fasthttp"
)

//go:embed web/*
var webFiles embed.FS

const repoRoot = "/Users/mervyn/Workspace/chataizoo/frontend"

func main() {
	s := server.Server{
		Port:     5187,
		WebFiles: &webFiles,
		Routes: map[string]fasthttp.RequestHandler{
			"/blame": func(ctx *fasthttp.RequestCtx) {
				r, err := repo.PlainOpen(repoRoot)
				if err != nil {
					ctx.Error(err.Error(), 500)
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

					headResp, _ := json.Marshal(headInfo)

					fmt.Fprintf(w, "%s", headResp)

					if err := w.Flush(); err != nil {
						return
					}

					time.Sleep(1 * time.Second)

					commit, _, err := r.GetBranchFiles(head.Name().Short())
					if err != nil {
						ctx.Error(err.Error(), 500)
					}

					commitResp, _ := json.Marshal(commit)

					fmt.Fprintf(w, "%s", commitResp)
				})
			},
			"/stream": responseStreamHandler,
		},
	}

	s.Start()
}

func responseStreamHandler(ctx *fasthttp.RequestCtx) {
	// Send the response in chunks and wait for a second between each chunk.
	ctx.SetBodyStreamWriter(func(w *bufio.Writer) {
		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "this is a message number %d", i)

			// Do not forget flushing streamed data to the client.
			if err := w.Flush(); err != nil {
				return
			}
			time.Sleep(time.Second)
		}
	})
}
