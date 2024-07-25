package solutions

import "fmt"

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}

	if n < 0 {
		n = -n
	}

	count := 0
	for n > 0 {
		n /= base
		count++
	}
	return count
}

func main() {
	fmt.Println(DigitLen(1010, 2))
}
