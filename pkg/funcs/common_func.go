package funcs

const (
	StartZeroFuncType = "StartZeroFuncType"
)

type StartZeroFunc struct {
	BaseFunc
	StartTime float64 `key:"StartValue"`
}

func (s StartZeroFunc) KeyMap() map[string]struct{} {
	return map[string]struct{}{
		UnknownKey: {},
	}
}

func (s StartZeroFunc) Type() TypeStr {
	return StartZeroFuncType
}

func (s StartZeroFunc) Expression() string {
	return "x(0->)"
}

func (s *StartZeroFunc) Init() {
	// do nothing
	s.StartTime = -1

	s.SetKeyFunc(UnknownKey, MetadataUnitFunction{})
}

func (s StartZeroFunc) Params() map[string]interface{} {
	return map[string]interface{}{}
}

func (s *StartZeroFunc) Call(f float64) (float64, error) {
	x, err := s.Keys()[UnknownKey].Call(f)
	if err != nil {
		return 0, err
	}
	if s.StartTime == -1 {
		s.StartTime = x
	}
	return x - s.StartTime, nil
}
