package server

import (
	"fmt"
	"io/fs"
	"log"
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
)

type Server struct {
	Port       int
	WebFiles   fs.FS
	Middleware func(next fasthttp.RequestHandler) fasthttp.RequestHandler
	Routes     map[string]fasthttp.RequestHandler
}

func (s *Server) Start() {
	fs := &fasthttp.FS{
		Root:               "/",
		IndexNames:         []string{"index.html"},
		GenerateIndexPages: false,
		Compress:           true,
		PathRewrite: func(ctx *fasthttp.RequestCtx) []byte {
			path := string(ctx.Path())
			if path == "/" || path == "/index.html" {
				return []byte("/web/index.html")
			}
			return []byte(path)
		},
		FS: s.WebFiles,
	}

	staticHandler := fs.NewRequestHandler()

	// 创建请求路由器
	router := s.Middleware(func(ctx *fasthttp.RequestCtx) {
		pathname := string(ctx.Path())
		for route, handler := range s.Routes {
			if pathname == route {
				handler(ctx)
				return
			}
		}

		if pathname == "/" || pathname == "/index.html" {
			staticHandler(ctx)
			return

		}

		if strings.HasPrefix(pathname, "/web") {
			staticHandler(ctx)
			return
		}

		ctx.NotFound()
	})

	listener, err := reuseport.Listen("tcp4", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		log.Fatalf("error in reuseport listener: %s", err)
	}

	if err := fasthttp.Serve(listener, router); err != nil {
		log.Fatalf("error in fasthttp server: %s", err)
	}
}
