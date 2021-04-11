package boot

import (
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"github.com/yunzhubanban/server/app/config"
	_ "github.com/yunzhubanban/server/packed"
	"os"
)

// init 用于服务初始化
func init() {
	loadConfig()
}

// loadConfig 加载配置文件
func loadConfig() {
	env, ok := os.LookupEnv("YZBBCONF")

	if !ok {
		g.Log().Error("在解析环境配置时发生了异常: 无配置信息")
		return
	}

	err := json.Unmarshal([]byte(env), &config.Config)

	if err != nil {
		g.Log().Error("在解析环境配置时发生了异常", err)
	}
}
