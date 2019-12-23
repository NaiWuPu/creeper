package main

import (
	//注册内部控制器
	_ "creeper/creeper_http_api"
	_ "creeper/open_http_api"
	"creeper/runner"
)

func main() {
	//启动http
	go runner.CreeperApiRunner()

	//启动open api
	go runner.OpenApiRunner()

	select {}
}
