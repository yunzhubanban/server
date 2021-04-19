package service

import (
	"github.com/yunzhubanban/server/app/api/endpoint/service"
	"time"

	"github.com/gogf/gf/frame/g"
)

// Service 服务信息
var Service = serviceService{}

// serviceService 服务信息服务
type serviceService struct{}

// StatusResp 服务状态响应
type StatusResp struct {
	Time int64  `json:"time"` // 当前时间（秒级时间戳）
	Name string `json:"name"` // 学校名
}

// Status 服务状态
func (a *serviceService) Status() *StatusResp {
	return &StatusResp{
		Time: time.Now().Unix(),
		Name: g.Cfg().GetString("yunzhubanban.Name"),
	}
}

type DeviceResp struct {
	HasSetting bool   `json:"has_setting"` //标识是否存在远程配置文件（布尔值）
	Token      string `json:"token"`       // JWT Token
}

func (a *serviceService) Devices(req *service.DeviceBindReq) *DeviceResp {
	return nil
}
