package main

import "fmt"

func myAtoi(s string) int {
	var sign int = 1
	var result int
	var i int

	// Skip leading whitespace
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// Check for sign
	if i < len(s) {
		if s[i] == '+' {
			i++
		} else if s[i] == '-' {
			sign = -1
			i++
		}
	}

	// Convert digits to integer
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		result = result*10 + int(s[i]-'0')
		i++
	}

	return result * sign
}

func main() {
	fmt.Println(myAtoi("42"))              // Output: 42
	fmt.Println(myAtoi("   -42"))          // Output: -42
	fmt.Println(myAtoi("4193 with words")) // Output: 4193
	fmt.Println(myAtoi("words and 987"))   // Output: 0
	fmt.Println(myAtoi("-91283472332"))    // Output: -2147483648
}
