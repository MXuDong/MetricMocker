package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"mmocker/conf"
	"mmocker/pkg/clients"
	"mmocker/utils"
	"mmocker/utils/log"
	"os"
	"syscall"
	"time"
)

func init() {
	log.InitLogrus()
}

// main func bootstrap the metric-mocker
func main() {
	//metrics := InitConf()
	//
	//// init client
	//var cs []clients.Client
	//for _, clientItem := range metrics.Clients {
	//	if c, err := clients.GetClient(clientItem.Name, clientItem.Type, clientItem.Params); err != nil {
	//		panic(err)
	//	} else {
	//		cs = append(cs, c)
	//	}
	//}
	//for _, p := range metrics.Processors {
	//	_ = p.Load()
	//
	//	// register to client
	//	for _, client := range p.ClientNames {
	//		log.Logger.Infof("Register {%s} to {%s}", p.Name, client)
	//		clientItem, err := clients.GetClient(client, "", map[string]interface{}{})
	//		if err != nil {
	//			continue
	//		}
	//		clientItem.Register(p)
	//	}
	//}
	//dur := metrics.Application.Ticker
	//if dur < 1 {
	//	dur = 5
	//}
	//Start(cs, dur)

	config := InitConf()

	for _, item := range config.Clients {
		clients.Client(item.Name, item.Type, item.Params)
	}

	for _, item := range config.Processors {
		item.Load()
	}

	for true {
		time.Sleep(5 * time.Second)
	}

}

func InitConf() *conf.Configs {
	log.Logger.Trace("Init the config")
	confPath := ""
	flag.StringVar(&confPath, "c", "./conf/config.yaml", "-c config-path")
	flag.Parse()

	if v := os.Getenv("CONF_PATH"); v != "" {
		fmt.Println("Check the env: CONF_PATH = " + v)
		confPath = v
	}

	log.Logger.Infof("Init log from path: %s", confPath)

	c := &conf.Configs{}
	viper.SetConfigFile(confPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Logger.Error(err)
		syscall.Exit(1)
	}

	if err := viper.Unmarshal(c); err != nil {
		log.Logger.Error(err)
		syscall.Exit(1)
	}

	// check the node-id, if the value is emtpy, set local directly.
	nodeId := os.Getenv("NODE_ID")
	if len(nodeId) == 0 {
		nodeId = utils.Local_str
	}
	c.Application.NodeId = nodeId

	return c
}
