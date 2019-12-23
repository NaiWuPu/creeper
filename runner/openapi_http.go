package runner

import (
	//加载api文档使用
	_ "creeper/creeper_http_docs"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

//路由
type openApiEngineRouter struct {
	routerPath  string
	method      []string
	handlerFunc gin.HandlerFunc
}

var openApiEngine *gin.Engine

//存储路由
var openApiEngineRouters []*openApiEngineRouter

func OpenApiRunner() {
	//启动设置端口
	cfg, err := goconfig.LoadConfigFile("etc/creeper.ini")
	if err != nil {
		panic(err)
	}
	mode, err := cfg.GetValue("open_api", "mode")
	if err != nil {
		panic(err)
	}
	gin.SetMode(mode)
	openApiEngine = gin.New()
	//允许使用跨域请求,全局中间件
	openApiEngine.Use(cors())
	httpPort, err := cfg.GetValue("open_api", "http_port")
	if err != nil {
		panic(err)
	}
	//路由加载
	loadOpenApiEngineRouter()
	if mode == "debug" {
		//swagger
		url := ginSwagger.URL(fmt.Sprintf("http://127.0.0.1:%s/swagger/doc.json", httpPort)) // The url pointing to API definition
		openApiEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
	//启动
	err = openApiEngine.Run(fmt.Sprintf(":%s", httpPort))
	if err != nil {
		panic(err)
	}
}

//给控制器注册路由使用
func RegisterOpenApiRunner(routerPath string, method []string, handlerFunc gin.HandlerFunc) {
	openApiEngineRouters = append(openApiEngineRouters, &openApiEngineRouter{
		routerPath:  routerPath,
		method:      method,
		handlerFunc: handlerFunc})
	logrus.Info("路由长度：", len(openApiEngineRouters))
}

//加载已经注册的路由
func loadOpenApiEngineRouter() {
	for _, router := range openApiEngineRouters {
		//method空就是所有
		if len(router.method) == 0 {
			openApiEngine.Any(router.routerPath, router.handlerFunc)
		} else {
			for _, m := range router.method {
				openApiEngine.Handle(m, router.routerPath, router.handlerFunc)
			}
		}
	}
}
