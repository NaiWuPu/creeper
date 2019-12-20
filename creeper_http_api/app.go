package creeper_http_api

import (
	"creeper/runner"
	"creeper/service"
	"github.com/gin-gonic/gin"
)

//
// @Summary 创建应用
// @Description 创建应用
// @Accept  json
// @Produce  json
// @Param   app_name	body	string	true	"应用名称"
// @Success 200 {string} string	"ok"
// @Router /app/create [post]
func CreateApp(c *gin.Context) {
	appName := c.Param("app_name")
	//内容验证
	ot := service.CreateApp(appName)
	if ot.Error != nil {
		c.JSON(200, (&runner.CreeperOutput{}).ErrorOutput(ot.Error.Error()))
	}

	c.JSON(200, (&runner.CreeperOutput{}).SuccessOutput(ot.Data, ""))
}
