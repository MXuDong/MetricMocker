package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mmocker/conf"
	"mmocker/internal/web/handler"
)

func Run() {
	r := gin.Default()

	r.Static("/statics", "./statics")

	r.GET("/processors", handler.ListProcessor)

	r.GET("/functions", handler.ListAllFunction)
	r.GET("/function", handler.ListAllFunction)
	r.GET("/function/:func", handler.GetFuncDescribe)
	r.GET("/function/:func/value", handler.GetFunctionValue)

	// get conf.ApplicationConfig
	r.GET("/application-config", handler.GetConfig)

	err := r.Run(conf.ApplicationConfig.Port)
	fmt.Printf("%v", err)
}
