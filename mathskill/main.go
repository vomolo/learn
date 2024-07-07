package main

import (
	"bufio"
	"fmt"
	"math"
	"mathskill/functions"
	"os"
	"strconv"
)

func main() {
	// Check if a file path is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		return
	}

	// Open the file
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("failed to open file: %s\n", err)
		return
	}
	defer file.Close()

	// Read the data

	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Printf("failed to parse number: %s\n", err)
			return
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("failed to scan file: %s\n", err)
		return
	}

	// Calculate mean
	average := math.Round(functions.ComputeAverage(numbers))

	// Calculate variance
	variance := math.Round(functions.ComputeVariance(numbers, average))

	// Calculate standard deviation
	stdDev := math.Round(functions.ComputeStandardDeviation(variance))

	// Calculate median
	median := math.Round(functions.ComputeMedian(numbers))

	// Print results
	fmt.Printf("Average: %d\n", int(average))
	fmt.Printf("Median: %d\n", int(median))
	fmt.Printf("Variance: %d\n", int(variance))
	fmt.Printf("Standard Deviation: %d\n", int(stdDev))

}
