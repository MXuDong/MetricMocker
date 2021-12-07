package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"mmocker/conf"
	"mmocker/pkg/clients"
	"mmocker/utils"
	"os"
	"time"
)

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

	// init client

	for _, clientItem := range metrics.Clients {
		if _, err := utils.GetClient(clientItem.Name, clientItem.Params); err != nil {
			panic(err)
		}
	}
	for _, group := range metrics.Groups {
		group.PushTag()
		group.Load()
		group.Do()
	}

	for {
		time.Sleep(1000)
	}
}

func backMain() {
	confPath := ""
	flag.StringVar(&confPath, "c", "./conf/config.yaml", "-c config-path")
	flag.Parse()

	if v := os.Getenv("CONF_PATH"); v != "" {
		fmt.Println("Check the env: CONF_PATH = " + v)
		confPath = v
	}

	metrics := InitConf(confPath)
	for _, group := range metrics.Groups {
		group.PushTag()
		group.Load()
	}

	stdoutClient := clients.StdoutClient{}
	params := map[string]interface{}{
		clients.StdoutFile: os.Stdout,
	}

	if err := stdoutClient.Init(params); err != nil {
		panic(err)
	}

	for _, item := range metrics.Groups {
		w := item.Workers
		if w != nil {
			for _, itemW := range w {
				go itemW.DoFunc(1)()
			}
		}
	}

	for {
		time.Sleep(1000)
	}
}

func InitConf(configPath string) *conf.Configs {
	c := &conf.Configs{}
	viper.SetConfigFile("/Users/bytedance/workspaces/go-pro/MetricMocker/conf/config.yaml")
	_ = viper.ReadInConfig()
	viper.Unmarshal(c)
	return c
}
