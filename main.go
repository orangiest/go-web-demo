package main

import (
	"cwm.wiki/web/config"
	"cwm.wiki/web/router"
	"flag"
	"fmt"
	"os"
)

// 希望实现的功能
// 1. 可自由配置的端口 doc 表示需要使用 环境变量
// 2. 单独文件的路由配置 已实现在 router 中
var (
	configFile = flag.String("c", "", "Configuration File name")
)

func main() {

	flag.Parse()

	// load configuration file./
	err := config.InitConfig(*configFile)
	if err != nil {
		fmt.Println("configuration load failed.")
		os.Exit(1)
	}

	// get configuration
	port := config.GetString("servers.port")

	if port == "" {
		port = "9998"
		fmt.Printf("Default port is %s", port)
	}

	router.RegiesterGin(port)

}
