package environment

import (
	"riskengine/utils/dict"
)

// NewCurve : instantiate a curve
func NewCurve(curveType string, env dict.Dict2) *Curve {
	curves := env["curves"]
	curve := curves.(dict.Dict2)[curveType].(dict.Dict2)
	return &Curve{
		Tenors: interfaceToStringArray(curve["tenors"]),
		Rates:  interfaceToFloat64Array(curve["rates"]),
	}
}

func interfaceToStringArray(in interface{}) []string {
	var out []string
	for _, val := range in.([]interface{}) {
		out = append(out, val.(string))
	}
	return out
}

func interfaceToFloat64Array(in interface{}) []float64 {
	var out []float64
	for _, val := range in.([]interface{}) {
		out = append(out, val.(float64))
	}
	return out
}
