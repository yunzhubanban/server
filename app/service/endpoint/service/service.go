package service

import (
	"github.com/yunzhubanban/server/app/model/endpoint/service"
	"github.com/yunzhubanban/server/app/service/util/endpoint/auth"
	"time"

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
	// Create the token
	token, _, err := auth.Auth.TokenGenerator(nil)

	if err != nil {
		return nil
	}

	return &model.DeviceResp{
		// FIXME: 从数据库中获取是否已有设置
		HasSetting: false,
		Token:      token,
	}
}
