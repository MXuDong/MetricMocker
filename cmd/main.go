package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"mmocker/conf"
	"mmocker/pkg/clients"
	"mmocker/utils/log"
	"os"
	"time"
)

func init() {
	log.InitLogrus()
}

// main func bootstrap the metric-mocker
func main() {
	metrics := InitConf()

	// init client
	var cs []clients.Client
	for _, clientItem := range metrics.Clients {
		if c, err := clients.GetClient(clientItem.Name, clientItem.Type, clientItem.Params); err != nil {
			panic(err)
		} else {
			cs = append(cs, c)
		}
	}
	for _, p := range metrics.Processors {
		_ = p.Load()

		// register to client
		for _, client := range p.ClientNames {
			log.Logger.Infof("Register {%s} to {%s}", p.Name, client)
			clientItem, err := clients.GetClient(client, "", map[string]interface{}{})
			if err != nil {
				continue
			}
			clientItem.Register(p)
		}
	}
	dur := metrics.Application.Ticker
	if dur < 1 {
		dur = 5
	}
	Start(cs, dur)
}

func InitConf() *conf.Configs {
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
	_ = viper.ReadInConfig()
	viper.Unmarshal(c)
	return c
}

// Start will hold on the application, and use debug-out client
func Start(cs []clients.Client, ticker int) {
	log.Logger.Infof("Start the applications")
	timeTickerChan := time.Tick(time.Duration(ticker) * time.Second)
	for {
		// keep main handler
		for _, ci := range cs {
			ci.Output()
		}
		<-timeTickerChan
	}
}
