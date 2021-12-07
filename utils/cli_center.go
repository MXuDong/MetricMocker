package utils

import (
	"mmocker/pkg/clients"
)

var cache map[string]clients.Client = make(map[string]clients.Client, 0)

func GetClient(name string, param map[string]interface{}) (clients.Client, error) {
	if v := IsInCache(name); v != nil {
		return v, nil
	}

	var client clients.Client
	switch name {
	case "StdoutClient":
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
