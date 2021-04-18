package service

import (
	"github.com/yunzhubanban/server/app/model"
	"github.com/yunzhubanban/server/app/service"

	"github.com/gogf/gf/net/ghttp"
)

// Devices 用于班级终端班级号绑定
func Devices(r *ghttp.Request) {
	var data *model.DeviceBindReq

	if err := r.Parse(data); err != nil {
		// TODO 使用通用错误响应体
		r.Response.WriteStatusExit(500, "Internal Server Error")
	}

	_ = r.Response.WriteJson(service.Service.Devices(data))
}
