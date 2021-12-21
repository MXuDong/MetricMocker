package clients

import (
	"fmt"
	"mmocker/pkg"
	"mmocker/utils/log"
	"os"
)

const StdoutFile = "STDOUT_FILE"

type StdoutClient struct {
	param   map[string]interface{}
	output  *os.File
	procs   []*pkg.Processor
	enabled bool
}

func (sc *StdoutClient) GetParam() map[string]interface{} {
	return sc.param
}

func (sc *StdoutClient) Init(value map[string]interface{}) error {
	sc.param = value
	sc.enabled = true
	if outP, ok := value[StdoutFile]; !ok {
		log.Logger.Warnf("Output file not find, checkout to disbale.")
		sc.enabled = false
		return nil
	} else {
		if output, ok := outP.(*os.File); ok {
			sc.output = output
			return nil
		} else if outputStr, ok := outP.(string); ok && outputStr == StdoutFile {
			sc.output = os.Stdout
			return nil
		} else {
			return fmt.Errorf("[%s] can't convert to [*os.File] or [string(STDOUT_FILE)]", outP)
		}
	}
}

func (sc *StdoutClient) Output() {
	if !sc.enabled {
		return
	}
	if sc.procs == nil {
		return
	}

	for _, procItem := range sc.procs {
		for _, valueItem := range procItem.Get() {
			fmt.Printf("FuncName: %s, tag: %v, value: %.2f\n", valueItem.Name, valueItem.Tags, valueItem.Value)
		}
	}
}

func (sc *StdoutClient) Register(processor *pkg.Processor) {
	if sc.procs == nil {
		sc.procs = make([]*pkg.Processor, 0)
	}
	if processor != nil {
		sc.procs = append(sc.procs, processor)
	}
}
