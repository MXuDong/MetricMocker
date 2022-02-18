package handler

import (
	"github.com/gin-gonic/gin"
	"mmocker/doc"
	"mmocker/pkg/funcs"
	"net/http"
	"sort"
)

func ListAllFunction(ctx *gin.Context) {
	funcNames := []string{}
	for typeName, _ := range funcs.FuncMap {
		funcNames = append(funcNames, string(typeName))
	}

	sort.Strings(funcNames)
	ctx.JSONP(http.StatusOK, funcNames)
}

func GetFuncDescribe(ctx *gin.Context) {
	funcName := ctx.Param("func")

	funcParam := funcs.FunctionParams{
		Type: funcs.TypeStr(funcName),
	}

	funcItem := funcs.Function(funcParam)
	if funcItem == nil {
		ctx.String(http.StatusNotFound, "Not found function: '%s'. Please check.", funcName)
		return
	}

	htmlInfo := doc.GetHtml(funcItem)
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, htmlInfo)
}
