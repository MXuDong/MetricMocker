package funcs

const (
	TimeSecondsFunctionType       = "TimeSecondsFunction"
	TimeSecondsInHourFunctionType = "TimeSecondsInHourFunction"
	TimeMinutesFunctionType       = "TimeMinutesFunction"
	TimeHoursFunctionType         = "TimeHoursFunction"
)

var (
	TimeSecondsFunctionInitiator FuncInitiator = func() BaseFuncInterface {
		return &ModularFunction{
			BaseFunc: BaseFunc{
				IsDerivedVar: &TrueP,
				DocValue: `TimeSecondsFunction always return value in 0-59, and set offset to zero(offsetX, offsetY).
If input value is time, the return value is the seconds value in one minute.
`,
			},
			ModularUnit: 60,
			OffsetX:     0,
			OffsetY:     0,
		}
	}
	TimeSecondsInHourFunctionInitiator FuncInitiator = func() BaseFuncInterface {
		return &ModularFunction{
			BaseFunc: BaseFunc{
				IsDerivedVar: &TrueP,
				DocValue: `TimeSecondsFunction always return value in 0-3599, and set offset to zero(offsetX, offsetY).
If input value is time, the return value is the seconds value in one Hour.
`,
			},
			ModularUnit: 60 * 60,
			OffsetX:     0,
			OffsetY:     0,
		}
	}
	TimeMinutesFunctionInitiator FuncInitiator = func() BaseFuncInterface {
		return &ModularFunction{
			BaseFunc: BaseFunc{
				IsDerivedVar: &TrueP,
				DocValue: `TimeMinutesFunction always return the value in 0-59, and set offset to zero(offsetX, offsetY).
If input value is time, the return value is the minute value in one Hour.
`,
				keyFunctions: map[string]BaseFuncInterface{
					UnknownKey: &FloatFloorFunction{
						BaseFunc: BaseFunc{
							keyFunctions: map[string]BaseFuncInterface{
								UnknownKey: &BaseDivisionFunction{
									BaseFunc: BaseFunc{
										IsDerivedVar: &TrueP,
									},
									Divisor: 60,
								},
							},
						},
					},
				},
			},
			ModularUnit: 60,
			OffsetX:     0,
			OffsetY:     0,
		}
	}
	TimeHoursFunctionInitiator FuncInitiator = func() BaseFuncInterface {
		return &ModularFunction{
			BaseFunc: BaseFunc{
				IsDerivedVar: &TrueP,
				DocValue: `TimeMinutesFunction always return the value in 0-59, and set offset to zero(offsetX, offsetY).
If input value is time, the return value is the minute value in one Hour.
`,
				keyFunctions: map[string]BaseFuncInterface{
					UnknownKey: &FloatFloorFunction{
						BaseFunc: BaseFunc{
							keyFunctions: map[string]BaseFuncInterface{
								UnknownKey: &BaseDivisionFunction{
									BaseFunc: BaseFunc{
										IsDerivedVar: &TrueP,
									},
									Divisor: 60*60,
								},
							},
						},
					},
				},
			},
			ModularUnit: 24,
			OffsetX:     0,
			OffsetY:     0,
		}
	}
)
