package handler

import (
	"github.com/gin-gonic/gin"
	"mmocker/pkg/proc"
	"net/http"
)

func ListProcessor(ctx *gin.Context) {
	processors := proc.Processors

	ctx.JSONP(http.StatusOK, processors)
}
