package handler

import (
	"github.com/gin-gonic/gin"
	"mmocker/conf"
	"net/http"
)

func GetConfig(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, conf.ApplicationConfig)
}
