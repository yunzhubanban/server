package api

import (
	"github.com/yunzhubanban/server/app/service"

	"github.com/gogf/gf/net/ghttp"
)

// Service 服务信息
var Service = serviceApi{}

// serviceApi 服务信息接口
type serviceApi struct{}

// Status 用于班级终端获取取得服务安装信息
func (a *serviceApi) Status(r *ghttp.Request) {
	_ = r.Response.WriteJson(service.Service.Status())
}
