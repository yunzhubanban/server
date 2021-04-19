package service

import (
	"github.com/yunzhubanban/server/app/service"

	"github.com/gogf/gf/net/ghttp"
)

type DeviceBindReq struct {
	Point    string `json:"point"`    // Point 班级号
	Password string `json:"password"` // Password 密码散列值 (SHA512)
}

// Devices 用于班级终端班级号绑定
func Devices(r *ghttp.Request) {
	var data *DeviceBindReq

	if err := r.Parse(data); err != nil {
		// TODO 使用通用错误响应体
		r.Response.WriteStatusExit(500, "Internal Server Error")
	}

	_ = r.Response.WriteJson(service.Service.Devices(data))
}
