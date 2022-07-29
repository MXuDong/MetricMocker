package main

import (
	"flag"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"mmocker/conf"
	"mmocker/instances"
	"mmocker/internal/web"
	"mmocker/om"
	"mmocker/pkg/clients"
	"mmocker/pkg/proc"
	"mmocker/utils"
	"mmocker/utils/log"
	"os"
	"syscall"
)

func init() {
	log.InitLogrus()
}

// main func bootstrap the metric-mocker
func main() {
	config := InitConf()

	if conf.ApplicationConfig.ObjectMockerConfig.Enable {
		om.Init() // init it

		// init cron job
		cronInstance := cron.New()
		_, err := instances.GlobalCron.AddFunc(conf.ApplicationConfig.ObjectMockerConfig.SyncInterval, func() {
			functionSync()
		})
		if err != nil {
			log.Logger.Error("%v", err)
		}
		cronInstance.Start()
	}
	for _, item := range config.Clients {
		clients.Client(item.Name, item.Type, item.Params)
	}

	for _, item := range config.Processors {
		if len(item.Holder) == 0 {
			item.Holder = conf.ApplicationConfig.NodeId
		}
		if conf.ApplicationConfig.ObjectMockerConfig.Enable {
			if _, err := om.RegisterProcessor(*item); err != nil {
				log.Logger.Error("%v", err)
			}
		}
		proc.AddProcessors(item)
	}

	web.Run()
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
	if c.Application.ObjectMockerConfig == nil {
		c.Application.ObjectMockerConfig = &conf.ObjectMockerConfig{
			Enable: false,
		}
	}
	conf.ApplicationConfig = c.Application

	return c
}

func functionSync() {
	log.Logger.Infof("Sync from object-mocker by interval: %s", conf.ApplicationConfig.ObjectMockerConfig.SyncInterval)
	processors, err := om.ListProcessor()
	if err != nil {
		log.Logger.Errorf("%v", err)
	}
	var procNames []string
	proc.AddProcessors(processors...)
	for _, procItem := range processors {
		procNames = append(procNames, procItem.Name)
	}
	proc.CutNotExistProcessors(procNames...)
}
