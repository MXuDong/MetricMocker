package clients

import (
	"fmt"
	"net/http"
)

type PrometheusClient struct {
	params map[string]interface{}
	values map[string]float64
	tags   map[string]map[string]string
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

	p.params = value
	p.values = make(map[string]float64, 0)
	p.tags = make(map[string]map[string]string, 0)
	s := http.NewServeMux()
	s.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		for key, value := range p.values {
			tags := p.tags[key]
			tagStr := ""
			if tags != nil {
				for tagKey, tagValue := range tags {
					tagStr += fmt.Sprintf("%s=\"%s\",", tagKey, tagValue)
				}
				tagStr = tagStr[:len(tagStr)-1]
			}
			itemStr := fmt.Sprintf("%s{%s} %f", key, tagStr, value)
			_, _ = writer.Write([]byte(itemStr))
			_, _ = writer.Write([]byte("\n"))
			writer.WriteHeader(http.StatusOK)
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
	return p.params
}

// PutValue in prometheus is cover put
func (p *PrometheusClient) PutValue(name string, value float64, tags map[string]string) {
	p.values[name] = value
	p.tags[name] = tags
}
