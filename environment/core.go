package environment

import (
	"github.com/ads91/utils"
)

// NewCurve : instantiate a curve
func NewCurve(curveType string, env utils.Dict2) *Curve {
	curves := env["curves"]
	curve := curves.(utils.Dict2)[curveType].(utils.Dict2)
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
