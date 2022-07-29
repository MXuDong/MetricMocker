package proc

import (
	"mmocker/pkg/funcs"
	"mmocker/utils/log"
	"sort"
	"strings"
	"sync"
	"time"
)

var Processors map[string]*Processor

// for cache
var lastUpdateTime int64
var listFlagTime int64
var processorsList []*Processor
var functionList []funcs.BaseFuncInterface

func ListProcessors() ([]*Processor, []funcs.BaseFuncInterface) {
	if lastUpdateTime == listFlagTime {
		return processorsList, functionList
	}

	processorsList = []*Processor{}

	lockProcessor()
	defer unLockProcessor()

	// double check
	if lastUpdateTime == listFlagTime {
		return processorsList, functionList
	}

	listFlagTime = lastUpdateTime

	for _, value := range Processors {
		funcCount := 0
		for _, funcItem := range value.Functions {
			functionList = append(functionList, funcItem)
			funcCount++
		}

		value.FunctionCount = funcCount
		value.ClientCount = len(value.Clients)

		processorsList = append(processorsList, value)
	}

	sort.Slice(processorsList, func(i, j int) bool {
		return strings.Compare(processorsList[i].Name, processorsList[j].Name) < 0
	})

	return processorsList, functionList
}

var lockItem sync.Mutex = sync.Mutex{}

func lockProcessor() {
	lockItem.Lock()
}

func unLockProcessor() {
	lockItem.Unlock()
}

// AddProcessors is thread-safe to add process. And AddProcessors do not cover old value.
func AddProcessors(processors ...*Processor) {
	lockProcessor()
	defer unLockProcessor()
	lastUpdateTime = time.Now().Unix()
	if Processors == nil {
		Processors = map[string]*Processor{}
	}
	if processors != nil {
		for _, item := range processors {
			if _, ok := Processors[item.Name]; ok {
				continue
			}
			log.Logger.Infof("Register processor to local center: %s", item.Name)
			Processors[item.Name] = item
			item.Load()
		}
	}
}

// CutNotExistProcessors will remove processor which can't find in processorNames.
// And if processorNames is nil, remove all processor.
func CutNotExistProcessors(processorNames ...string) int {
	lockProcessor()
	defer unLockProcessor()

	removeCount := 0

	if len(processorNames) == 0 {
		// remove all
		removeCount = len(Processors)
		for _, item := range Processors {
			item.Unload()
		}
		// remove point, for gc
		Processors = map[string]*Processor{}
		return removeCount
	}

	nameMap := map[string]struct{}{}
	for _, name := range processorNames {
		nameMap[name] = struct{}{}
	}
	for name, pItem := range Processors {
		if _, ok := nameMap[name]; !ok {
			pItem.Unload()
			delete(Processors, name)
		}
	}

	return removeCount
}
