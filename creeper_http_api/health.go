package creeper_http_api

import (
	"creeper/runner"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	runner.RegisterCreeperApiRunner("/", nil, Ping)
	logrus.Info("health 的 Ping 注册到 CreeperApiRunner")
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
