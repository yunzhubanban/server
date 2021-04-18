package api

import (
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/yunzhubanban/server/app/config"
	"time"
)

var (
	// Auth JWT 中间件
	Auth *jwt.GfJWTMiddleware
)

// 初始化 JWT
func init() {
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:         "YZBB",
		Key:           []byte(config.Config.JWTKey),
		Timeout:       time.Minute * config.Config.JWTTimeout,
		MaxRefresh:    time.Minute * config.Config.JWTTimeout,
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
