package main

import "fmt"

func Max(a []int) int {
	if len(a) == 0 {
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
	a := []int{23, 45, 67, 22, 12, 67, 897}
	res := Max(a)
	fmt.Println(res)
}
