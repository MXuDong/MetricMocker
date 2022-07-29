package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mmocker/conf"
	"mmocker/internal/web/handler"
	"mmocker/pkg/proc"
)

func Run() {
	r := gin.Default()

	r.Static("/statics", "./statics")

	r.GET("/monitor-info", handler.MonitorHandler)

	r.GET("/processors", handler.ListProcessor)
	r.GET("/processors-remove-all", func(context *gin.Context) {
		proc.CutNotExistProcessors()
	})

	r.GET("/functions", handler.ListAllFunction)
	r.GET("/function", handler.ListAllFunction)
	r.GET("/function/:func", handler.GetFuncDescribe)
	r.GET("/function/:func/value", handler.GetFunctionValue)

	// get conf.ApplicationConfig
	r.GET("/application-config", handler.GetConfig)

	//--------------------------------------------------
	// for cloud monitors, all handler support any method.
	r.Any("/cloud-monitor/resource/count", handler.GetCountForCloudMonitor)
	r.Any("/cloud-monitor/resources", handler.ListResourceForCloudMonitor)

	err := r.Run(conf.ApplicationConfig.Port)
	fmt.Printf("%v", err)
}
