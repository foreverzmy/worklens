package main

import (
	"os"

	"github.com/foreverzmy/worklens/routes"
	"github.com/foreverzmy/worklens/server"
	"github.com/valyala/fasthttp"
)

//// go:embed web/dist/web/*
// var webFiles embed.FS

func main() {
	webFS := os.DirFS("./web/dist")

	s := server.Server{
		Port: 5187,
		// WebFiles: &webFiles,
		WebFiles:   webFS,
		Middleware: routes.RepoMiddleware,
		Routes: map[string]fasthttp.RequestHandler{
			"/blame":     routes.Blame,
			"/repo/info": routes.RepoInfo,
			"/stream":    routes.ResponseStreamHandler,
		},
	}

	s.Start()
}
