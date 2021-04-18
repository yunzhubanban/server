package service

import (
	"github.com/yunzhubanban/server/app/service"

	"github.com/gogf/gf/net/ghttp"
)

// Status 用于班级终端获取取得服务安装信息
func Status(r *ghttp.Request) {
	_ = r.Response.WriteJson(service.Service.Status())
}
