package open_http_api

import (
	"creeper/app"
	"creeper/runner"
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	runner.RegisterOpenApiRunner("/postDemo", []string{"POST"}, PostDemo)
	runner.RegisterOpenApiRunner("/getDemo", []string{"GET"}, GetDemo)
	runner.RegisterOpenApiRunner("/putDemo", []string{"PUT"}, PutDemo)
	logrus.Info("demo 的内容注册到 RegisterOpenApiRunner")
}

func PostDemo(c *gin.Context) {
	cfg, err := goconfig.LoadConfigFile("etc/http_proxy.ini")
	if err != nil {
		panic(err)
	}
	proxyInfo, err := cfg.GetSection("demo_proxy")
	if err != nil {
		panic(err)
	}
	app.CreeperProxy(
		&app.CreeperProxyHost{
			Ctx:    c,
			Host:   proxyInfo["host"],
			Path:   "/Demo3",
			Scheme: proxyInfo["scheme"]})
}

func GetDemo(c *gin.Context) {
	cfg, err := goconfig.LoadConfigFile("etc/http_proxy.ini")
	if err != nil {
		panic(err)
	}
	proxyInfo, err := cfg.GetSection("demo_proxy")
	if err != nil {
		panic(err)
	}
	app.CreeperProxy(
		&app.CreeperProxyHost{
			Ctx:    c,
			Host:   proxyInfo["host"],
			Path:   "/Demo3",
			Scheme: proxyInfo["scheme"]})
}

func PutDemo(c *gin.Context) {
	cfg, err := goconfig.LoadConfigFile("etc/http_proxy.ini")
	if err != nil {
		panic(err)
	}
	proxyInfo, err := cfg.GetSection("demo_proxy")
	if err != nil {
		panic(err)
	}
	app.CreeperProxy(
		&app.CreeperProxyHost{
			Ctx:    c,
			Host:   proxyInfo["host"],
			Path:   "/Demo3",
			Scheme: proxyInfo["scheme"]})
}
