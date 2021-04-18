package service

import (
	"time"

	"github.com/yunzhubanban/server/app/model"
	"github.com/gogf/gf/frame/g"
)

// Service 服务信息
var Service = serviceService{}

// serviceService 服务信息服务
type serviceService struct{}

// Status 服务状态
func (a *serviceService) Status() *model.StatusResp {
	return &model.StatusResp{
		Time: time.Now().Unix(),
		Name: g.Cfg().GetString("yunzhubanban.Name"),
	}
}

func (a *serviceService) Devices(req *model.DeviceBindReq) *model.DeviceResp {
	return nil
}
