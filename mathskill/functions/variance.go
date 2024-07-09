package functions

import "math"

func ComputeVariance(numbers []float64, average float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}

	var sum float64
	for _, num := range numbers {
		sum += math.Pow(num-average, 2)
	}
	return sum / float64(len(numbers))
}
