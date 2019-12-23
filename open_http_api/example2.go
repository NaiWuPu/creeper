package open_http_api

import (
	"creeper/app"
	"creeper/runner"
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
	app.CreeperProxy(
		&app.CreeperProxyHost{
			Ctx:    c,
			Host:   "127.0.0.1:8081",
			Path:   "/Demo3",
			Scheme: "http"})
}

func GetDemo(c *gin.Context) {
	app.CreeperProxy(
		&app.CreeperProxyHost{
			Ctx:    c,
			Host:   "127.0.0.1:8081",
			Path:   "/Demo3",
			Scheme: "http"})
}

func PutDemo(c *gin.Context) {
	app.CreeperProxy(
		&app.CreeperProxyHost{
			Ctx:    c,
			Host:   "127.0.0.1:8081",
			Path:   "/Demo3",
			Scheme: "http"})
}
