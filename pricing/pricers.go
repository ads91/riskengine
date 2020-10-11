package pricing

import (
	"math"
	"math/rand"
	"time"
)

// Price a bond
func (b Bond) Price() float64 {
	c := b.Coupon
	r := b.Curve.Rates
	p := 0.0
	for t, rate := range r {
		p += c / math.Pow(1+rate, float64(t+1))
	}
	return p
}

// Price a european call option
func (ec EuropeanCall) Price() float64 {
	// seed time for monte-carlo simulation
	rand.Seed(time.Now().UnixNano())
	r := ec.R
	t := ec.T
	vol := ec.Vol
	dt := DailyTimeStep
	n := ec.N
	k := ec.K
	s1 := 0.0
	for i := 0; i < n; i++ {
		s2 := ec.s0
		// run a price path
		for j := 0; j < t*365; j++ {
			s2 = s2 * (1.0 + r*dt + vol*math.Sqrt(dt)*rand.NormFloat64())
		}
		// calculate the payoff
		s1 += math.Max(0.0, s2-k)
	}
	// get the average of the simulations
	s1 = s1 / float64(n)
	// discount back to present day
	return s1 / math.Pow(1+r, float64(t))
}
