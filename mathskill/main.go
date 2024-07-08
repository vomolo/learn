package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"mathskill/functions"
)

func main() {
	useMes := "Expecting: go run your-program.go data.txt"

	// Length of arguments check
	if len(os.Args) != 2 {
		fmt.Println(useMes)
		fmt.Println("Check length of arguments")
		return
	}

	filePath := os.Args[1]

	// Check .txt file extension
	if !strings.HasSuffix(filePath, ".txt") {
		fmt.Println(useMes)
		fmt.Println("Make sure it's a .txt file extension")
		os.Exit(1)
	}

	fileCheck, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			return
		}
		fmt.Println("Error checking file:", err)
		return
	}

	if fileCheck.Size() == 0 {
		fmt.Println("File is empty")
		return
	}
	file, _ := os.Open(filePath)

	defer file.Close()

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
	fmt.Printf("Average: %.0f\n", average)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard Deviation: %.0f\n", stdDev)
}
