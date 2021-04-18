package boot

import (
	_ "github.com/yunzhubanban/server/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

// init 初始化参数设置
func init() {
	s := g.Server()
	s.Plugin(new(swagger.Swagger))
}
