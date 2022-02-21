package handler

import (
	"github.com/gin-gonic/gin"
	"mmocker/doc"
	"mmocker/internal/web/model"
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

// GetFunctionValue return the value of specify function with params.
func GetFunctionValue(ctx *gin.Context) {
	funcName := ctx.Param("func")
	params := &model.FunctionParams{}
	funcParam := funcs.FunctionParams{
		Type: funcs.TypeStr(funcName),
	}

	funcItem := funcs.Function(funcParam)

	// parse base params
	_ = ctx.BindQuery(params)

	// parse input params
	var err error
	params.Params, err = funcs.ConvertMapStringToMapInterface(ctx.QueryMap("Params"), funcs.GetParamFields(funcItem))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	funcParam.Params = params.Params
	// re-get function
	funcItem = funcs.Function(funcParam)

	if params.From > params.End || params.Step <= 0 {
		ctx.String(http.StatusBadRequest, "Error range, out of limit, please set From < End and set Step > 0")
		return
	}
	if (params.End-params.From)/params.Step > 1000 {
		ctx.String(http.StatusBadRequest, "The range of (end - from)/step is so large, please make sure the point count less than 1000(can equals).")
		return
	}

	var resValue []model.ValueMap

	for inputValue := params.From; inputValue <= params.End; inputValue += params.Step {
		outputValue, err := funcItem.Call(inputValue)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		resValue = append(resValue, model.ValueMap{
			Input:  inputValue,
			Output: outputValue,
		})
	}

	res := model.FunctionCallValue{
		Expression: funcItem.Expression(),
		Values:     resValue,
	}

	ctx.JSONP(http.StatusOK, res)
}
