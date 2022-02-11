package conf

import "mmocker/pkg/proc"

type Configs struct {
	Clients     []Client    `yaml:"Clients" json:"clients"`
	Application Application `yaml:"Application"`
	Processors []proc.Processor `yaml:"Processors"`
}

type Client struct {
	Name   string                 `yaml:"name" json:"name"`
	Type   string                 `yaml:"type" json:"type"`
	Params map[interface{}]interface{} `yaml:"params" json:"params"`
}

type Application struct {
	Ticker             int                `yaml:"ticker" json:"ticker"`
	NodeId             string             `yaml:"nodeId"` // Bind with the environment. If empty, set local directly.
	ObjectMockerConfig ObjectMockerConfig `yaml:"objectMockerConfig"`
}

type ObjectMockerConfig struct {
	Enable       bool   `yaml:"enable"`
	Host         string `yaml:"host"`
	SyncInterval string `yaml:"syncInterval"`
}
