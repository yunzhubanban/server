package router

import (
	"github.com/gogf/gf-jwt/example/service"
	"github.com/yunzhubanban/server/app/api"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// authHook is the HOOK function implements JWT logistics.
func middlewareAuth(r *ghttp.Request) {
	api.Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}

func init() {
	s := g.Server()
	s.Group("/apis", func(group *ghttp.RouterGroup) {
		group.Group("/endpoint/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.CORS, middlewareAuth)
			group.GET("/service", api.Service)
		})
	})
}
