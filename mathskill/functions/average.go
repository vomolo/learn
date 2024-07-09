package functions

func ComputeAverage(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}

	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}
