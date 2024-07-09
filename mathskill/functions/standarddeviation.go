package functions

import "math"

func ComputeStandardDeviation(variance float64) float64 {
	if variance < 0 {
		return 0.0 // Standard deviation is always non-negative
	}
	return math.Sqrt(variance)
}
