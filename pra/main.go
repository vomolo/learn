package main

import "fmt"

func Max(a []int) int {
	if len(a) < 1 {
		return 0
	}

	maxVal := a[0]

	for _, val := range a {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func main() {
	a := []int{123, 433, 435, 567, 879}
	res := Max(a)
	fmt.Println(res)
}
