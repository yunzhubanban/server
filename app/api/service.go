package api

import (
	"github.com/yunzhubanban/server/app/model"
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

// Devices 用于班级终端班级号绑定
func (a *serviceApi) Devices(r *ghttp.Request) {
	var data *model.DeviceBindReq

	if err := r.Parse(data); err != nil {
		// TODO 使用通用错误响应体
		r.Response.WriteStatusExit(500, "Internal Server Error")
	}

	_ = r.Response.WriteJson(service.Service.Devices(data))
}
