package internal

import "time"

// Config 云助班班配置文件
type Config struct {
	Name       string        // 学校名称
	Password   string        // 管理员密码
	JWTKey     string        // JWT Key
	JWTTimeout time.Duration // JWT Token 超时时间, 单位为秒
}
