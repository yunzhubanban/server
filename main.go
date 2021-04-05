package main

import (
	_ "github.com/yunzhubanban/server/boot"
	_ "github.com/yunzhubanban/server/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
