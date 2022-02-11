package proc

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"mmocker/pkg/clients"
	"mmocker/pkg/common"
	"mmocker/pkg/funcs"
	"time"
)

type FunctionName string

type Processor struct {
	Name string `yaml:"Name"`

	Tags map[string]string `yaml:"Tags"`

	FunctionParamsList []funcs.FunctionParams `yaml:"FunctionParamsList"`

	Functions    map[FunctionName]funcs.BaseFuncInterface `yaml:"Functions"`
	FunctionTags map[FunctionName]map[string]string       `yaml:"FunctionTags"`

	// IgnoreFunctionTag is a boolean value, if true, ignore the function param, but keep name and type tag.
	IgnoreFunctionTag bool `yaml:"IgnoreFunctionTag"`

	Clients []string `yaml:"Clients"`

	ClientInstances []clients.BaseClientInterface
}

func (p *Processor) Load() {
	// load function
	p.FunctionTags = map[FunctionName]map[string]string{}
	p.Functions = map[FunctionName]funcs.BaseFuncInterface{}
	if p.Tags == nil {
		p.Tags = map[string]string{}
	}

	for _, funcParamItem := range p.FunctionParamsList {
		name := funcParamItem.Name
		f := funcs.Function(funcParamItem)
		if f == nil {
			// log here
			continue
		}
		p.Functions[FunctionName(name)] = f
		p.FunctionTags[FunctionName(name)] = map[string]string{}
		for processorTagName, processorTagValue := range p.Tags {
			p.FunctionTags[FunctionName(name)][processorTagName] = processorTagValue
		}
		p.FunctionTags[FunctionName(name)]["function_name"] = name
		p.FunctionTags[FunctionName(name)]["function_type"] = string(funcParamItem.Type)

		if !p.IgnoreFunctionTag {
			params := p.Functions[FunctionName(name)].Params()
			for key, value := range params {
				keyStr := fmt.Sprintf("%v", key)
				valueStr := fmt.Sprintf("%v", value)

				p.FunctionTags[FunctionName(name)][keyStr] = valueStr
			}
		}
	}
	// load client
	for _, clientName := range p.Clients {
		if client := clients.Client(clientName, "", nil); client != nil {
			p.ClientInstances = append(p.ClientInstances, client)
		}
	}

	cItem := cron.New()
	cItem.Start()
	cItem.AddFunc("@every 5s", func() {
		res := p.Values()
		for _, clientItem := range p.ClientInstances {
			clientItem.Push(p.Name, res)
		}
	})
}

// Values will call all function. And return with processor and function param tag.
func (p Processor) Values() map[string]common.FunctionResult {
	nowTime := time.Now().Unix()
	result := make(map[string]common.FunctionResult, len(p.Functions))
	for funcName, funcItem := range p.Functions {
		fr := common.FunctionResult{}

		res, err := funcItem.Call(float32(nowTime))
		if err != nil {
			// log err
		} else {
			fr.Value = res
			fr.Tags = p.FunctionTags[funcName]

			result[string(funcName)] = fr
		}
	}

	return result
}