package model

type FunctionParams struct {
	UnknownKeyName string
	From           float64
	End            float64
	Step           float64
	Params         map[string]interface{}
}

// ValueMap save the input and output of function
type ValueMap struct {
	Input  float64 `json:"input"`
	Output float64 `json:"output"`
}

type FunctionCallValue struct {
	Expression string     `json:"expression"`
	Values     []ValueMap `json:"values"`
}
