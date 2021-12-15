package clients

import (
	"mmocker/pkg"
)

type Client interface {
	Init(value map[string]interface{}) error
	GetParam() map[string]interface{}
	Output()
	Register(processor *pkg.Processor)
}
