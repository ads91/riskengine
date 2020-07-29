package pricing

import (
	"math"
	"math/rand"
)

func (b Bond) Price() float64 {
	c := b.Coupon
	r := b.Curve.Rates
	p := 0.0
	for t, rate := range r {
		p += c / math.Pow(1+rate, float64(t+1))
	}
	return p
}
func (ec EuropeanCall) price() float64 {
	r := ec.R
	t := ec.T
	vol := ec.Vol
	dt := DailyTimeStep
	n := ec.N
	k := ec.K
	s1 := 0.0
	for i := 0; i < n; i++ {
		s2 := ec.S0
		for j := 0; j < t*365; j++ {
			s2 = s2 * (1.0 + r*dt + vol*math.Sqrt(dt)*rand.NormFloat64())
		}
		s1 += math.Max(0.0, s2-k)
	}
	return float64(s1/float64((t*365)+1)) / math.Pow(1+r, float64(t))
}
