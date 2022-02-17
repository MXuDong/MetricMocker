package handler

import (
	"github.com/gin-gonic/gin"
	"mmocker/pkg/funcs"
	"net/http"
)

func ListAllFunction(ctx *gin.Context) {
	funcNames := []funcs.TypeStr{}
	for typeName, _ := range funcs.FuncMap {
		funcNames = append(funcNames, typeName)
	}
	ctx.JSONP(http.StatusOK, funcNames)
}
