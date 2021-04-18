package router

import (
	"github.com/yunzhubanban/server/app/api/endpoint/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/apis", func(group *ghttp.RouterGroup) {
		group.Group("/endpoint", func(group *ghttp.RouterGroup) {
			group.Group("/service", func(group *ghttp.RouterGroup) {
				group.
					GET("status", service.Status).
					PUT("devices", service.Devices)
			})
		})
	})
}
