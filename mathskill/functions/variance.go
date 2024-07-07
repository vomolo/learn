package functions

func ComputeVariance(numbers []float64, average float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += (num - average) * (num - average)
	}
	return sum / float64(len(numbers))
}
