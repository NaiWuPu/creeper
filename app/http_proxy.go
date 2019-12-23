package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

type CreeperProxyHost struct {
	Scheme string
	Host   string
	Path   string
	Ctx    *gin.Context
}

func CreeperProxy(cph *CreeperProxyHost) {
	proxy := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = cph.Scheme
			req.URL.Host = cph.Host
			req.URL.Path = cph.Path
			req.Host = cph.Host
		},
	}
	proxy.ServeHTTP(cph.Ctx.Writer, cph.Ctx.Request)
}
