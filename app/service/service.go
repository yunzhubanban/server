package service

import (
	"time"

	"github.com/yunzhubanban/server/app/config"
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
		Name: config.Config.Name,
	}
}
