package boot

import (
	_ "github.com/yunzhubanban/server/packed"

	"encoding/json"
	"os"

	"github.com/yunzhubanban/server/app/config"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

// init 初始化参数设置
func init() {
	s := g.Server()
	s.Plugin(new(swagger.Swagger))

	loadConfig()
}

// loadConfig 加载配置信息
func loadConfig() {

	env, ok := os.LookupEnv("YZBBCONF")
	if !ok {
		g.Log().Fatal("请于环境变量中加入配置信息")
	}

	err := json.Unmarshal([]byte(env), &config.Config)
	if err != nil {
		g.Log().Fatal("解析配置信息失败：", err)
	}
}
