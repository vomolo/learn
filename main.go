package main

import "fmt"

func TrimSpaces(str string) string {
	start := 0
	end := len(str) - 1

	for start < end && str[start] == ' ' {
		start++
	}

	for end > start && str[end] == ' ' {
		end--
	}

	return str[start : end+1]
}

func InputSpaces(str string) string {
	res := ""
	spaceFlag := false

	for _, char := range str {
		if char != ' ' {
			res += string(char)
			spaceFlag = false
		}else if !spaceFlag {
			res += " "
			spaceFlag =true
		}
	}
	return res
func TrimSpaces(str string) string {
	start := 0
	end := len(str) - 1

	for start <= end && str[start] == ' ' {
		start++
	}

	for end >= start && str[end] == ' ' {
		end--
	}

	return str[start : end+1]
}

func RemSpaces(str string) string {
	res := ""

	for _, char := range str {
		if char != ' ' {
			res += string(char)
		} else if char == ' ' {
			res += " "
		}
	}
	return res
}

func UnderScore(str string) string {
	res := ""

	for _, char := range str {
		if char != ' ' {
			res += string(char)
		} else if char == ' ' {
			res += "_"
		}
	}
	return res
}

func UnderscoreCapital(str string) string {
	res := ""

	for i, char := range str {
		if char >= 'A' && char <= 'Z' && i > 0 && str[i-1] != '_' {
			res += "_" // Add underscore before the capital letter
		}
		res += string(char)
	}

	return res
}

func main() {
	str := "   hey JudeTimeYYes  "
	trim := TrimSpaces(str)
	rem := RemSpaces(trim)
	fmt.Println(rem)
	under := UnderScore(rem)
	un := UnderscoreCapital(under)
	fmt.Println(un)
}

func main() {
	str := "    hello       you       there.      "
	res1 := TrimSpaces(str)
	res := InputSpaces(res1)
	fmt.Println(res)
}


package main

import (
    "fmt"
)

// Contains checks if substr is within str.
func Contains(str, substr string) bool {
    // If substr is an empty string, return true
    if substr == "" {
        return true
    }
    
    // Lengths of the strings
    strLen := len(str)
    substrLen := len(substr)

    // If the substring is longer than the string, it cannot be contained
    if substrLen > strLen {
        return false
    }

    // Iterate through the string
    for i := 0; i <= strLen-substrLen; i++ {
        // Check if the substring matches
        if str[i:i+substrLen] == substr {
            return true
        }
    }

    return false
}

func main() {
    fmt.Println(Contains("hello world", "world")) // true
    fmt.Println(Contains("hello world", "golang")) // false
    fmt.Println(Contains("hello world", ""))       // true (empty substring)
}