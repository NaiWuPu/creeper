package open_http_api

import (
	"creeper/runner"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	runner.RegisterOpenApiRunner("/openDemo", nil, OpenDemo)
	logrus.Info("openDemo 的内容注册到 RegisterOpenApiRunner")
}

func OpenDemo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OpenDemo pong",
	})
}
