package model

// StatusResp 服务状态响应
type StatusResp struct {
	Time int64  `json:"time"` // 当前时间（秒级时间戳）
	Name string `json:"name"` // 学校名
}
