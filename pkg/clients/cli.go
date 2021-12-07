package clients

type Client interface {
	Init(value map[string]interface{}) error
	GetParam()map[string]interface{}
	PutValue(value float64, tags map[string]string)
}
