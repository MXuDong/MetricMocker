package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mmocker/pkg/proc"
	"net/http"
)

//--------------------------------------------------
// custom define

// CommonResult define the result by TOP rule.
type CommonResult struct {
	ResponseMetadata struct {
		RequestId string `json:"RequestId"`
		Action    string `json:"Action"`
		Version   string `json:"Version"`
		Service   string `json:"Service"`
		Region    string `json:"Region"`
		Error     struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error"`
	} `json:"ResponseMetadata"`
	Result struct {
		Data  interface{} `json:"Data"`  // Data in this program is an array.
		Count int         `json:"Count"` // Count is len of the data.
	} `json:"Result"`
}

// CloudMonitorRequest is the param of cloud-monitor request.
type CloudMonitorRequest struct {
	Namespace    string `json:"Namespace"`    // the product namespace info.
	SubNamespace string `json:"SubNamespace"` // the sub-namespace in cloud-monitor define.
	TargetName   string `json:"TargetName"`   // specify the target value name.
	//Filters      Filters `json:"Filters"`      // the query filters.
	PageNumber int `json:"PageNumber"` // like  Offset.
	PageSize   int `json:"PageSize"`   // like Limit.
	Offset     int `json:"Offset"`     // like PageNumber.
	Limit      int `json:"Limit"`      // like PageSize.
}

type Filters struct {
	FilterItems []FilterItem `json:"-"`
}

type FilterItem struct {
	Name   string   `json:"Name"`
	Values []string `json:"Values"`
}

func GetCountForCloudMonitor(ctx *gin.Context) {
	processors, _ := proc.ListProcessors()

	res := CommonResult{}
	res.Result.Count = len(processors)

	ctx.JSON(http.StatusOK, res)
}

func ListResourceForCloudMonitor(ctx *gin.Context) {

	// define res object
	res := CommonResult{}

	// parse the cloud-monitor request param.
	reqParam := CloudMonitorRequest{
		PageSize: 10, // page size default is 10
		Limit:    10, // limit default is 10
	}

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		res.ResponseMetadata.Error.Message = err.Error()
		res.ResponseMetadata.Error.Code = fmt.Sprintf("%d", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
	}

	fmt.Println(string(body))

	if len(body) == 0 {
		body = []byte("{}")
	}

	if err := json.Unmarshal(body, &reqParam); err != nil {
		res.ResponseMetadata.Error.Message = err.Error()
		res.ResponseMetadata.Error.Code = fmt.Sprintf("%d", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)

	}

	// get all resources
	processors, functions := proc.ListProcessors()

	start := reqParam.PageNumber * reqParam.Offset
	end := (reqParam.PageNumber + 1) * reqParam.Limit

	if reqParam.SubNamespace == "Processor" || len(reqParam.SubNamespace) == 0 {
		//FIXME: maybe start and end should greater than zero? :-)
		if start >= len(processors) {
			// if start index is greater than processors' length, return empty array directly.
			res.Result.Data = []struct{}{}
			ctx.JSON(http.StatusOK, res)
			return
		} else if end >= len(processors) {
			end = len(processors)
		}

		// set result
		res.Result.Data = processors[start:end]
		res.Result.Count = len(processors)
	} else {
		if start >= len(functions) {
			// if start index is greater than processors' length, return empty array directly.
			res.Result.Data = []struct{}{}
			ctx.JSON(http.StatusOK, res)
			return
		} else if end >= len(functions) {
			end = len(functions)
		}
		res.Result.Data = functions[start:end]
		res.Result.Count = len(functions)
	}

	ctx.JSON(http.StatusOK, res)
}

func GetDimensionForCloudMonitor(ctx *gin.Context) {

}
