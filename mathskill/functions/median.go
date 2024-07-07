package functions

import "sort"

func ComputeMedian(numbers []float64) float64 {
	sort.Float64s(numbers)
	n := len(numbers)
	if n%2 == 0 {
		return (numbers[(n/2)-1] + numbers[n/2]) / 2
	}
	return numbers[n/2]
}
