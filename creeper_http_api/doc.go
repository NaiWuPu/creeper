package creeper_http_api

//主要为了文档，也可以作为注册使用
import (
	"creeper/runner"
	"github.com/sirupsen/logrus"
)

func init() {
	runner.RegisterCreeperApiRunner("/app/create", []string{"POST"}, CreateApp)
	logrus.Info("CreateApp 的内容注册到 CreeperApiRunner")
}

// @title 系统管理api
// @version 1.0
// @description creeper内部管理api.
// @termsOfService http://creeper.io/

// @contact.name lvxin
// @contact.url http://creeper.io/
// @contact.email lvxin@creeper.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8081
// @BasePath /
