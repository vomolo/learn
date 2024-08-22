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
	Sort(a)
	fmt.Println(res)
	fmt.Println(a)
	}

func Sort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
		
}

}
