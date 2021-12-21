package pkg

import (
	"fmt"
	"mmocker/pkg/funcs"
	"mmocker/utils"
	"mmocker/utils/log"
	"time"
)

type ValueItem struct {
	Name  string
	Tags  map[string]string
	Value float64
}

type FuncPackage struct {
	f             funcs.Function
	FuncName      string // the function name for find func
	FuncNameAlias string // the function name in processor
	Params        map[string]float64

	tags map[string]string
}

func (fp *FuncPackage) Load() {
	fp.f = funcs.GetFunc(fp.FuncName, fp.Params)
}

type Processor struct {
	Name string
	Tags map[string]string

	Functions []*FuncPackage

	ClientNames []string

	start     bool
	startTime int64
}

func (p *Processor) Load() error {
	log.Logger.Trace("The processor {%s} is loading.", p.Name)
	if p.Functions != nil && len(p.Functions) != 0 {
		for _, funcItem := range p.Functions {
			funcItem.Load()
			// append tags
			funcItem.tags = map[string]string{}
			for k, v := range p.Tags {
				funcItem.tags[k] = v
			}
			funcItem.tags[utils.TagFuncStr] = funcItem.FuncNameAlias
			funcItem.tags[utils.TagFuncNameStr] = funcItem.FuncName
			if funcItem.f != nil {
				for k, v := range funcItem.f.Params() {
					funcItem.tags[k] = fmt.Sprintf("%.2f", v)
				}
			}
		}
	}
	return nil
}

func (p *Processor) Reset() {
	p.start = false
}

func (p *Processor) Get() []ValueItem {
	nowTime := time.Now().UnixNano()
	if !p.start {
		p.start = true
		p.startTime = nowTime
	}

	r := make([]ValueItem, 0)
	for _, fp := range p.Functions {
		value := fp.f.Execute(float64(nowTime-p.startTime) / float64(time.Second) * float64(time.Nanosecond))
		vitem := ValueItem{
			Name:  p.Name,
			Tags:  fp.tags,
			Value: value,
		}
		r = append(r, vitem)
	}
	return r
}
