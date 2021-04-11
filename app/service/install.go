package service

import (
	"github.com/yunzhubanban/server/app/config"
	"time"
)

var Install = installService{}

// installService 班级终端服务安装处理器
type installService struct{}

// InstallServiceResp 服务安装信息响应
type InstallServiceResp struct {
	time int64  // 当前时间（秒级时间戳）
	name string // 学校名
}

// GetServiceInstallInfo 获取服务安装信息
// GET /endpoint/service/status
func (a *installService) GetServiceInstallInfo() (resp *InstallServiceResp) {
	return &InstallServiceResp{
		time.Now().Unix(),
		config.Config.Name,
	}
}
