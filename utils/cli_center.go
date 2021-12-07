package utils

import (
	"mmocker/pkg/clients"
)

var cache map[string]clients.Client = make(map[string]clients.Client, 0)

func GetClient(name, typ string, param map[string]interface{}) (clients.Client, error) {
	if v := IsInCache(name); v != nil {
		return v, nil
	}

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
