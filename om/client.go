package om

import (
	"encoding/json"
	oc "github.com/application-mocker/object-client"
	"mmocker/conf"
	"mmocker/pkg/proc"
)

var (
	ObjectClient    *oc.ObjectClient
	ProcessorClient *oc.ObjectClient
)

func Init() {
	if conf.ApplicationConfig.ObjectMockerConfig.Enable {
		var err error
		ObjectClient, err = oc.NewObjectClient(conf.ApplicationConfig.ObjectMockerConfig.Host, "__metric_mocker__")
		if err != nil {
			panic(err)
		}

		ProcessorClient, err = ObjectClient.SubClient("processor")
		if err != nil {
			panic(err)
		}
	}
}

func DeleteProcessor(procName string, holder string) error {
	if len(holder) == 0 {
		holder = conf.ApplicationConfig.NodeId
	}
	objs, err := ProcessorClient.ListAllValue()
	if err != nil {
		return err
	}

	for _, item := range objs {
		jsonObj, err := json.Marshal(item.DataValue)
		if err != nil {
			return err
		}
		procItem := &proc.Processor{}
		if err := json.Unmarshal(jsonObj, procItem); err != nil {
			return err
		}
		if procItem.Name == procName && holder == procItem.Holder {
			if _, err := ProcessorClient.DeleteById(item.Id); err != nil {
				return err
			}
		}
	}
	return nil
}

// RegisterProcessor register a processor
func RegisterProcessor(proc proc.Processor) (string, error) {
	if err := DeleteProcessor(proc.Name, proc.Holder); err != nil {
		return "", err
	}
	id, err := ProcessorClient.InsertOne(proc)
	if err != nil {
		return "", err
	}
	return id, nil
}

func ListProcessor() ([]*proc.Processor, error) {
	res := make([]*proc.Processor, 0)
	if ProcessorClient == nil {
		return res, nil
	}

	objs, err := ProcessorClient.ListAllValue()
	if err != nil {
		return res, err
	}
	if objs == nil {
		return res, nil
	}

	res = make([]*proc.Processor, len(objs))
	for index, item := range objs {
		jsonObj, err := json.Marshal(item.DataValue)
		if err != nil {
			return make([]*proc.Processor, 0), err
		}
		resItem := &proc.Processor{}

		if json.Unmarshal(jsonObj, resItem) != nil {
			return make([]*proc.Processor, 0), err
		}
		res[index] = resItem
	}
	return res, nil
}
