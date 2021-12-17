package clients

import (
	"mmocker/utils/log"
)

var cache map[string]Client = make(map[string]Client, 0)

func GetClient(name, typ string, param map[string]interface{}) (Client, error) {
	log.Logger.Infof("Get client: {%s} in type: {%s} with params: {%v}", name, typ, param)
	if v := IsInCache(name); v != nil {
		log.Logger.Infof("Load client: {%s} in type: {%s} from cache.", name, typ)
		return v, nil
	}

	log.Logger.Warnf("Not cached fount special client: {%s} in type: {%s}, load.", name, typ)

	var client Client
	switch typ {
	case "StdoutClient":
		client = &StdoutClient{}
		if err := client.Init(param); err != nil {
			return nil, err
		}
	case "PrometheusClient":
		client = &PrometheusClient{}
		if err := client.Init(param); err != nil {
			return nil, err
		}
	default:
		client = &StdoutClient{}
		if err := client.Init(param); err != nil {
			return nil, err
		}
	}

	RegisterClient(name, client)
	return client, nil
}

func RegisterClient(name string, client Client) {
	cache[name] = client
}

func IsInCache(name string) Client {
	return cache[name]
}
