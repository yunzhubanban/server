package router

import (
	"github.com/yunzhubanban/server/app/api"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/apis", func(group *ghttp.RouterGroup) {
		group.Group("/endpoint/", func(group *ghttp.RouterGroup) {
			group.GET("/service", api.Service)
			group.PUT("/service", api.Device)
		})
	})
}
