package clients

import (
	"fmt"
	"os"
)

const StdoutFile = "STDOUT_FILE"

type StdoutClient struct {
	param  map[string]interface{}
	output *os.File
}

func (sc *StdoutClient) GetParam() map[string]interface{} {
	return sc.param
}

func (sc *StdoutClient) Init(value map[string]interface{}) error {
	sc.param = value
	if outP, ok := value[StdoutFile]; !ok {
		return fmt.Errorf("Stdout file can't find ")
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

func (sc *StdoutClient) PutValue(name string, value float64, tags map[string]string) {
	_, err := fmt.Fprintf(sc.output, "%s: '%v': tags:'%v'\n", name, value, tags)
	if err != nil {
		fmt.Println(err)
	}
}
