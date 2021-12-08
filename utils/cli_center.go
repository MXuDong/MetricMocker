package utils

import (
	"mmocker/pkg/clients"
	"mmocker/utils/log"
)

var cache map[string]clients.Client = make(map[string]clients.Client, 0)

func GetClient(name, typ string, param map[string]interface{}) (clients.Client, error) {
	log.Logger.Infof("Get client: {%s} in type: {%s} with params: {%v}", name, typ, param)
	if v := IsInCache(name); v != nil {
		log.Logger.Infof("Load client: {%s} in type: {%s} from cache.", name, typ)
		return v, nil
	}

	log.Logger.Infof("Not cached fount special client: {%s} in type: {%s}, load.", name, typ)

	var client clients.Client
	switch typ {
	case "StdoutClient":
		client = &clients.StdoutClient{}
		if err := client.Init(param); err != nil {
			return nil, err
		}
	case "PrometheusClient":
		client = &clients.PrometheusClient{}
		if err := client.Init(param); err != nil {
			return nil, err
		}
	default:
		client = &clients.StdoutClient{}
		if err := client.Init(param); err != nil {
			return nil, err
		}
	}

	RegisterClient(name, client)
	return client, nil
}

func RegisterClient(name string, client clients.Client) {
	cache[name] = client
}

func IsInCache(name string) clients.Client {
	return cache[name]
}
