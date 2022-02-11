package clients

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"mmocker/pkg/common"
)

type StdoutClient struct {
	keys map[string]map[string]common.FunctionResult
}

func (s *StdoutClient) Init(param map[interface{}]interface{}) {
	cronInstance := cron.New()
	cronInstance.Start()

	eId, err := cronInstance.AddFunc("@every 5s", func() {
		fmt.Printf("%v", s.keys)
	})

	if err != nil {
		panic(err)
	}

	s.keys = map[string]map[string]common.FunctionResult{}
	println(eId)

}

func (s *StdoutClient) InitP(param map[interface{}]interface{}) BaseClientInterface {
	s.Init(param)
	return s
}

func (s *StdoutClient) Push(processorName string, result map[string]common.FunctionResult) {
	s.keys[processorName] = result
}
