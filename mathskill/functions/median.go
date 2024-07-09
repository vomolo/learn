package functions

func ComputeMedian(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}

	// Sort the slice
	sortedNumbers := make([]float64, len(numbers))
	copy(sortedNumbers, numbers)
	quickSort(sortedNumbers, 0, len(sortedNumbers)-1)

	// Calculate the median
	mid := len(sortedNumbers) / 2
	if len(sortedNumbers)%2 == 0 {
		return (sortedNumbers[mid-1] + sortedNumbers[mid]) / 2
	} else {
		return sortedNumbers[mid]
	}
}

func quickSort(arr []float64, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []float64, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
