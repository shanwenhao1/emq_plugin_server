package main

import (
	"emq_plugin_server/emqWeb"
	"emq_plugin_server/emqWeb/Infra/config"
)

func main() {
	// 加载server配置
	config.InitConfig("./config.ini")

	ser := emqWeb.Server{}
	// 初始化启动
	ser.InitializedSystem("config/server_gin.xml")
	// 通过回调函数注册路由, 并启动服务
	ser.Run()
}
