package clients

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"mmocker/pkg/common"
	"mmocker/utils"
	"time"
)

type InfluxdbV1Client struct {
	client influxdb2.Client
	database string
	table  string
}

const (
	InfluxDBServerURLKey = "ServerURL"
	InfluxDBTableKey = "Table"
	InfluxDBDatabaseKey = "Database"
	InfluxDBV1ClientType = "InfluxDBV1"
)

func (i *InfluxdbV1Client) Init(param map[interface{}]interface{}) {

	url := utils.GetStringWithDefault(param, InfluxDBServerURLKey, "")
	if len(url) == 0{
		panic(InfluxDBV1ClientType + " find empty of ServerURL")
	}
	database := utils.GetStringWithDefault(param, InfluxDBDatabaseKey, "metric-mocker/autogen")
	table := utils.GetStringWithDefault(param, InfluxDBTableKey, "test-metric")
	i.table = table
	i.database = database
	i.client = influxdb2.NewClient(url, "")
}

func (i *InfluxdbV1Client) InitP(param map[interface{}]interface{}) BaseClientInterface {
	i.Init(param)
	return i
}

func (i InfluxdbV1Client) Push(processorName string, result map[string]common.FunctionResult) {
	writeApi := i.client.WriteAPI("", i.database)
	for key, resultItem := range result {
		// set base key
		resultItem.Tags["proc"] = processorName
		resultItem.Tags["func"] = key
		p := influxdb2.NewPoint(i.table,
			resultItem.Tags,
			map[string]interface{}{
				"value": resultItem.Value,
			},
			time.Now(),
		)
		writeApi.WritePoint(p)
	}

}
