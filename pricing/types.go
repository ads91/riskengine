package pricing

import (
	"riskengine/environment"
)

var DailyTimeStep = float64(1.0 / 365.0)

type Bond struct {
	Curve  *environment.Curve
	Coupon float64
}

type EuropeanCall struct {
	s0  float64
	K   float64
	T   int     // years
	R   float64 // decimal i.e. 0.05 = 5%
	Vol float64
	N   int // number of monte-carlo paths
}
