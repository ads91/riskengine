package environment

// Curve : a data structure representing a fixed income curve
type Curve struct {
	Tenors []string
	Rates  []float64
}
