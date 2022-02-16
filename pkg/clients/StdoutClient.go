package clients

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"mmocker/pkg/common"
)

type StdoutClient struct {
	keys map[string]map[string]common.FunctionResult
}

func (s *StdoutClient) Init(param map[string]interface{}) {
	cronInstance := cron.New()
	cronInstance.Start()

	eId, err := cronInstance.AddFunc("@every 5s", func() {
		if s.keys != nil {
			for processorName, item := range s.keys {
				for funcName, value := range item {
					fmt.Printf("Proc: %s, FuncName: %s, Value: %.2f, Tags: %v\n", processorName, funcName, value.Value, value.Tags)
				}
			}
			fmt.Printf("\n")
		}
	})

	if err != nil {
		panic(err)
	}

	s.keys = map[string]map[string]common.FunctionResult{}
	println(eId)

}

func (s *StdoutClient) InitP(param map[string]interface{}) BaseClientInterface {
	s.Init(param)
	return s
}

func (s *StdoutClient) Push(processorName string, result map[string]common.FunctionResult) {
	s.keys[processorName] = result
}
