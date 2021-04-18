package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/yunzhubanban/server/app/model"
	"github.com/yunzhubanban/server/app/service"
)

// Device 班级终端注册
var Device = deviceApi{}

// serviceApi 服务信息接口
type deviceApi struct{}

// Devices 用于班级终端班级号绑定
func (a *deviceApi) Devices(r *ghttp.Request) {
	var data *model.DeviceBindReq

	if err := r.Parse(data); err != nil {
		// TODO 使用通用错误响应体
		r.Response.WriteStatusExit(500, "Internal Server Error")
	}

	_ = r.Response.WriteJson(service.Service.Devices(data))
}
