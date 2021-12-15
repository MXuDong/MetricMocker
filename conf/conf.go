package conf

import (
	"mmocker/pkg"
)

type Configs struct {
	Processors []*pkg.Processor `yaml:"processors" json:"processors"`
	Clients    []Client         `yaml:"clients" json:"clients"`
}

type Client struct {
	Name   string                 `yaml:"name" json:"name"`
	Type   string                 `yaml:"type" json:"type"`
	Params map[string]interface{} `yaml:"params" json:"params"`
}
