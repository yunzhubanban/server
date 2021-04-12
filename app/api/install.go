package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/yunzhubanban/server/app/service"
)

var Install = installApi{}

type installApi struct {
}

// ServiceInstall 用于班级终端获取取得服务安装信息.
func (a *installApi) Status(r *ghttp.Request) {
	_ = r.Response.WriteJson(service.Install.GetServiceInstallInfo())
}
