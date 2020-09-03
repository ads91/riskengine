package pricing

import (
	"riskengine/environment"
	"riskengine/utils/dict"
)

// DailyTimeStep : a day expressed as a fraction of a year
var DailyTimeStep = float64(1.0 / 365.0)

// HTTPPricer : a wrapper for an HTTP handler func to accept a dict
type HTTPPricer struct {
	Env dict.Dict2
}

// Bond : the representation of a bond
type Bond struct {
	Curve  *environment.Curve
	Coupon float64
}

// EuropeanCall : the representation of a european call option
type EuropeanCall struct {
	s0  float64
	K   float64
	T   int     // years
	R   float64 // decimal i.e. 0.05 = 5%
	Vol float64
	N   int // number of monte-carlo paths
}
