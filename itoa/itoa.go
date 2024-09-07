package main

import (
	"fmt"
)

func main() {
	str := 123
	res := itoa(str)
	fmt.Println(res) // Output: "123"

	strNegative := -456
	resNegative := itoa(strNegative)
	fmt.Println(resNegative) // Output: "-456"

	strZero := 0
	resZero := itoa(strZero)
	fmt.Println(resZero) // Output: "0"
}

// itoa converts an integer to a string.
func itoa(num int) string {
	if num == 0 {
		return "0"
	}

	isNegative := num < 0
	if isNegative {
		num = -num
	}

	var result []byte

	for num > 0 {
		remainder := num % 10
		result = append(result, byte(remainder+'0'))
		num /= 10
	}

	if isNegative {
		result = append(result, '-')
	}

	reverse(result)

	return string(result) // Convert byte slice to string
}

// reverse reverses a byte slice in place.
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
