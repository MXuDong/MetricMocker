package web

import (
	"github.com/gin-gonic/gin"
	"mmocker/conf"
	"mmocker/internal/web/handler"
)

func Run() {
	r := gin.Default()

	r.GET("/processors", handler.ListProcessor)

	r.GET("/functions", handler.ListAllFunction)
	r.GET("/function", handler.ListAllFunction)

	r.GET("/function/:func", handler.GetFuncDescribe)

	// get conf.ApplicationConfig
	r.GET("/application-config", handler.GetConfig)

	r.Run(conf.ApplicationConfig.Port)
}
