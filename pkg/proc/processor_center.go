package proc

import "mmocker/utils/log"

var Processors map[string]*Processor

// AddProcessors is thread-safe to add process. And AddProcessors do no cover old value.
func AddProcessors(processors ...*Processor) {
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
