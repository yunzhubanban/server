package api

import (
	"time"

	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
)

var (
	// Auth JWT 中间件
	Auth *jwt.GfJWTMiddleware
)

// 初始化 JWT
func init() {
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:         "YZBB",
		Key:           g.Cfg().GetBytes("yunzhubanban.jwt.Key"),
		Timeout:       time.Minute * g.Cfg().GetDuration("yunzhubanban.jwt.Timeout"), // TODO 可能不需要配置文件加载
		MaxRefresh:    time.Minute * g.Cfg().GetDuration("yunzhubanban.jwt.Timeout"),
		IdentityKey:   "id",
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		g.Log().Error("JWT 错误" + err.Error())
	}
	Auth = authMiddleware
}
