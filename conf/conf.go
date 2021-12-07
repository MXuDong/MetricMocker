package conf

import "mmocker/internal"

type Configs struct {
	Groups  []internal.Group `yaml:"groups" json:"groups"`
	Clients []Client         `yaml:"clients" json:"clients"`
}

type Client struct {
	Name   string                 `yaml:"name" json:"name"`
	Params map[string]interface{} `yaml:"params" json:"params"`
}
