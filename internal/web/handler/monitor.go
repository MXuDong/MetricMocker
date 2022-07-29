package handler

import (
	"github.com/gin-gonic/gin"
	"mmocker/internal/web/model"
	"runtime"
)

func MonitorHandler(ctx *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	res := model.SystemInfo{}
	res.Sys = m.Sys

	ctx.JSONP(200, res)
}
