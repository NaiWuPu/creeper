package creeper_http_api

import (
	"creeper/runner"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	runner.RegisterCreeperApiRunner("/postDemo", []string{"POST"}, PostDemo)
	runner.RegisterCreeperApiRunner("/getDemo", []string{"GET"}, GetDemo)
	runner.RegisterCreeperApiRunner("/putDemo", []string{"PUT"}, PutDemo)
	runner.RegisterCreeperApiRunner("/Demo3", []string{"POST", "GET", "PUT"}, Demo3)
	logrus.Info("demo 的内容注册到 CreeperApiRunner")
}

func PostDemo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PostDemo pong",
	})
}

func GetDemo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetDemo pong",
	})
}

func PutDemo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PutDemo pong",
	})
}

func Demo3(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Demo3 pong",
	})
}
