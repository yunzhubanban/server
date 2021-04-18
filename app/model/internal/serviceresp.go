package internal

// StatusResp 服务状态响应
type StatusResp struct {
	Time int64  `json:"time"` // 当前时间（秒级时间戳）
	Name string `json:"name"` // 学校名
}

type DeviceResp struct {
	HasSetting bool   `json:"has_setting"` //标识是否存在远程配置文件（布尔值）
	Token      string `json:"token"`       // JWT Token
}
