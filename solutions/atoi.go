package solutions

import (
	"fmt"
)

// Atoi converts a string to an integer.
func Atoi(s string) (int, error) {
	// Initialize variables
	sign := 1
	result := 0
	i := 0
	n := len(s)

	// Skip leading whitespace
	// for i < n && (s[i] == ' ' || s[i] == '\t') {
	// 	i++
	// }

	// Remove all spaces and tabs manually
	var cleanedString []byte
	for j := 0; j < n; j++ {
		if s[j] != ' ' && s[j] != '\t' {
			cleanedString = append(cleanedString, s[j])
		}
	}
	s = string(cleanedString)

	n = len(s) // Update length after cleaning

	// Check for optional sign
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}

	// Process digits
	for i < n {
		digit := int(s[i] - '0')
		if digit < 0 || digit > 9 {
			return 0, fmt.Errorf("invalid character: %c", s[i])
		}

		// Check for overflow
		if result > (2147483647-digit)/10 {
			return 0, fmt.Errorf("integer overflow")
		}

		result = result*10 + digit
		i++
	}

	return sign * result, nil
}

func solutions() {
	tests := []string{"123  ", "-456", "   +789", "42abc", "2147483648", "", "   -42"}

	for _, test := range tests {
		result, err := Atoi(test)
		if err != nil {
			fmt.Printf("Atoi(%q) = error: %v\n", test, err)
		} else {
			fmt.Printf("Atoi(%q) = %d\n", test, result)
		}
	}
}
