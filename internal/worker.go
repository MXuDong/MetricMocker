package internal

import (
	"mmocker/pkg/clients"
	"mmocker/pkg/funcs"
	"mmocker/utils"
	"mmocker/utils/log"
	"strconv"
	"time"
)

type Group struct {
	Name string            `json:"Names" yaml:"Names"`
	Tags map[string]string `json:"Tags" yaml:"Tags"`

	Groups  []*Group  `json:"groups" yaml:"groups"`
	Workers []*Worker `json:"workers" yaml:"workers"`
}

func (g *Group) Do() {
	for _, groupItem := range g.Groups {
		groupItem.Do()
	}

	for _, workItem := range g.Workers {
		if workItem != nil {
			go workItem.DoFunc(1)()
		}
	}
}

func (g *Group) Load() {
	log.Logger.Infof("Load the group: %s ", g.Name)
	for _, groupItem := range g.Groups {
		groupItem.Load()
	}

	for _, workItem := range g.Workers {
		workItem.Load()
	}
}

// PushTag will push all tag to children group and workers
func (g *Group) PushTag() {
	log.Logger.Infof("Push tag from %s : %v", g.Name, g.Tags)
	for k, v := range g.Tags {
		g.PutTag(k, v)
	}
}

// PutTag will put a tag to current group and push it to children groups and workers
func (g *Group) PutTag(key, value string) {
	if g.Groups != nil {
		for _, item := range g.Groups {
			item.PutTag(key, value)
		}
	}
	if g.Workers != nil {
		for _, item := range g.Workers {
			item.PutTag(key, value)
		}
	}
	if g.Tags == nil {
		g.Tags = map[string]string{}
	}

	g.Tags[key] = value
}

// GetTag will get curretn group tags
func (g *Group) GetTag() map[string]string {
	return g.Tags
}

type Worker struct {
	Tags map[string]string `json:"Tags" yaml:"Tags"`
	Name string            `json:"Name" yaml:"Name"`

	FunctionName   string             `json:"functionName" yaml:"functionName"`
	FunctionParams map[string]float64 `json:"functionParams" yaml:"functionParams"`
	ClientsName    []string           `json:"clients" yaml:"clients"`
	Clients        []clients.Client
	f              funcs.Function
	startTime      int64
}

func (w *Worker) DoFunc(duration int64) func() {

	return func() {
		tags := w.Tags
		w.Reset()
		timeTickerChan := time.Tick(time.Duration(duration) * time.Second)
		for {
			for _, clientItem := range w.Clients {
				if clientItem != nil {
					(clientItem).PutValue(w.Name, w.Value(), tags)
				}
			}
			<-timeTickerChan
		}
	}
}

// Load will load function and clients
func (w *Worker) Load() {
	log.Logger.Infof("Load the worker: %s with func: %s and clients: %v", w.Name, w.FunctionName, w.ClientsName)
	w.f = utils.GetFunc(w.FunctionName, w.FunctionParams)
	for _, workerItem := range w.ClientsName {
		c, _ := utils.GetClient(workerItem, "", map[string]interface{}{})
		w.Clients = append(w.Clients, c)
	}
	if w.f != nil {
		for k, v := range w.f.Params() {
			w.PutTag(k, strconv.FormatFloat(v, 'f', -1, 64))
		}
	}

	w.PutTag("function_name", w.FunctionName)
}

// GetTag will return current worker tags
func (w *Worker) GetTag() map[string]string {
	return w.Tags
}

// PutTag will put a tag to current worker
func (w *Worker) PutTag(key, value string) {
	if w.Tags == nil {
		w.Tags = map[string]string{}
	}
	w.Tags[key] = value
}

// Reset the time of start
func (w *Worker) Reset() {
	w.startTime = time.Now().UnixNano()
}

// Value return value of now time(in second)
func (w *Worker) Value() float64 {
	// time can't less than zero
	now := time.Now().UnixNano()
	if w.startTime <= 0 {
		w.startTime = now
	}

	// turn nano to second :
	sec := float64(now-w.startTime) / float64(time.Second) * float64(time.Nanosecond)
	return w.f.Execute(sec)
}
