package clients

import (
	"fmt"
	"mmocker/pkg"
	"net/http"
	"strings"
	"sync"
)

type PrometheusClient struct {
	params map[string]interface{}
	tags   map[string]map[string]string
	lock   *sync.RWMutex
	procs  []*pkg.Processor
}

const PrometheusOutputPort = "PROMETHEUS_OUTPUT_PORT"

func (p *PrometheusClient) Init(value map[string]interface{}) error {

	port := ""

	if outP, ok := value[PrometheusOutputPort]; !ok {
		return fmt.Errorf("Can find port of proemtheus! ")
	} else {
		if port, ok = outP.(string); !ok {
			return fmt.Errorf("Cant convert to string: [%v] ", outP)
		}
	}

	p.lock = &sync.RWMutex{}

	p.params = value
	p.tags = make(map[string]map[string]string, 0)
	s := http.NewServeMux()
	s.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		for _, processor := range p.procs {
			for _, valueItem := range processor.Get() {
				tagStr := ""
				for tagKey, tagValue := range valueItem.Tags {
					tagVStr := strings.ReplaceAll(tagValue, " ", "_")
					tagStr += fmt.Sprintf("%s=\"%s\",", tagKey, tagVStr)
				}
				tagStr = tagStr[:len(tagStr)-1]
				itemStr := fmt.Sprintf("%s{%s} %e", valueItem.Name, tagStr, valueItem.Value)
				_, _ = writer.Write([]byte(itemStr))
				_, _ = writer.Write([]byte("\n"))
			}
		}
	})
	go func() {
		if err := http.ListenAndServe(port, s); err != nil {
			fmt.Println(err)
		}
	}()

	return nil
}

func (p *PrometheusClient) GetParam() map[string]interface{} {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.params
}

// Output in prometheus is cover put
func (p *PrometheusClient) Output() {

}

func (pc *PrometheusClient) Register(processor *pkg.Processor) {
	if pc.procs == nil {
		pc.procs = make([]*pkg.Processor, 0)
	}
	if processor != nil {
		pc.procs = append(pc.procs, processor)
	}
}
