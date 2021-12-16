package conf

import (
	"mmocker/pkg"
)

type Configs struct {
	Processors  []*pkg.Processor `yaml:"processors" json:"processors"`
	Clients     []Client         `yaml:"clients" json:"clients"`
	Application Application      `yaml:"application"`
}

type Client struct {
	Name   string                 `yaml:"name" json:"name"`
	Type   string                 `yaml:"type" json:"type"`
	Params map[string]interface{} `yaml:"params" json:"params"`
}

type Application struct {
	Ticker int `yaml:"ticker" json:"ticker"`
}
