package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"mmocker/conf"
	"mmocker/utils"
	"mmocker/utils/log"
	"os"
	"time"
)

func init() {
	log.InitLogrus()
}

// main func bootstrap the metric-mocker
func main() {

	confPath := ""
	flag.StringVar(&confPath, "c", "./conf/config.yaml", "-c config-path")
	flag.Parse()

	if v := os.Getenv("CONF_PATH"); v != "" {
		fmt.Println("Check the env: CONF_PATH = " + v)
		confPath = v
	}

	metrics := InitConf(confPath)
	log.Logger.Infof("Init log from path: %s", confPath)
	// init client

	for _, clientItem := range metrics.Clients {
		if _, err := utils.GetClient(clientItem.Name, clientItem.Type, clientItem.Params); err != nil {
			panic(err)
		}
	}
	for _, group := range metrics.Groups {
		group.PushTag()
		group.Load()
		group.Do()
	}

	log.Logger.Infof("Start the applications")
	for {
		// keep main handler
		time.Sleep(1000)
	}
}

func InitConf(configPath string) *conf.Configs {
	c := &conf.Configs{}
	viper.SetConfigFile(configPath)
	_ = viper.ReadInConfig()
	viper.Unmarshal(c)
	return c
}
